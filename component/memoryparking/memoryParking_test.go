package memoryparking

import "testing"

func TestParkCar(t *testing.T) {
	InitParkingLot(5)
	pl, _ := NewInMemoryParkingLot()
	tests := []struct {
		name         string
		licensePlate string
		age          int
		wantres      int
		wantErr      bool
	}{
		{
			name:         "success",
			licensePlate: "KA-23-AB-1234",
			age:          21,
			wantres:      1,
			wantErr:      false,
		},
		{
			name:         "success",
			licensePlate: "PB-23-AB-1234",
			age:          21,
			wantres:      2,
			wantErr:      false,
		},
		{
			name:         "success",
			licensePlate: "DL-23-AB-1234",
			age:          27,
			wantres:      3,
			wantErr:      false,
		},
		{
			name:         "success",
			licensePlate: "AP-23-AB-1234",
			age:          4,
			wantres:      4,
			wantErr:      false,
		},
		{
			name:         "success",
			licensePlate: "HP-23-AB-1234",
			age:          40,
			wantres:      5,
			wantErr:      false,
		},
		{
			name:         "failure",
			licensePlate: "KA-23-AB-1234",
			age:          21,
			wantErr:      true,
		},
		{
			name:         "failure",
			licensePlate: "HR-23-AB-1234",
			age:          0,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := pl.ParkCar(tt.licensePlate, tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkCar() error: Got error = %v, Want error = %v", err, tt.wantErr)
				return
			}
			if err == nil && res != tt.wantres {
				t.Errorf("ParkCar() error: Want %v, Got %v", tt.wantres, res)
			}
		})
	}
}

func TestSearchCar(t *testing.T) {
	pl, _ := NewInMemoryParkingLot()
	tests := []struct {
		name         string
		licensePlate string
		wantres      int
		found        bool
	}{
		{
			name:         "success",
			licensePlate: "KA-23-AB-1234",
			wantres:      1,
			found:        true,
		},
		{
			name:         "success",
			licensePlate: "HR-23-AB-1234",
			found:        false,
		},
		{
			name:         "success",
			licensePlate: "HP-23-AB-1234",
			wantres:      5,
			found:        true,
		},
		{
			name:         "failure",
			licensePlate: "DL-23-AB-1234",
			wantres:      3,
			found:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, res := pl.SearchCar(tt.licensePlate)
			if found != tt.found {
				t.Errorf("SearchCar() error: Got Found = %v, Want Found = %v", found, tt.found)
				return
			}
			if found && res != tt.wantres {
				t.Errorf("SearchCar() error: Want %v, Got %v", tt.wantres, res)
			}
		})
	}

}

func TestSlotForAge(t *testing.T) {
	pl, _ := NewInMemoryParkingLot()
	tests := []struct {
		name    string
		age     int
		wantres string
	}{
		{
			name:    "success",
			age:     21,
			wantres: "1,2",
		},
		{
			name:    "failure",
			age:     62,
			wantres: "",
		},
		{
			name:    "success",
			age:     40,
			wantres: "5",
		},
		{
			name:    "failure",
			age:     0,
			wantres: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pl.SlotForAge(tt.age)
			if s != tt.wantres {
				t.Errorf("SlotForAge() error: Want %v, Got %v", tt.wantres, s)
			}
		})
	}

}

func TestVehicleForAge(t *testing.T) {
	pl, _ := NewInMemoryParkingLot()
	tests := []struct {
		name    string
		age     int
		wantres string
	}{
		{
			name:    "success",
			age:     21,
			wantres: "KA-23-AB-1234,PB-23-AB-1234",
		},
		{
			name:    "failure",
			age:     62,
			wantres: "",
		},
		{
			name:    "success",
			age:     40,
			wantres: "HP-23-AB-1234",
		},
		{
			name:    "failure",
			age:     0,
			wantres: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pl.VehicleForAge(tt.age)
			if s != tt.wantres {
				t.Errorf("VehicleForAge() error: Want %v, Got %v", tt.wantres, s)
			}
		})
	}

}

func TestLeaveSlot(t *testing.T) {
	pl, _ := NewInMemoryParkingLot()
	tests := []struct {
		name        string
		slot        int
		wantLicense string
		wantAge     int
		wantErr     bool
	}{
		{
			name:        "success",
			slot:        1,
			wantLicense: "KA-23-AB-1234",
			wantAge:     21,
			wantErr:     false,
		},
		{
			name:    "failure",
			slot:    1,
			wantErr: true,
		},
		{
			name:        "success",
			slot:        5,
			wantLicense: "HP-23-AB-1234",
			wantAge:     40,
			wantErr:     false,
		},
		{
			name:    "failure",
			slot:    0,
			wantErr: true,
		},
		{
			name:    "failure",
			slot:    10,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			license, age, err := pl.LeaveSlot(tt.slot)
			if (err != nil) != tt.wantErr {
				t.Errorf("LeaveSlot() error: Got error = %v, Want error = %v", err, tt.wantErr)
				return
			}
			if err == nil && license != tt.wantLicense {
				t.Errorf("LeaveSlot() error: Want %v, Got %v", tt.wantLicense, license)
			}
			if err == nil && age != tt.wantAge {
				t.Errorf("LeaveSlot() error: Want %v, Got %v", tt.wantAge, age)
			}
		})
	}
}
