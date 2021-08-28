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
	if err != nil {
		return response, errors.New(fmt.Sprintf("Could not create Parking Lot, Invalid size: %v", request[0]))
	}
	if size <= 0 {
		return response, errors.New(fmt.Sprintf("Could not create Parking Lot, Invalid size: %v", request[0]))
	}
	err = repository.InitParkingLot(size)
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
	if err != nil {
		return response, errors.New("Invalid Driver Age")
	}

	pl := repository.GetParkingLot()
	slot, err := pl.ParkCar(licensePlate, age)
	if err != nil {
		return response, err
	}
	response.Message = fmt.Sprintf("Car with vehicle registration number %s has been parked at slot number %d", licensePlate, slot)
	return response, nil
}

func slotForAge(request []string) (model.Response, error) {
	response := model.Response{}
	_, err := strconv.Atoi(request[0])
	if err != nil {
		return response, errors.New(fmt.Sprintf("Invalid age: %v\n", request[0]))
	}

	//call Method

	response.Message = fmt.Sprintf("1,2")
	return response, nil
}

func slotForVehicle(request []string) (model.Response, error) {
	response := model.Response{}
	if len(request[0]) < 13 {
		return response, errors.New("Invalid License Plate")
	}
	//call method

	response.Message = fmt.Sprintf("2")
	return response, nil
}

func leaveSlot(request []string) (model.Response, error) {
	response := model.Response{}
	slot, err := strconv.Atoi(request[0])
	if err != nil {
		return response, errors.New(fmt.Sprintf("Invalid slot: %v\n", request[0]))
	}

	//call method

	response.Message = fmt.Sprintf("Slot number %d vacated, the car with vehicle registration number %s left the space, the driver of the car was of age %d", slot, "ka-01-hh-1234", 21)
	return response, nil
}

func vehicleForAge(request []string) (model.Response, error) {
	response := model.Response{}
	_, err := strconv.Atoi(request[0])
	if err != nil {
		return response, errors.New(fmt.Sprintf("Invalid age: %v\n", request[0]))
	}
	//call method

	response.Message = fmt.Sprintf("ka-1234,pb-1234")
	return response, nil
}
