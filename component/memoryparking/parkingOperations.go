package memoryparking

import (
	"errors"
	"fmt"
	"strconv"
)

func (pl *parkingLot) ParkCar(licensePlate string, age int) (int, error) {
	if !pl.isSpaceAvailable() {
		return -1, errors.New("No slots are available, Parking Lot is full")
	}
	if duplicate, existingslot := pl.findCarLocation(licensePlate); duplicate {
		return -1, errors.New(fmt.Sprintf("Duplicate License Plate, %s is already parked at slot %d", licensePlate, existingslot))
	}

	allotedSlot := pl.addCar(licensePlate, age)
	return allotedSlot, nil
}

func (pl *parkingLot) SearchCar(licensePlate string) (bool, int) {
	return pl.findCarLocation(licensePlate)
}

func (pl *parkingLot) SlotForAge(age int) string {
	s := ""
	idx := pl.indexForAge(age)
	if len(idx) <= 0 {
		return s
	}
	for _, i := range idx {
		if len(s) != 0 {
			s += ","
		}
		s += strconv.Itoa(i + 1)
	}
	return s
}

func (pl *parkingLot) VehicleForAge(age int) string {
	s := ""
	idx := pl.indexForAge(age)
	if len(idx) <= 0 {
		return s
	}
	for _, i := range idx {
		if len(s) != 0 {
			s += ","
		}
		s += pl.slots[i].vehicleNumber
	}
	return s
}
