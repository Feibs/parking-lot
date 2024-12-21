package parking

import (
	e "parking-lot/entity"
	errMsg "parking-lot/errorcase"
)

type Attendant struct {
	assignedLot []*Lot
	availLot    []*Lot
	style       Style
}

func NewAttendant() *Attendant {
	return &Attendant{
		assignedLot: []*Lot{},
		availLot:    []*Lot{},
		style:       DefaultStyle{},
	}
}

func (a *Attendant) SetParkingStyle(s Style) {
	a.style = s
}

func (a *Attendant) Update(p *Lot, isAvailable bool) {
	if !isAvailable {
		for i := 0; i < len(a.availLot); i++ {
			if a.availLot[i] == p {
				a.availLot = append(a.availLot[:i], a.availLot[i+1:]...)
				return
			}
		}
	}
	a.availLot = append([]*Lot{p}, a.availLot...)
}

func (a *Attendant) AssignParkingLot(p *Lot) {
	a.assignedLot = append(a.assignedLot, p)
	if !p.IsFull() {
		a.availLot = append(a.availLot, p)
	}
	p.Register(a)
}

func (a *Attendant) AvailableParkingLot() []*Lot {
	return a.availLot
}

func (a *Attendant) Park(c *e.Car) (*e.Ticket, error) {
	for _, lot := range a.assignedLot {
		if lot.IsCarFound(c) {
			return nil, errMsg.ErrCannotParkTwice
		}
	}
	if len(a.availLot) == 0 {
		return nil, errMsg.ErrNoPosition
	}
	chosenLot := a.style.ChooseLot(a.availLot)
	return chosenLot.Park(c)
}

func (a *Attendant) Unpark(t *e.Ticket) (*e.Car, error) {
	for _, lot := range a.assignedLot {
		if lot.IsTicketFound(t) {
			return lot.Unpark(t)
		}
	}
	return nil, errMsg.ErrUnrecognizedTicket
}
