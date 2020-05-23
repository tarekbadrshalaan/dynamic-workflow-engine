package tests

import (
	"context"
	"gwf/controller"
	"testing"
)

//!+test
//go test -v

func TestAdminForceChangeState(t *testing.T) {

	// Initialize only one vehicle to test stress changes
	v, _ := controller.InitializeVehicle(context.Background(), "vec1", "Ready", 99)
	defer v.Terminate()

	tt := []struct {
		testname        string
		adminState      string
		availableStates []string
		usertype        int
		err             string
	}{
		{
			testname:        "Ready-Bounty",
			adminState:      "Bounty",
			usertype:        controller.ADMIN,
			availableStates: []string{"Collected"},
		},
		{
			testname:        "Riding-Riding",
			adminState:      "Riding",
			usertype:        controller.ADMIN,
			availableStates: []string{"Ready", "Battery-Low"},
		},
		{
			testname:        "Riding-Battery-Low",
			adminState:      "Battery-Low",
			usertype:        controller.ADMIN,
			availableStates: []string{"Bounty"},
		},
		{
			testname:   "Battery-Low-Unknown",
			adminState: "Unknown",
			usertype:   controller.ADMIN,
		},
		{
			testname:        "Unknown-Bounty",
			adminState:      "Bounty",
			usertype:        controller.ADMIN,
			availableStates: []string{"Collected"},
		},
		{
			testname:        "Bounty-Dropped",
			adminState:      "Dropped",
			usertype:        controller.ADMIN,
			availableStates: []string{"Ready"},
		},
		{
			testname:        "Dropped-Battery-Low",
			adminState:      "Battery-Low",
			usertype:        controller.ADMIN,
			availableStates: []string{"Bounty"},
		},
		{
			testname:        "Battery-Low-Ready",
			adminState:      "Ready",
			usertype:        controller.ADMIN,
			availableStates: []string{"Riding", "Unknown", "Bounty"},
		},
		{
			testname:   "Ready-non-exist",
			adminState: "non-exist",
			usertype:   controller.ADMIN,
			err:        "State (non-exist) is not exist",
		},
		{
			testname:   "Ready-Riding",
			adminState: "Riding",
			usertype:   controller.USER,
			err:        "Non Admin User try to force change vehicle state (0)",
		},
		{
			testname:   "Ready-Battery-Low",
			adminState: "Battery-Low",
			usertype:   controller.HUNTER,
			err:        "Non Admin User try to force change vehicle state (1)",
		},
		{
			testname:   "Ready-Bounty",
			adminState: "Bounty",
			usertype:   controller.SYSTEM,
			err:        "Non Admin User try to force change vehicle state (4)",
		},
		{
			testname:   "Ready-Unknown",
			adminState: "Unknown",
			usertype:   controller.VEHICLE,
			err:        "Non Admin User try to force change vehicle state (3)",
		},
	}
	for _, tc := range tt {
		t.Run(tc.testname, func(t *testing.T) {

			err := v.AdminForceChangeState(tc.adminState, tc.usertype)
			if tc.err != "" {
				if tc.err != err.Error() {
					t.Errorf("expected error %v; got %v", tc.err, err)
					return
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error %v", err)
			}

			actualState, actualAvailableStates := v.AvailableStates()
			if tc.adminState != actualState {
				t.Errorf("expected state %v; got %v", tc.adminState, actualState)
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
