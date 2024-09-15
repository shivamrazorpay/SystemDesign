package internal

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func NewParkingLot(name string, totalslots int, levels int) *ParkingLot {
	slots := make([]ParkingSlot, totalslots)
	for i := 0; i < totalslots; i++ {
		slots[i] = ParkingSlot{
			Id:          i,
			SlotNumber:  i + 1,
			Level:       1,
			IsAvailable: true,
		}
	}
	return &ParkingLot{
		Id:         uuid.New().String(),
		Name:       name,
		TotalSlots: totalslots,
		Levels:     levels,
		Slots:      slots,
	}
}

func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) (ParkingTicket, error) {

	// Validate Vehicle
	err := vehicle.Validate()
	if err != nil {
		return ParkingTicket{}, err
	}

	// Allot Slot
	for i := range pl.Slots {
		if pl.Slots[i].IsAvailable {
			pl.Slots[i].IsAvailable = false
			pl.Slots[i].Vehicle = vehicle
			pl.Slots[i].EntryTime = time.Now()
			return ParkingTicket{
				Vehicle:     vehicle,
				ParkingSlot: pl.Slots[i],
			}, nil
		}
	}
	return ParkingTicket{}, errors.New("no slots available")
}

func (pl *ParkingLot) UnparkVehicle(vehicleNumber string) error {
	for _, slot := range pl.Slots {
		if slot.Vehicle.VehicleNumber == vehicleNumber {
			fmt.Println("Vehicle un-parked successfully")
			charge := pl.CalculateParkingCharges(slot)
			fmt.Printf("Parking charges for vehicle %s is %f\n", vehicleNumber, charge)
			slot.IsAvailable = true
			slot.Vehicle = Vehicle{}
			return nil
		}
	}
	return errors.New("vehicle not found")
}

func (pl *ParkingLot) CalculateParkingCharges(slot ParkingSlot) float64 {
	duration := time.Now().Sub(slot.EntryTime)
	charge, _ := VehicleCharge(slot.Vehicle.VehicleType)
	return duration.Minutes() * (float64(charge) / 60)
}
