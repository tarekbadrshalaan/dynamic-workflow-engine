package benchmarks

import (
	"context"
	"dew/controller"
	"testing"
)

//!+bench
// go test -bench=BenchmarkInitializeVehicle -benchmem
// go test -bench=BenchmarkInitializeVehicle -benchmem -benchtime=10000x

func BenchmarkInitializeVehicle(b *testing.B) {

	tt := []struct {
		testname          string
		id                string
		state             string
		batteryPercentage int
	}{
		{
			testname:          "Ready",
			id:                "vec1",
			state:             "Ready",
			batteryPercentage: 50,
		},
		{
			testname:          "Riding",
			id:                "vec1",
			state:             "Riding",
			batteryPercentage: 70,
		},
		{
			testname:          "Battery-Low",
			id:                "vec1",
			state:             "Battery-Low",
			batteryPercentage: 10,
		},
		{
			testname:          "Bounty",
			id:                "vec1",
			state:             "Bounty",
			batteryPercentage: 3,
		},
		{
			testname:          "Collected",
			id:                "vec1",
			state:             "Collected",
			batteryPercentage: 0,
		},
		{
			testname:          "Dropped",
			id:                "vec1",
			state:             "Dropped",
			batteryPercentage: 99,
		},
		{
			testname:          "Unknown",
			id:                "vec1",
			state:             "Unknown",
			batteryPercentage: 99,
		},
	}

	for _, tc := range tt {
		b.Run(tc.testname, func(b *testing.B) {
			for index := 0; index < b.N; index++ {
				v, _ := controller.InitializeVehicle(context.Background(), tc.id, tc.state, tc.batteryPercentage)
				v.Terminate()
			}
		})
	}

}

/* result from go1.12
go test -bench=BenchmarkInitializeVehicle -benchmem
goos: linux
goarch: amd64
pkg: dew/controller/benchmarks
BenchmarkInitializeVehicle/Ready-8         	 1000000	      1412 ns/op	     451 B/op	       7 allocs/op
BenchmarkInitializeVehicle/Riding-8        	 1000000	      1713 ns/op	     444 B/op	       6 allocs/op
BenchmarkInitializeVehicle/Battery-Low-8   	 1000000	      1835 ns/op	     450 B/op	       7 allocs/op
BenchmarkInitializeVehicle/Bounty-8        	 1000000	      1667 ns/op	     434 B/op	       6 allocs/op
BenchmarkInitializeVehicle/Collected-8     	 1000000	      1835 ns/op	     442 B/op	       6 allocs/op
BenchmarkInitializeVehicle/Dropped-8       	 1000000	      1564 ns/op	     455 B/op	       6 allocs/op
BenchmarkInitializeVehicle/Unknown-8       	 1000000	      1495 ns/op	     430 B/op	       6 allocs/op
PASS
ok  	dew/controller/benchmarks	86.777s
*/

/* result from go1.14
// NOTE: it will stay running forever without add -benchtime=10000x
go test -bench=BenchmarkInitializeVehicle -benchmem -benchtime=10000x
goos: linux
goarch: amd64
pkg: dew/controller/benchmarks
BenchmarkInitializeVehicle/Ready-8         	   10000	       729 ns/op	     446 B/op	       7 allocs/op
BenchmarkInitializeVehicle/Riding-8        	   10000	       738 ns/op	     441 B/op	       7 allocs/op
BenchmarkInitializeVehicle/Battery-Low-8   	   10000	       665 ns/op	     432 B/op	       7 allocs/op
BenchmarkInitializeVehicle/Bounty-8        	   10000	       540 ns/op	     450 B/op	       7 allocs/op
BenchmarkInitializeVehicle/Collected-8     	   10000	       525 ns/op	     423 B/op	       6 allocs/op
BenchmarkInitializeVehicle/Dropped-8       	   10000	       539 ns/op	     440 B/op	       6 allocs/op
BenchmarkInitializeVehicle/Unknown-8       	   10000	       525 ns/op	     462 B/op	       7 allocs/op
PASS
ok  	dew/controller/benchmarks	0.172s
*/

//!-bench
