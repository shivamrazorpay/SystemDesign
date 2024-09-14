package main

import (
	"Parking_Lot/internal"
	"time"
)

func main() {
	pl := internal.NewParkingLot("MyParkingLot", 10, 1)
	parkingTicket, err := pl.ParkVehicle(internal.Vehicle{
		VehicleNumber: "KA01AB1234",
		VehicleType:   internal.Car,
	})
	if err != nil {
		panic(err)
	}

	// Simulating parking for 5 seconds
	time.Sleep(5 * time.Second)

	err = pl.UnparkVehicle(parkingTicket)
	if err != nil {
		panic(err)
	}
	return
}
