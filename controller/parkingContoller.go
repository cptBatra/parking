package controller

import (
	"errors"
	"fmt"
	"parking/model"
	"parking/repository"
	"strconv"
)

func createParkingLot(request []string) (model.Response, error) {
	response := model.Response{}
	size, err := strconv.Atoi(request[0])
	if err != nil || size <= 0 {
		return response, errors.New(fmt.Sprintf("Storage Error: Could not create Parking Lot, Invalid size: %v", request[0]))
	}
	err = repository.InitParking(size)
	if err != nil {
		return response, err
	}
	response.Message = fmt.Sprintf("Created parking of %d slots", size)
	return response, nil
}

func parkCar(request []string) (model.Response, error) {
	response := model.Response{}
	if len(request) < 3 {
		return response, errors.New("Incomplete Parameters for Parking car")
	}
	if len(request[0]) < 13 {
		return response, errors.New("Invalid License Plate")
	}
	licensePlate := request[0]
	age, err := strconv.Atoi(request[2])
	if err != nil || age <= 0 {
		return response, errors.New(fmt.Sprintf("Invalid Driver Age %v", request[2]))
	}

	pl, err := repository.GetParking()
	if err != nil {
		return response, err
	}
	slot, err := pl.ParkCar(licensePlate, age)
	if err != nil {
		return response, err
	}
	response.Message = fmt.Sprintf("Car with vehicle registration number \"%s\" has been parked at slot number %d", licensePlate, slot)
	return response, nil
}

func slotForAge(request []string) (model.Response, error) {
	response := model.Response{}
	age, err := strconv.Atoi(request[0])
	if err != nil || age <= 0 {
		return response, errors.New(fmt.Sprintf("Invalid age: %v", request[0]))
	}

	pl, err := repository.GetParking()
	if err != nil {
		return response, err
	}
	s := pl.SlotForAge(age)
	if len(s) < 1 {
		return response, errors.New(fmt.Sprintf("No parked car matches the query"))
	}
	response.Message = fmt.Sprintf(s)
	return response, nil
}

func slotForVehicle(request []string) (model.Response, error) {
	response := model.Response{}
	if len(request[0]) < 13 {
		return response, errors.New("Invalid License Plate")
	}
	pl, err := repository.GetParking()
	if err != nil {
		return response, err
	}
	found, slot := pl.SearchCar(request[0])
	if !found {
		return response, errors.New("No parked car matches the query")
	}
	response.Message = fmt.Sprintf("%d", slot)
	return response, nil
}

func leaveSlot(request []string) (model.Response, error) {
	response := model.Response{}
	slot, err := strconv.Atoi(request[0])
	if err != nil || slot <= 0 {
		return response, errors.New(fmt.Sprintf("Invalid slot: %v", request[0]))
	}
	pl, err := repository.GetParking()
	if err != nil {
		return response, err
	}
	vehicleNumber, age, err := pl.LeaveSlot(slot)
	if err != nil {
		return response, err
	}

	response.Message = fmt.Sprintf("Slot number %d vacated, the car with vehicle registration number \"%s\" left the space, the driver of the car was of age %d", slot, vehicleNumber, age)
	return response, nil
}

func vehicleForAge(request []string) (model.Response, error) {
	response := model.Response{}
	age, err := strconv.Atoi(request[0])
	if err != nil || age <= 0 {
		return response, errors.New(fmt.Sprintf("Invalid age: %v", request[0]))
	}
	pl, err := repository.GetParking()
	if err != nil {
		return response, err
	}
	s := pl.VehicleForAge(age)
	if len(s) < 1 {
		return response, errors.New(fmt.Sprintf("No parked car matches the query"))
	}
	response.Message = fmt.Sprintf(s)
	return response, nil
}
