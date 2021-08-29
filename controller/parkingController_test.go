package controller

import "testing"

func TestParkCar(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		wantres string
		wantErr bool
	}{
		{
			name:    "failure",
			args:    []string{"PB-01-AA-123", "21"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"PB-01-AA-1234", "21"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"PB-01-AA-1234", "driver_age", "ab"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := parkCar(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkCar() error: Got error = %v, Want error = %v", err, tt.wantErr)
				return
			}
			if res.Message != tt.wantres {
				t.Errorf("ParkCar() error: Want %v, Got %v", tt.wantres, res.Message)
			}
		})
	}
}

func TestSlotForAge(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		wantres string
		wantErr bool
	}{
		{
			name:    "failure",
			args:    []string{"ab"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"-1"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{" "},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := slotForAge(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("slotForAge() error: Got error = %v, Want error = %v", err, tt.wantErr)
				return
			}
			if res.Message != tt.wantres {
				t.Errorf("slotForAge() error: Want %v, Got %v", tt.wantres, res.Message)
			}
		})
	}
}

func TestSlotForVehicle(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		wantres string
		wantErr bool
	}{
		{
			name:    "failure",
			args:    []string{" "},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"KA-12-AB-333"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"ab"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := slotForVehicle(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("slotForVehicle() error: Got error = %v, Want error = %v", err, tt.wantErr)
				return
			}
			if res.Message != tt.wantres {
				t.Errorf("slotForVehicle() error: Want %v, Got %v", tt.wantres, res.Message)
			}
		})
	}
}

func TestVehicleForAge(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		wantres string
		wantErr bool
	}{
		{
			name:    "failure",
			args:    []string{"ab"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"-1"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{" "},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := vehicleForAge(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("vehicleForAge() error: Got error = %v, Want error = %v", err, tt.wantErr)
				return
			}
			if res.Message != tt.wantres {
				t.Errorf("vehicleForAge() error: Want %v, Got %v", tt.wantres, res.Message)
			}
		})
	}
}

func TestLeaveSlot(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		wantres string
		wantErr bool
	}{
		{
			name:    "failure",
			args:    []string{"ab"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"-1"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"0"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := leaveSlot(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("leaveSlot() error: Got error = %v, Want error = %v", err, tt.wantErr)
				return
			}
			if res.Message != tt.wantres {
				t.Errorf("leaveSlot() error: Want %v, Got %v", tt.wantres, res.Message)
			}
		})
	}
}

func TestCreateParkingLot(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		wantres string
		wantErr bool
	}{
		{
			name:    "failure",
			args:    []string{"ab"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"-1"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"0"},
			wantErr: true,
		},
		{
			name:    "failure",
			args:    []string{"2.4"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := leaveSlot(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("createParkingLot() error: Got error = %v, Want error = %v", err, tt.wantErr)
				return
			}
			if res.Message != tt.wantres {
				t.Errorf("createParkingLot() error: Want %v, Got %v", tt.wantres, res.Message)
			}
		})
	}
}
