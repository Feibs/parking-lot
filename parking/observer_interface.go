package parking

type Observer interface {
    Update(*Lot, bool)
}