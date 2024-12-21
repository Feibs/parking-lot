package parking

type Style interface {
	ChooseLot(lots []*Lot) *Lot
}
