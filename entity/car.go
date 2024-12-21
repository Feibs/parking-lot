package entity

type Car struct {
	PlateNumber string
}

func NewCar(plateNumber string) *Car {
	return &Car{
		PlateNumber: plateNumber,
	}
}
