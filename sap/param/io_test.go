package param

import (
	"gitlab.suse.de/guohouzuo/saptune/system"
	"testing"
)

func TestIOElevators(t *testing.T) {
	if system.IsUserOBS() {
		t.Skip() // IO elevators cannot be read on build service
	}
	inspected, err := BlockDeviceSchedulers{}.Inspect()
	if err != nil {
		t.Fatal(err, inspected)
	}
	if len(inspected.(BlockDeviceSchedulers).SchedulerChoice) == 0 {
		t.Fatal(inspected)
	}
	for name, elevator := range inspected.(BlockDeviceSchedulers).SchedulerChoice {
		if name == "" || elevator == "" {
			t.Fatal(inspected)
		}
	}
	optimised, err := inspected.Optimise("noop")
	if err != nil {
		t.Fatal(err)
	}
	if len(optimised.(BlockDeviceSchedulers).SchedulerChoice) == 0 {
		t.Fatal(optimised)
	}
	for name, elevator := range optimised.(BlockDeviceSchedulers).SchedulerChoice {
		if name == "" || elevator != "noop" {
			t.Fatal(optimised)
		}
	}
}