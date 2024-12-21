package parking

import (
	e "parking-lot/entity"
	err "parking-lot/errorcase"
)

type Lot struct {
	limit     int
	ticketCars map[*e.Ticket]*e.Car
	parkedCars map[*e.Car]bool
	observers  []Observer
}

func NewParkingLot(limit int) *Lot {
	return &Lot{
		limit:      limit,
		ticketCars: map[*e.Ticket]*e.Car{},
		parkedCars: map[*e.Car]bool{},
		observers: []Observer{},
	}
}

func (p *Lot) Register(o Observer) {
	p.observers = append(p.observers, o)
}

func (p *Lot) NotifyAll(isAvailable bool) {
	for _, observer := range(p.observers) {
        observer.Update(p, isAvailable)
    }
}

func (p *Lot) Limit() int {
	return p.limit
}

func (p *Lot) CalculateVacancy() int {
	return p.limit - len(p.ticketCars)
}

func (p *Lot) TicketCars() map[*e.Ticket]*e.Car {
	return p.ticketCars
}

func (p *Lot) IsFull() bool {
	return len(p.ticketCars) == p.limit
}

func (p *Lot) IsTicketFound(t *e.Ticket) bool {
	_, found := p.ticketCars[t]
	return found
}

func (p *Lot) IsCarFound(c *e.Car) bool {
	_, found := p.parkedCars[c]
	return found
}

func (p *Lot) CompareLotByMaxLimit(otherLot *Lot) *Lot {
	if p.limit > otherLot.limit {
		return p
	}
	return otherLot
}

func (p *Lot) CompareLotByMaxVacancy(otherLot *Lot) *Lot {
	spaceLot1 := p.CalculateVacancy()
	spaceLot2 := otherLot.CalculateVacancy()
	if spaceLot1 > spaceLot2 {
		return p
	}
	return otherLot
}

func (p *Lot) Park(c *e.Car) (*e.Ticket, error) {
	if p.IsFull() {
		return nil, err.ErrNoPosition
	}
	_, found := p.parkedCars[c]
	if found {
		return nil, err.ErrCannotParkTwice
	}
	ticket := e.NewTicket()
	p.ticketCars[&ticket] = c
	p.parkedCars[c] = true
	if p.IsFull() {
		p.NotifyAll(false)
	}
	return &ticket, nil
}

func (p *Lot) Unpark(t *e.Ticket) (*e.Car, error) {
	car, found := p.ticketCars[t]
	if !found {
		return nil, err.ErrUnrecognizedTicket
	}
	isPrevFull := p.IsFull()
	delete(p.ticketCars, t)
	delete(p.parkedCars, car)
	isCurrAvail := !p.IsFull()
	if isPrevFull && isCurrAvail {
		p.NotifyAll(true)
	}
	return car, nil
}