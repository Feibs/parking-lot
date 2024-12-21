package parking_test

import (
	"testing"

	e "parking-lot/entity"
	s "parking-lot/errorcase"
	"parking-lot/mocks"
	p "parking-lot/parking"

	"github.com/stretchr/testify/assert"
)

func TestNotifyObserver(t *testing.T) {
	t.Run("should invoke observer's update when the lot is fully parked", func(t *testing.T) {
		mockObserver := mocks.NewObserver(t)
		lot := p.NewParkingLot(2)
		car1 := e.NewCar("B1234AAA")
		car2 := e.NewCar("B1234BBB")

		mockObserver.On("Update", lot, false)
		lot.Register(mockObserver)
		lot.Park(car1)
		lot.Park(car2)

		mockObserver.AssertNumberOfCalls(t, "Update", 1)
	})
	t.Run("should invoke observer's update when unparking from a fully parked lot", func(t *testing.T) {
		mockObserver := mocks.NewObserver(t)
		lot := p.NewParkingLot(1)
		car := e.NewCar("B1234AAA")

		mockObserver.On("Update", lot, false)
		lot.Register(mockObserver)
		ticket, _ := lot.Park(car)

		mockObserver.On("Update", lot, true)
		lot.Unpark(ticket)

		mockObserver.AssertNumberOfCalls(t, "Update", 2)
	})
}

func TestParkCar(t *testing.T) {
	t.Run("should return a ticket successfully when parking a car", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		car := e.NewCar("B1234AAA")

		ticket, _ := lot.Park(car)

		assert.NotNil(t, ticket)
	})
	t.Run("should raise a no-place-error when the parking lot is full", func(t *testing.T) {
		lot := p.NewParkingLot(1)
		cars := []*e.Car{e.NewCar("B1234AAA"), e.NewCar("B5678BBB")}

		var err error
		for _, car := range cars {
			_, err = lot.Park(car)
		}

		assert.Equal(t, s.ErrNoPosition, err)
	})
	t.Run("should raise a cannot-park-twice-error when the car has been parked", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		car := e.NewCar("B1234AAA")

		var err error
		for i := 0; i < 2; i++ {
			_, err = lot.Park(car)
		}

		assert.Equal(t, s.ErrCannotParkTwice, err)
	})
	t.Run("should occupy the parking lot when parking a car", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		car := e.NewCar("B1234AAA")

		ticket, _ := lot.Park(car)
		ticketCars := lot.TicketCars()
		_, foundInTicketCars := ticketCars[ticket]

		assert.True(t, foundInTicketCars)
		assert.Equal(t, lot.CalculateVacancy(), lot.Limit()-1)
	})
}

func TestUnparkCar(t *testing.T) {
	t.Run("should return the right car when unparking with the corresponding ticket", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		car := e.NewCar("B1234AAA")
		parkedTicket, _ := lot.Park(car)

		unparkedCar, _ := lot.Unpark(parkedTicket)

		assert.Equal(t, car, unparkedCar)
	})
	t.Run("should remove the car from parking lot when unparking", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		car := e.NewCar("B1234AAA")
		parkedTicket, _ := lot.Park(car)

		lot.Unpark(parkedTicket)
		ticketCars := lot.TicketCars()
		_, foundInTicketCars := ticketCars[parkedTicket]

		assert.False(t, foundInTicketCars)
	})
	t.Run("should raise an unrecognized-ticket-error when unparking with invalid ticket", func(t *testing.T) {
		lot := p.NewParkingLot(100)
		ticket := e.NewTicket()

		_, err := lot.Unpark(&ticket)

		assert.Equal(t, s.ErrUnrecognizedTicket, err)
	})
}
