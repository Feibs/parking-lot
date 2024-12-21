package parking_test

import (
	e "parking-lot/entity"
	"testing"

	p "parking-lot/parking"

	"github.com/stretchr/testify/assert"
)

func TestChooseLot(t *testing.T) {
	t.Run("should return the first vacant lot when using DefaultStyle", func(t *testing.T) {
		lot1 := p.NewParkingLot(3)
		lot2 := p.NewParkingLot(5)
		lots := []*p.Lot{lot1, lot2}
		defaultStyle := p.DefaultStyle{}

		firstVacantLot := defaultStyle.ChooseLot(lots)

		assert.Equal(t, lot1, firstVacantLot)
	})
	t.Run("should return the lot with the highest limit when using MaxStyle", func(t *testing.T) {
		lot1 := p.NewParkingLot(3)
		lot2 := p.NewParkingLot(2)
		lot3 := p.NewParkingLot(5)
		lots := []*p.Lot{lot1, lot2, lot3}
		maxStyle := p.MaxStyle{}

		largestLot := maxStyle.ChooseLot(lots)

		assert.Equal(t, lot3, largestLot)
	})
	t.Run("should return the lot with the highest free space when using VacantStyle", func(t *testing.T) {
		attendant := p.NewAttendant()
		lot1 := p.NewParkingLot(3)
		lot2 := p.NewParkingLot(2)
		lot3 := p.NewParkingLot(1)
		attendant.AssignParkingLot(lot1)
		attendant.AssignParkingLot(lot2)
		attendant.AssignParkingLot(lot3)
		attendant.Park(e.NewCar("B1234AAA"))
		attendant.Park(e.NewCar("B1234BBB"))
		lots := []*p.Lot{lot1, lot2, lot3}
		vacantStyle := p.VacantStyle{}

		spaciousLot := vacantStyle.ChooseLot(lots)

		assert.Equal(t, lot2, spaciousLot)
	})
}
