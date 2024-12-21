package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"fmt"

	e "parking-lot/entity"
	errMsg "parking-lot/errorcase"
	p "parking-lot/parking"
)

func displayMainMenu() {
	fmt.Println()
	fmt.Println("----------------------")
	fmt.Println("Parking Lot Management")
	fmt.Println("----------------------")
	fmt.Println("0 : Display Available Parking Lot")
	fmt.Println("1 : Register Parking Lot")
	fmt.Println("2 : Park Car")
	fmt.Println("3 : Unpark Car")
	fmt.Println("4 : Change Parking Style")
	fmt.Println("5 : Exit")
	fmt.Println()
}

func displayStyle() {
	fmt.Println()
	fmt.Println("----------------------")
	fmt.Println("Parking Style Menu")
	fmt.Println("----------------------")
	fmt.Println("0 : Default")
	fmt.Println("1 : Largest Limit Lot")
	fmt.Println("2 : Most Vacant Lot")
	fmt.Println()
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return err.Error()
	}
	input = strings.TrimSpace(input)
	return input
}

func displayAvailableLots(attendant *p.Attendant) {
	fmt.Println("Available Parking Lot:")
	lots := attendant.AvailableParkingLot()
	if len(lots) == 0 {
		fmt.Println("None at the moment")
		return
	}
	for i := 0; i < len(lots); i++ {
		lot := lots[i]
		fmt.Printf("%d. Limit: %d | Vacant: %d\n", i+1, lot.Limit(), lot.CalculateVacancy())
	}
}

func registerLot(attendant *p.Attendant) {
	isMenuLimitRunning := true
	for isMenuLimitRunning {
		fmt.Print("Lot's limit: ")  
		inputLimit := input()
		limit, err := strconv.Atoi(inputLimit)
		if limit > 0 && err == nil {
			lot := p.NewParkingLot(limit)
			attendant.AssignParkingLot(lot)
			fmt.Printf("Lot of %s space(s) is registered successfully!\n", inputLimit)
			isMenuLimitRunning = false
		} else {
			fmt.Println(errMsg.ErrLimitInvalid)
		}
	}
}

func changeStyle(attendant *p.Attendant) {
	isMenuStyleRunning := true
	for isMenuStyleRunning {
		displayStyle()
		fmt.Print("Choose a parking style: ")
		optStyle := input()
		switch optStyle {
		case "0":
			attendant.SetParkingStyle(p.DefaultStyle{})
			isMenuStyleRunning = false
		case "1":
			attendant.SetParkingStyle(p.MaxStyle{})
			isMenuStyleRunning = false
		case "2":
			attendant.SetParkingStyle(p.VacantStyle{})
			isMenuStyleRunning = false
		default:
			fmt.Println(errMsg.ErrUnrecognizedStyle)
		}
	}
	fmt.Println("Parking style is set successfully!")
}

func registerCarWithPlateNum() *e.Car {
	fmt.Print("Car's plate number: ")
	inputPlateNum := input()
	return e.NewCar(inputPlateNum)
}

func park(attendant *p.Attendant, tickets *map[string]*e.Ticket) {
	car := registerCarWithPlateNum()
	ticket, err := attendant.Park(car)
	if err == nil {
		(*tickets)[ticket.ID] = ticket
		fmt.Printf("Car %s is successfully parked with Ticket %s\n", car.PlateNumber, ticket.ID)
		return
	}
	fmt.Println(err)
}

func unpark(attendant *p.Attendant, tickets *map[string]*e.Ticket) {
	fmt.Print("Ticket ID: ")
	inputTicketID := input()
	car, err := attendant.Unpark((*tickets)[inputTicketID])
	if err == nil {
		delete(*tickets, inputTicketID)
		fmt.Printf("Car %s is successfully unparked with Ticket %s\n", car.PlateNumber, inputTicketID)
		return
	}
	fmt.Println(err)
}

func main() {
	attendant := p.NewAttendant()
	tickets := make(map[string]*e.Ticket)

	isRunning := true
	for isRunning {
		displayMainMenu()
		fmt.Print("Choose a menu: ")
		opt := input()
		switch opt {
		case "0":
			displayAvailableLots(attendant)
		case "1":
			registerLot(attendant)
		case "2":
			park(attendant, &tickets)
		case "3":
			unpark(attendant, &tickets)
		case "4":
			changeStyle(attendant)
		case "5":
			fmt.Println("Thank you for using Parking Lot Management!")
			isRunning = false
		default:
			fmt.Println(errMsg.ErrUnrecognizedOptionMenu)
		}
	}
}
