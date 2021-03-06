package model

type ParkingOperator interface {
	ParkCar(licensePlate string, age int) (int, error)
	SearchCar(licensePlate string) (bool, int)
	SlotForAge(age int) string
	VehicleForAge(age int) string
	LeaveSlot(id int) (vehicleNumber string, age int, err error)
}
