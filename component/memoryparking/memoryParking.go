package memoryparking

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

func NewInMemoryParkingLot() (*parkingLot, error) {
	if parking == nil {
		return nil, errors.New("Storage Error: Parking Lot Doesn't Exist")
	}
	return parking, nil
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
	pl.availableIdx = pl.findEmptySlotIdx(pl.availableIdx)
	return slot.id
}

func (pl *parkingLot) findEmptySlotIdx(idx int) int {
	for ; idx < pl.size; idx++ {
		if pl.slots[idx].id == 0 {
			return idx
		}
	}
	return idx
}

func (pl *parkingLot) isSpaceAvailable() bool {
	if pl.filledCount >= pl.size {
		return false
	}
	return true
}

func (pl *parkingLot) indexForAge(age int) []int {
	ans := make([]int, 0)
	for i, v := range pl.slots {
		if v.driverAge == age {
			ans = append(ans, i)
		}
	}
	return ans
}

func (pl *parkingLot) emptySlot(id int) (vehicleNumber string, age int, err error) {
	slot := pl.slots[id-1]
	vehicleNumber = slot.vehicleNumber
	age = slot.driverAge

	pl.slots[id-1].id = 0
	pl.slots[id-1].vehicleNumber = ""
	pl.slots[id-1].driverAge = 0
	pl.filledCount -= 1
	pl.availableIdx = min(pl.availableIdx, id-1)
	return
}
