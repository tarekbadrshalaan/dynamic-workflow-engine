package benchmarks

import (
	"context"
	"dwf/controller"
	"testing"
)

//!+bench

// go test -bench=BenchmarkChangeStateVoidhandlersUser -benchmem
// BenchmarkChangeStateVoidhandlersUser to Benchmark only voidHandler
func BenchmarkChangeStateVoidhandlersUser(b *testing.B) {

	tt := []struct {
		testname  string
		state     string
		nextState string
		user      int
	}{
		{
			testname:  "state-Ready",
			state:     "Ready",
			nextState: "Riding",
			user:      controller.USER,
		},
		{
			testname:  "state-Riding",
			state:     "Ready",
			nextState: "Riding",
			user:      controller.USER,
		},
		{
			testname:  "state-Battery-Low",
			state:     "Battery-Low",
			nextState: "Bounty",
			user:      controller.VEHICLE,
		},
		{
			testname:  "state-Bounty",
			state:     "Bounty",
			nextState: "Collected",
			user:      controller.ADMIN,
		},
		{
			testname:  "state-Collected",
			state:     "Collected",
			nextState: "Dropped",
			user:      controller.HUNTER,
		},
		{
			testname:  "state-Dropped",
			state:     "Dropped",
			nextState: "Ready",
			user:      controller.SYSTEM,
		},
	}

	for _, tc := range tt {
		b.Run(tc.testname, func(b *testing.B) {
			v, _ := controller.InitializeVehicle(context.Background(), "vec1", tc.state, 100)
			defer v.Terminate()
			for index := 0; index < b.N; index++ {
				v.ChangeState(tc.nextState, tc.user)
			}
		})
	}
}

/*
go test -bench=BenchmarkChangeStateVoidhandlersUser -benchmem
goos: linux
goarch: amd64
pkg: dwf/controller/benchmarks
BenchmarkChangeStateVoidhandlersUser/state-Ready-8         	 1930252	       644 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Riding-8        	 1835732	       623 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Battery-Low-8   	 1961193	       637 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Bounty-8        	 1949448	       673 ns/op	     272 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Collected-8     	 1660357	       698 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Dropped-8       	 1948905	       627 ns/op	     240 B/op	       9 allocs/op
PASS
ok  	dwf/controller/benchmarks	13.476s
*/

// go test -bench=BenchmarkChangeStateNonVoidhandlersUser -benchmem
// BenchmarkChangeStateNonVoidhandlersUser to Benchmark only non voidHandler
func BenchmarkChangeStateNonVoidhandlersUser(b *testing.B) {
	tt := []struct {
		testname  string
		state     string
		nextState string
		user      int
	}{
		{
			testname:  "Ready-Unknown",
			state:     "Ready",
			nextState: "Unknown",
			user:      controller.SYSTEM,
		},
		{
			testname:  "Ready-Bounty",
			state:     "Ready",
			nextState: "Bounty",
			user:      controller.SYSTEM,
		},
		{
			testname:  "Riding-Battery-Low",
			state:     "Riding",
			nextState: "Battery-Low",
			user:      controller.SYSTEM,
		},
	}

	for _, tc := range tt {
		b.Run(tc.testname, func(b *testing.B) {
			v, _ := controller.InitializeVehicle(context.Background(), "vec1", tc.state, 100)
			defer v.Terminate()
			for index := 0; index < b.N; index++ {
				v.ChangeState(tc.nextState, tc.user)
			}
		})
	}
}

/*
go test -bench=BenchmarkChangeStateNonVoidhandlersUser -benchmem
goos: linux
goarch: amd64
pkg: dwf/controller/benchmarks
BenchmarkChangeStateNonVoidhandlersUser/Ready-Unknown-8         	 3053121	       377 ns/op	     128 B/op	       5 allocs/op
BenchmarkChangeStateNonVoidhandlersUser/Ready-Bounty-8          	 2640531	       453 ns/op	     128 B/op	       5 allocs/op
BenchmarkChangeStateNonVoidhandlersUser/Riding-Battery-Low-8    	 3930580	       316 ns/op	     128 B/op	       5 allocs/op
PASS
ok  	dwf/controller/benchmarks	4.763s
*/

/*
// NOTE: NON Reliable test, because of changeing time.Now() frequently

// go test -bench=BenchmarkChangeStateNonVoidhandlersLogic -benchmem
// BenchmarkChangeStateNonVoidhandlersLogic to Benchmark non voidHandler logic
func BenchmarkChangeStateNonVoidhandlersLogic(b *testing.B) {

	tt := []struct {
		testname      string
		state         string
		expectedState string
		timeNow       time.Time
		BatteryLevel  int
	}{
		{
			testname:      "Ready-Unknown",
			state:         "Ready",
			expectedState: "Unknown",
			timeNow:       time.Now().Add(49 * time.Hour),
			BatteryLevel:  99,
		},
		{
			testname:      "Ready-Bounty",
			state:         "Ready",
			expectedState: "Bounty",
			timeNow:       time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 31, 0, 0, time.UTC),
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
			testname:      "Riding-Battery-Low",
			state:         "Riding",
			expectedState: "Battery-Low",
			BatteryLevel:  10,
		},
	}

	for _, tc := range tt {
		b.Run(tc.testname, func(b *testing.B) {
			for index := 0; index < b.N; index++ {
				func() {
					v, _ := controller.InitializeVehicle(context.Background(), "vec1", tc.state, tc.BatteryLevel)
					defer v.Terminate()
					if !tc.timeNow.IsZero() {
						patch := monkey.Patch(time.Now, func() time.Time { return tc.timeNow })
						defer patch.Unpatch()
					}

					time.Sleep(12 * time.Millisecond)

					if tc.expectedState != v.State() {
						b.Errorf("expected state %v; got %v", tc.expectedState, v.State())
						return
					}
				}()
			}
		})
	}
}

/*
go test -bench=BenchmarkChangeStateNonVoidhandlersLogic -benchmem
goos: linux
goarch: amd64
pkg: dwf/controller/benchmarks
BenchmarkChangeStateNonVoidhandlersLogic/Ready-Unknown-8         	     100	  12224754 ns/op	    1384 B/op	      30 allocs/op
BenchmarkChangeStateNonVoidhandlersLogic/Ready-Bounty-8          	     100	  12193880 ns/op	    1184 B/op	      25 allocs/op
BenchmarkChangeStateNonVoidhandlersLogic/UnknownVSBounty-8       	     100	  12166997 ns/op	    1308 B/op	      29 allocs/op
BenchmarkChangeStateNonVoidhandlersLogic/Riding-Battery-Low-8    	     100	  12091766 ns/op	     881 B/op	      16 allocs/op
PASS
ok  	dwf/controller/benchmarks	4.972s
*/
