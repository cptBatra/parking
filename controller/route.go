package controller

import (
	"errors"
	"fmt"
	"parking/model"
)

func RouteInstruction(i model.Instruction) (model.Response, error) {

	switch i.Command {
	case "Create_parking_lot":
		return createParkingLot(i.Arguments)
	case "Park":
		return parkCar(i.Arguments)
	case "Slot_numbers_for_driver_of_age":
		return slotForAge(i.Arguments)
	case "Slot_number_for_car_with_number":
		return slotForVehicle(i.Arguments)
	case "Leave":
		return leaveSlot(i.Arguments)
	case "Vehicle_registration_number_for_driver_of_age":
		return vehicleForAge(i.Arguments)
	default:
		return model.Response{}, errors.New(fmt.Sprintf("Unknown command: %s\n", i.Command))
	}
}
