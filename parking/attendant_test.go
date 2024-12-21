package parking_test

import (
	e "parking-lot/entity"
	s "parking-lot/errorcase"
	p "parking-lot/parking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvailableLot(t *testing.T) {
	t.Run("should append the list of available lots when the lot's limit suffice", func(t *testing.T) {
		attendant := p.NewAttendant()
		car := e.NewCar("B1234AAA")
		lot1 := p.NewParkingLot(1)
		lot2 := p.NewParkingLot(5)
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)
		attendant.Park(car)

		availLots := attendant.AvailableParkingLot()

		assert.Equal(t, 1, len(availLots))
	})
}

func TestParkCarWithAttendant(t *testing.T) {
	t.Run("should return a ticket successfully when parking a car", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		attendant := p.NewAttendant()
		car := e.NewCar("B1234AAA")
		attendant.AssignParkingLot(lot)

		ticket, _ := attendant.Park(car)

		assert.NotNil(t, ticket)
	})
	t.Run("should raise a no-place-error when the parking lot is full", func(t *testing.T) {
		lot := p.NewParkingLot(1)
		attendant := p.NewAttendant()
		attendant.AssignParkingLot(lot)
		cars := []*e.Car{e.NewCar("B1234AAA"), e.NewCar("B5678BBB")}
		
		var err error
		for _, car := range(cars){
			_, err = attendant.Park(car)
		}

		assert.Equal(t, s.ErrNoPosition, err)
	})
	t.Run("should raise a cannot-park-twice-error when the car has been parked", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		attendant := p.NewAttendant()
		attendant.AssignParkingLot(lot)
		car := e.NewCar("B1234AAA")

		var err error
		for i := 0; i < 2; i++ {
			_, err = attendant.Park(car)
		}

		assert.Equal(t, s.ErrCannotParkTwice, err)
	})
	t.Run("should park car to another lot when the current lot is full", func(t *testing.T) {
		attendant := p.NewAttendant()
		car1 := e.NewCar("B1234AAA")
		car2 := e.NewCar("B5678BBB")
		lot1 := p.NewParkingLot(1)
		lot2 := p.NewParkingLot(5)
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)
		attendant.Park(car1)

		ticket, _ := attendant.Park(car2)

		assert.NotNil(t, ticket)
	})
	t.Run("should raise a no-place-error when parking a full lot", func(t *testing.T) {
		attendant := p.NewAttendant()
		car1 := e.NewCar("B1234AAA")
		car2 := e.NewCar("B5678BBB")
		lot1 := p.NewParkingLot(1)
		lot2 := p.NewParkingLot(1)
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)
		attendant.Park(car1)
		attendant.Park(car2)

		_, err := attendant.Park(e.NewCar("B1111XYZ"))

		assert.Equal(t, s.ErrNoPosition, err)
	})
	t.Run("should not park twice even in different parking lot", func(t *testing.T) {
		attendant := p.NewAttendant()
		car := e.NewCar("B1234AAA")
		lot1 := p.NewParkingLot(1)
		lot2 := p.NewParkingLot(5)
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)

		var err error
		for i := 0; i < 2; i++ {
			_, err = attendant.Park(car)
		}

		assert.Equal(t, s.ErrCannotParkTwice, err)
	})
	t.Run("should be notified when the subscribed lot is full", func(t *testing.T) {
		attendant := p.NewAttendant()
		car1 := e.NewCar("B1234AAA")
		car2 := e.NewCar("B1234BBB")
		lot1 := p.NewParkingLot(2)
		lot2 := p.NewParkingLot(2)
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)
		attendant.Park(car1)
		attendant.Park(car2)

		availLots := attendant.AvailableParkingLot()
		countAvailLots := len(availLots)

		assert.Equal(t, 1, countAvailLots)
		assert.Equal(t, lot2, availLots[0])
	})
	t.Run("should use the new parking style when setting the parking style accordingly", func(t *testing.T) {
		lot1 := p.NewParkingLot(3)
		lot2 := p.NewParkingLot(5)
		attendant := p.NewAttendant()
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)
		car := e.NewCar("B1234AAA")

		attendant.SetParkingStyle(p.MaxStyle{})
		ticket, _ := attendant.Park(car)

		assert.True(t, lot2.IsTicketFound(ticket))
	})
}

func TestUnparkCarWithAttendant(t *testing.T) {
	t.Run("should return the right car when attendant unparking with the corresponding ticket", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		car := e.NewCar("B1234AAA")
		attendant := p.NewAttendant()
		attendant.AssignParkingLot(lot)
		parkedTicket, _ := attendant.Park(car)

		unparkedCar, _ := attendant.Unpark(parkedTicket)
		
		assert.Equal(t, car, unparkedCar)
	})
	t.Run("should raise an unrecognized-ticket-error when unparking with invalid ticket", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		attendant := p.NewAttendant()
		attendant.AssignParkingLot(lot)
		ticket := e.NewTicket()
		
		_, err := attendant.Unpark(&ticket)

		assert.Equal(t, s.ErrUnrecognizedTicket, err)
	})
	t.Run("should return the car among all the lots when unparking that car", func(t *testing.T) {
		attendant := p.NewAttendant()
		car1 := e.NewCar("B1234AAA")
		car2 := e.NewCar("B5678BBB")
		lot1 := p.NewParkingLot(1)
		lot2 := p.NewParkingLot(5)
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)
		attendant.Park(car1)
		ticket, _ := attendant.Park(car2)
		
		car, _ := attendant.Unpark(ticket)

		assert.NotNil(t, car)
	})
	t.Run("should be notified when the subscribed lot is available", func(t *testing.T) {
		attendant := p.NewAttendant()
		car1 := e.NewCar("B1234AAA")
		car2 := e.NewCar("B1234BBB")
		lot1 := p.NewParkingLot(1)
		lot2 := p.NewParkingLot(2)
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)
		ticketCar1, _ := attendant.Park(car1)
		attendant.Park(car2)
		attendant.Unpark(ticketCar1)
		
		availLots := attendant.AvailableParkingLot()
		countAvailLots := len(availLots)

		assert.Equal(t, 2, countAvailLots)
	})
}