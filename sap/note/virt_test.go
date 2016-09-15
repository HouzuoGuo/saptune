package note

import (
	"gitlab.suse.de/guohouzuo/saptune/system"
	"testing"
)

func TestVmwareGuestIOElevator(t *testing.T) {
	if system.IsUserOBS() {
		t.Skip("IO elevator settings cannot be read on build service")
	}
	ioel, err := VmwareGuestIOElevator{}.Initialise()
	if ioel.Name() == "" {
		t.Fatal(ioel.Name())
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(ioel.(VmwareGuestIOElevator).BlockDeviceSchedulers.SchedulerChoice) == 0 {
		t.Fatal(ioel)
	}
	for name, elevator := range ioel.(VmwareGuestIOElevator).BlockDeviceSchedulers.SchedulerChoice {
		if name == "" || elevator == "" {
			t.Fatal(ioel)
		}
	}
	optimised, err := ioel.Optimise()
	if err != nil {
		t.Fatal(err)
	}
	if len(ioel.(VmwareGuestIOElevator).BlockDeviceSchedulers.SchedulerChoice) == 0 {
		t.Fatal(ioel)
	}
	// All elevators now must be set to noop
	for name, elevator := range optimised.(VmwareGuestIOElevator).BlockDeviceSchedulers.SchedulerChoice {
		if name == "" || elevator != "noop" {
			t.Fatal(optimised)
		}
	}
}
