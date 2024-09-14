package internal

import "errors"

func VehicleCharge(vehicleType VehicleType) (int, error) {
	switch vehicleType {
	case Car:
		return 100, nil
	case Bike:
		return 50, nil
	case Auto:
		return 20, nil
	default:
		return 0, errors.New("invalid vehicle type")
	}
}
