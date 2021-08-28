package repository

import (
	"parking/component/memoryparking"
	"parking/model"
)

func GetParking() model.ParkingOperator {
	return memoryparking.NewInMemoryParkingLot()
}

func InitParking(size int) error {
	return memoryparking.InitParkingLot(size)
}
