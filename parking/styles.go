package parking

type DefaultStyle struct{}

type MaxStyle struct{}

type VacantStyle struct{}

func (d DefaultStyle) ChooseLot(lots []*Lot) *Lot {
	return lots[0]
}

func (m MaxStyle) ChooseLot(lots []*Lot) *Lot {
	largestLot := lots[0]
	for _, lot := range lots[1:] {
		largestLot = largestLot.CompareLotByMaxLimit(lot)
	}
	return largestLot
}

func (v VacantStyle) ChooseLot(lots []*Lot) *Lot {
	spaciousLot := lots[0]
	for _, lot := range lots[1:] {
		spaciousLot = spaciousLot.CompareLotByMaxVacancy(lot)
	}
	return spaciousLot
}
