package repository

import (
	"parking/component/memoryparking"
	"parking/model"
)

//GetParking contorls the type of parking implementation to be used
//i.e it controls which underlying component is to be used as storage engine to
//store parking data
//This enables easy swapping of components if required

// For Example: we can read from config file which storage to use and modify the return value accordingly
// like return sqlparking.NewSQLParkingLot()

func GetParking() (model.ParkingOperator, error) {
	return memoryparking.NewInMemoryParkingLot()
}

func InitParking(size int) error {
	return memoryparking.InitParkingLot(size)
}
