package tests

import (
	"context"
	"dwf/exvehicle"
	"testing"
)

//!+test
//go test -v

func TestInitializeVehicle(t *testing.T) {

	tt := []struct {
		testname          string
		id                string
		state             string
		availableStates   []string
		batteryPercentage int
		err               string
	}{
		{
			testname:          "success-Ready",
			id:                "vec1",
			state:             "Ready",
			availableStates:   []string{"Riding", "Unknown", "Bounty"},
			batteryPercentage: 50,
		},
		{
			testname:          "success-Riding",
			id:                "vec1",
			state:             "Riding",
			availableStates:   []string{"Ready", "Battery-Low"},
			batteryPercentage: 70,
		},
		{
			testname:          "success-Riding-Battery-Low",
			id:                "vec1",
			state:             "Riding",
			availableStates:   []string{"Ready", "Battery-Low"},
			batteryPercentage: 1,
		},
		{
			testname:          "success-Battery-Low",
			id:                "vec1",
			state:             "Battery-Low",
			availableStates:   []string{"Bounty"},
			batteryPercentage: 10,
		},
		{
			testname:          "success-Bounty",
			id:                "vec1",
			state:             "Bounty",
			availableStates:   []string{"Collected"},
			batteryPercentage: 3,
		},
		{
			testname:          "success-Collected",
			id:                "vec1",
			state:             "Collected",
			availableStates:   []string{"Dropped"},
			batteryPercentage: 0,
		},
		{
			testname:          "success-Dropped",
			id:                "vec1",
			state:             "Dropped",
			availableStates:   []string{"Ready"},
			batteryPercentage: 99,
		},
		{
			testname:          "success-Unknown",
			id:                "vec1",
			state:             "Unknown",
			batteryPercentage: 99,
		},
		{
			testname:          "nonexist state",
			id:                "vec1",
			state:             "nonexist",
			batteryPercentage: 50,
			err:               "State (nonexist) is not exist",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testname, func(t *testing.T) {
			v, err := exvehicle.InitializeVehicle(context.Background(), tc.id, tc.state, tc.batteryPercentage)
			if tc.err != "" {
				if tc.err != err.Error() {
					t.Errorf("expected error %v; got %v", tc.err, err)
					return
				}
				return
			}
			defer v.Terminate()

			if tc.state != v.State() {
				t.Errorf("expected state %v; got %v", tc.state, v.State())
				return
			}

			if tc.id != v.ID {
				t.Errorf("expected Id %v; got %v", tc.id, v.ID)
				return
			}

			actualState, actualAvailableStates := v.AvailableStates()
			if tc.state != actualState {
				t.Errorf("expected state %v; got %v", tc.state, actualState)
				return
			}

			if len(tc.availableStates) != len(actualAvailableStates) {
				t.Errorf("expected available state length %v; got %v", len(tc.availableStates), len(actualAvailableStates))
				return
			}

			/*WARNING: issue performance, don't use it in production*/
			for _, expectedAVS := range tc.availableStates {
				founded := false
				for _, actualAVS := range actualAvailableStates {
					if expectedAVS == actualAVS {
						founded = true
						continue
					}
				}
				if !founded {
					t.Errorf("expected available state %v; didn't find", expectedAVS)
				}
			}
			/* */
		})
	}
}

//!-tests
