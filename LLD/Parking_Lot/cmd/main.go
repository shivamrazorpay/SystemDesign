package main

import (
	"Parking_Lot/internal"
	"fmt"
	"time"
)

func main() {
	pl := internal.NewParkingLot("MyParkingLot", 10, 1)
	parkingTicket, err := pl.ParkVehicle(internal.Vehicle{
		VehicleNumber:  "KA01AB1234",
		VehicleType:    internal.Car,
		VehicleSubType: internal.SUV,
		PhoneNumber:    "9876543210",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Vehicle parked successfully")
	fmt.Println("Parking Ticket: ", parkingTicket)

	// Simulating parking for 5 seconds
	time.Sleep(5 * time.Second)

	err = pl.UnparkVehicle(parkingTicket.Vehicle.VehicleNumber)
	if err != nil {
		panic(err)
	}
	return
}
