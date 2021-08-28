package repository

import "errors"

var parking *parkingLot

type parkingLot struct {
	size         int
	slots        []spot
	filledCount  int
	availableIdx int
}

type spot struct {
	id            int
	vehicleNumber string
	driverAge     int
}

type parkingOperater interface {
	ParkCar(licensePlate string, age int) (int, error)
	SearchCar(licensePlate string) (bool, int)
}

func GetParkingLot() *parkingLot {
	return parking
}

func InitParkingLot(n int) error {

	if parking != nil {
		return errors.New("Parking Lot already exists")
	}
	spots := make([]spot, n)
	lot := &parkingLot{
		size:         n,
		slots:        spots,
		filledCount:  0,
		availableIdx: 0,
	}
	parking = lot
	return nil
}

func (pl *parkingLot) findCarLocation(licensePlate string) (bool, int) {
	for _, v := range pl.slots {
		if licensePlate == v.vehicleNumber {
			return true, v.id
		}
	}
	return false, -1
}

func (pl *parkingLot) addCar(licensePlate string, age int) int {
	slot := spot{
		id:            pl.availableIdx + 1,
		vehicleNumber: licensePlate,
		driverAge:     age,
	}
	pl.slots[pl.availableIdx] = slot
	pl.filledCount += 1
	pl.availableIdx += 1
	return pl.availableIdx
}

func (pl *parkingLot) isSpaceAvailable() bool {
	if pl.filledCount >= pl.size {
		return false
	}
	return true
}
