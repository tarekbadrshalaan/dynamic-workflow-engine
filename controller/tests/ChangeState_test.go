package tests

import (
	"context"
	"dwf/controller"
	"testing"
	"time"

	"bou.ke/monkey"
)

//!+test
//go test -v

// TestChangeStateVoidhandlersUser done
// TestChangeStateNonVoidhandlersUser
// TestChangeStateToWrongState done
// TestChangeStateNonVoidhandlersLogic done

// TestChangeStateUser to test only voidHandler
func TestChangeStateVoidhandlersUser(t *testing.T) {

	type changeState struct {
		nextState string
		user      int
		err       string
	}
	tt := []struct {
		testname    string
		state       string
		changeState []changeState
	}{
		{
			testname: "state-Ready",
			state:    "Ready",
			changeState: []changeState{
				{"Riding", controller.USER, ""},
				{"Riding", controller.HUNTER, ""},
				{"Riding", controller.ADMIN, ""},
				{"Riding", controller.SYSTEM, ""},
				{"Riding", controller.VEHICLE, "change statues from (Ready) to (Riding) is not valid ERROR:Change vehicle state from (Ready) to (Riding) is not available for (Vehicle user)"},
			},
		},
		{
			testname: "state-Riding",
			state:    "Riding",
			changeState: []changeState{
				{"Ready", controller.USER, ""},
				{"Ready", controller.HUNTER, ""},
				{"Ready", controller.ADMIN, ""},
				{"Ready", controller.SYSTEM, ""},
				{"Ready", controller.VEHICLE, "change statues from (Riding) to (Ready) is not valid ERROR:Change vehicle state from (Riding) to (Ready) is not available for (Vehicle user)"},
			},
		},
		{
			testname: "state-Battery-Low",
			state:    "Battery-Low",
			changeState: []changeState{
				{"Bounty", controller.USER, "change statues from (Battery-Low) to (Bounty) is not valid ERROR:Change vehicle state from (Battery-Low) to (Bounty) is not available for (User user)"},
				{"Bounty", controller.HUNTER, "change statues from (Battery-Low) to (Bounty) is not valid ERROR:Change vehicle state from (Battery-Low) to (Bounty) is not available for (Hunter user)"},
				{"Bounty", controller.ADMIN, "change statues from (Battery-Low) to (Bounty) is not valid ERROR:Change vehicle state from (Battery-Low) to (Bounty) is not available for (Admin user)"},
				{"Bounty", controller.SYSTEM, ""},
				{"Bounty", controller.VEHICLE, ""},
			},
		},
		{
			testname: "state-Bounty",
			state:    "Bounty",
			changeState: []changeState{
				{"Collected", controller.USER, "change statues from (Bounty) to (Collected) is not valid ERROR:Change vehicle state from (Bounty) to (Collected) is not available for (User user)"},
				{"Collected", controller.HUNTER, ""},
				{"Collected", controller.ADMIN, ""},
				{"Collected", controller.SYSTEM, ""},
				{"Collected", controller.VEHICLE, "change statues from (Bounty) to (Collected) is not valid ERROR:Change vehicle state from (Bounty) to (Collected) is not available for (Vehicle user)"},
			},
		},
		{
			testname: "state-Collected",
			state:    "Collected",
			changeState: []changeState{
				{"Dropped", controller.USER, "change statues from (Collected) to (Dropped) is not valid ERROR:Change vehicle state from (Collected) to (Dropped) is not available for (User user)"},
				{"Dropped", controller.HUNTER, ""},
				{"Dropped", controller.ADMIN, ""},
				{"Dropped", controller.SYSTEM, ""},
				{"Dropped", controller.VEHICLE, "change statues from (Collected) to (Dropped) is not valid ERROR:Change vehicle state from (Collected) to (Dropped) is not available for (Vehicle user)"},
			},
		},
		{
			testname: "state-Dropped",
			state:    "Dropped",
			changeState: []changeState{
				{"Ready", controller.USER, "change statues from (Dropped) to (Ready) is not valid ERROR:Change vehicle state from (Dropped) to (Ready) is not available for (User user)"},
				{"Ready", controller.HUNTER, ""},
				{"Ready", controller.ADMIN, ""},
				{"Ready", controller.SYSTEM, ""},
				{"Ready", controller.VEHICLE, "change statues from (Dropped) to (Ready) is not valid ERROR:Change vehicle state from (Dropped) to (Ready) is not available for (Vehicle user)"},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testname, func(t *testing.T) {
			for _, s := range tc.changeState {

				t.Run(s.nextState+"-"+controller.UsersType[s.user], func(t *testing.T) {
					v, _ := controller.InitializeVehicle(context.Background(), "vec1", tc.state, 100)
					defer v.Terminate()
					err := v.ChangeState(s.nextState, s.user)
					if s.err != "" {
						if s.err != err.Error() {
							t.Errorf("expected error %v; got %v", s.err, err)
							return
						}
						return
					}

					if s.nextState != v.State() {
						t.Errorf("expected state %v; got %v", s.nextState, v.State())
						return
					}
				})
			}
		})
	}
}

// TestChangeStateUser to test only non voidHandler
func TestChangeStateNonVoidhandlersUser(t *testing.T) {

	type changeState struct {
		nextState string
		user      int
		err       string
	}
	tt := []struct {
		testname    string
		state       string
		changeState []changeState
	}{
		{
			testname: "state-Ready",
			state:    "Ready",
			changeState: []changeState{
				{"Unknown", controller.USER, "change statues from (Ready) to (Unknown) is not valid ERROR:Change vehicle state from (Ready) to (Unknown) is not available for (User user)"},
				{"Unknown", controller.HUNTER, "change statues from (Ready) to (Unknown) is not valid ERROR:Change vehicle state from (Ready) to (Unknown) is not available for (Hunter user)"},
				{"Unknown", controller.ADMIN, "change statues from (Ready) to (Unknown) is not valid ERROR:Change vehicle state from (Ready) to (Unknown) is not available for (Admin user)"},
				{"Unknown", controller.SYSTEM, "not allowed to change statues from (Ready) to (Unknown)"},
				{"Unknown", controller.VEHICLE, "change statues from (Ready) to (Unknown) is not valid ERROR:Change vehicle state from (Ready) to (Unknown) is not available for (Vehicle user)"},

				{"Bounty", controller.USER, "change statues from (Ready) to (Bounty) is not valid ERROR:Change vehicle state from (Ready) to (Bounty) is not available for (User user)"},
				{"Bounty", controller.HUNTER, "change statues from (Ready) to (Bounty) is not valid ERROR:Change vehicle state from (Ready) to (Bounty) is not available for (Hunter user)"},
				{"Bounty", controller.ADMIN, "change statues from (Ready) to (Bounty) is not valid ERROR:Change vehicle state from (Ready) to (Bounty) is not available for (Admin user)"},
				// {"Bounty", controller.SYSTEM, "not allowed to change statues from (Ready) to (Bounty)"},
				{"Bounty", controller.VEHICLE, "change statues from (Ready) to (Bounty) is not valid ERROR:Change vehicle state from (Ready) to (Bounty) is not available for (Vehicle user)"},
			},
		},
		{
			testname: "state-Riding",
			state:    "Riding",
			changeState: []changeState{
				{"Battery-Low", controller.USER, "change statues from (Riding) to (Battery-Low) is not valid ERROR:Change vehicle state from (Riding) to (Battery-Low) is not available for (User user)"},
				{"Battery-Low", controller.HUNTER, "change statues from (Riding) to (Battery-Low) is not valid ERROR:Change vehicle state from (Riding) to (Battery-Low) is not available for (Hunter user)"},
				{"Battery-Low", controller.ADMIN, "change statues from (Riding) to (Battery-Low) is not valid ERROR:Change vehicle state from (Riding) to (Battery-Low) is not available for (Admin user)"},
				{"Battery-Low", controller.SYSTEM, "not allowed to change statues from (Riding) to (Battery-Low)"},
				{"Battery-Low", controller.VEHICLE, "not allowed to change statues from (Riding) to (Battery-Low)"},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testname, func(t *testing.T) {
			for _, s := range tc.changeState {

				t.Run(s.nextState+"-"+controller.UsersType[s.user], func(t *testing.T) {
					v, _ := controller.InitializeVehicle(context.Background(), "vec1", tc.state, 100)
					defer v.Terminate()
					err := v.ChangeState(s.nextState, s.user)
					if s.err != "" {
						if s.err != err.Error() {
							t.Errorf("expected error %v; got %v", s.err, err)
							return
						}
						return
					}

					if s.nextState != v.State() {
						t.Errorf("expected state %v; got %v", s.nextState, v.State())
						return
					}
				})
			}
		})
	}
}

// TestChangeStateToWrongState: test change state to wrong state,
// it will be satisfied with two states (Ready,Riding)
func TestChangeStateToWrongState(t *testing.T) {

	type changeState struct {
		nextState string
		user      int
		err       string
	}
	tt := []struct {
		testname    string
		state       string
		changeState []changeState
	}{
		{
			testname: "state-Ready",
			state:    "Ready",
			changeState: []changeState{
				{"Battery-Low", controller.USER, "change statues from (Ready) to (Battery-Low) is not valid ERROR:nextState(Battery-Low) is not exist"},
				{"Battery-Low", controller.HUNTER, "change statues from (Ready) to (Battery-Low) is not valid ERROR:nextState(Battery-Low) is not exist"},
				{"Battery-Low", controller.ADMIN, "change statues from (Ready) to (Battery-Low) is not valid ERROR:nextState(Battery-Low) is not exist"},
				{"Battery-Low", controller.SYSTEM, "change statues from (Ready) to (Battery-Low) is not valid ERROR:nextState(Battery-Low) is not exist"},
				{"Battery-Low", controller.VEHICLE, "change statues from (Ready) to (Battery-Low) is not valid ERROR:nextState(Battery-Low) is not exist"},

				{"Collected", controller.USER, "change statues from (Ready) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},
				{"Collected", controller.HUNTER, "change statues from (Ready) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},
				{"Collected", controller.ADMIN, "change statues from (Ready) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},
				{"Collected", controller.SYSTEM, "change statues from (Ready) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},
				{"Collected", controller.VEHICLE, "change statues from (Ready) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},

				{"Dropped", controller.USER, "change statues from (Ready) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
				{"Dropped", controller.HUNTER, "change statues from (Ready) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
				{"Dropped", controller.ADMIN, "change statues from (Ready) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
				{"Dropped", controller.SYSTEM, "change statues from (Ready) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
				{"Dropped", controller.VEHICLE, "change statues from (Ready) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
			},
		},
		{
			testname: "state-Riding",
			state:    "Riding",
			changeState: []changeState{
				{"Unknown", controller.USER, "change statues from (Riding) to (Unknown) is not valid ERROR:nextState(Unknown) is not exist"},
				{"Unknown", controller.HUNTER, "change statues from (Riding) to (Unknown) is not valid ERROR:nextState(Unknown) is not exist"},
				{"Unknown", controller.ADMIN, "change statues from (Riding) to (Unknown) is not valid ERROR:nextState(Unknown) is not exist"},
				{"Unknown", controller.SYSTEM, "change statues from (Riding) to (Unknown) is not valid ERROR:nextState(Unknown) is not exist"},
				{"Unknown", controller.VEHICLE, "change statues from (Riding) to (Unknown) is not valid ERROR:nextState(Unknown) is not exist"},

				{"Bounty", controller.USER, "change statues from (Riding) to (Bounty) is not valid ERROR:nextState(Bounty) is not exist"},
				{"Bounty", controller.HUNTER, "change statues from (Riding) to (Bounty) is not valid ERROR:nextState(Bounty) is not exist"},
				{"Bounty", controller.ADMIN, "change statues from (Riding) to (Bounty) is not valid ERROR:nextState(Bounty) is not exist"},
				{"Bounty", controller.SYSTEM, "change statues from (Riding) to (Bounty) is not valid ERROR:nextState(Bounty) is not exist"},
				{"Bounty", controller.VEHICLE, "change statues from (Riding) to (Bounty) is not valid ERROR:nextState(Bounty) is not exist"},

				{"Collected", controller.USER, "change statues from (Riding) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},
				{"Collected", controller.HUNTER, "change statues from (Riding) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},
				{"Collected", controller.ADMIN, "change statues from (Riding) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},
				{"Collected", controller.SYSTEM, "change statues from (Riding) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},
				{"Collected", controller.VEHICLE, "change statues from (Riding) to (Collected) is not valid ERROR:nextState(Collected) is not exist"},

				{"Dropped", controller.USER, "change statues from (Riding) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
				{"Dropped", controller.HUNTER, "change statues from (Riding) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
				{"Dropped", controller.ADMIN, "change statues from (Riding) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
				{"Dropped", controller.SYSTEM, "change statues from (Riding) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
				{"Dropped", controller.VEHICLE, "change statues from (Riding) to (Dropped) is not valid ERROR:nextState(Dropped) is not exist"},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testname, func(t *testing.T) {
			for _, s := range tc.changeState {

				t.Run(s.nextState+"-"+controller.UsersType[s.user], func(t *testing.T) {
					v, _ := controller.InitializeVehicle(context.Background(), "vec1", tc.state, 100)
					defer v.Terminate()
					err := v.ChangeState(s.nextState, s.user)
					if s.err != "" {
						if s.err != err.Error() {
							t.Errorf("expected error %v; got %v", s.err, err)
							return
						}
						return
					}

					if s.nextState != v.State() {
						t.Errorf("expected state %v; got %v", s.nextState, v.State())
						return
					}
				})
			}

		})
	}
}

// TestChangeStateUser to test non voidHandler logic
func TestChangeStateNonVoidhandlersLogic(t *testing.T) {

	tt := []struct {
		testname      string
		state         string
		expectedState string
		timeNow       time.Time
		BatteryLevel  int
	}{
		{
			testname:      "Unknown-success",
			state:         "Ready",
			expectedState: "Unknown",
			timeNow:       time.Now().Add(49 * time.Hour),
			BatteryLevel:  99,
		},
		{
			testname:      "Unknown-Failed",
			state:         "Ready",
			expectedState: "Ready",
			timeNow:       time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 29, 0, 0, time.UTC),
			BatteryLevel:  99,
		},
		{
			testname:      "Bounty-success",
			state:         "Ready",
			expectedState: "Bounty",
			timeNow:       time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 31, 0, 0, time.UTC),
			BatteryLevel:  99,
		},
		{
			testname:      "Bounty-Failed",
			state:         "Ready",
			expectedState: "Ready",
			timeNow:       time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 29, 0, 0, time.UTC),
			BatteryLevel:  99,
		},
		{
			testname:      "UnknownVSBounty",
			state:         "Ready",
			expectedState: "Unknown",
			timeNow:       time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+3, 21, 31, 0, 0, time.UTC),
			BatteryLevel:  99,
		},
		{
			testname:      "BountyVSUnknown",
			state:         "Ready",
			expectedState: "Bounty",
			timeNow:       time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 21, 31, 0, 0, time.UTC),
			BatteryLevel:  99,
		},
		{
			testname:      "Battery-Low-success",
			state:         "Riding",
			expectedState: "Battery-Low",
			timeNow:       time.Now(),
			BatteryLevel:  10,
		},
		{
			testname:      "Battery-Low-Failed",
			state:         "Riding",
			expectedState: "Riding",
			timeNow:       time.Now(),
			BatteryLevel:  21,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testname, func(t *testing.T) {
			v, _ := controller.InitializeVehicle(context.Background(), "vec1", tc.state, tc.BatteryLevel)
			defer v.Terminate()
			/* start mock time */
			patch := monkey.Patch(time.Now, func() time.Time { return tc.timeNow })
			defer patch.Unpatch()
			/* end mock time */

			time.Sleep(12 * time.Millisecond)

			if tc.expectedState != v.State() {
				t.Errorf("expected state %v; got %v", tc.expectedState, v.State())
				return
			}
		})
	}
}
