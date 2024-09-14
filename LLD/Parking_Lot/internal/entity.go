package internal

import "time"

type ParkingLot struct {
	Id         string        `json:"id"`
	Name       string        `json:"name"`
	TotalSlots int           `json:"total_slots"`
	Levels     int           `json:"levels"`
	Slots      []ParkingSlot `json:"slots"`
}

type ParkingSlot struct {
	Id          int         `json:"id"`
	SlotNumber  int         `json:"slot_number"`
	Level       int         `json:"level"`
	IsAvailable bool        `json:"is_available"`
	Vehicle     Vehicle     `json:"vehicle"`
	Type        VehicleType `json:"type"`
	EntryTime   time.Time   `json:"entry_time"`
}

type Vehicle struct {
	Id             int         `json:"id"`
	VehicleType    VehicleType `json:"vehicle_type"`
	VehicleSubType int         `json:"vehicle_sub_type"`
	VehicleNumber  string      `json:"vehicle_number"`
	PhoneNumber    string      `json:"phone_number"`
}

// VehicleCar struct which is of type vehicle
type VehicleCar struct {
	Vehicle
	Pricing float64 `json:"pricing"`
}

// VehicleBike struct which is of type vehicle
type VehicleBike struct {
	Vehicle
	Pricing float64 `json:"pricing"`
}

// VehicleAuto struct which is of type vehicle
type VehicleAuto struct {
	Vehicle
	Pricing float64 `json:"pricing"`
}
