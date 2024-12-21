package parking

type Publisher interface {
    Register(observer Observer)
    NotifyAll(isAvailable bool)
}