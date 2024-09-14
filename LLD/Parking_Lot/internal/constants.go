package internal

type VehicleType int

const (
	Car VehicleType = iota
	Bike
	Auto
)

type CarSubType int

const (
	SUV CarSubType = iota
	Sedan
)
