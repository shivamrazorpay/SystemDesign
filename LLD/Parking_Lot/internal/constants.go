package internal

type VehicleType int
type VehicleSubType int

const (
	InvalidVehicle VehicleType = iota
	Car
	Bike
	Auto
)

const (
	InvalidSubType VehicleSubType = iota
	// Car subtypes
	SUV
	Sedan
	Hatchback

	// Bike subtypes
	SportsBike
	Cruiser
	Scooter

	// Auto subtypes
	PassengerAuto
	GoodsAuto
)
