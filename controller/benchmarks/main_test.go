package benchmarks

import (
	"gwf/controller"
	"gwf/logger"
	"os"
	"testing"
)

//!+bench
//go test -bench . -benchmem

func TestMain(m *testing.M) {
	// here you can add general code before run test cases ...

	/* logger initialize start */
	mylogger := logger.NewEmptyLogger()
	logger.InitializeLogger(&mylogger)
	defer logger.Close()
	/* logger initialize end */

	// build state list from predefined one
	controller.BuildDefaultStateList()

	code := m.Run()
	os.Exit(code)
}

//!-bench

/*
go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: gwf/controller/benchmarks
BenchmarkAdminForceChangeState/Ready-Bounty-8         	 5025529	       224 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Bounty-Riding-8        	 5444305	       241 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Riding-Battery-Low-8   	 5297530	       226 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Battery-Low-Unknown-8  	 5153222	       233 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Unknown-Bounty-8       	 5291767	       231 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Bounty-Dropped-8       	 4845734	       241 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Dropped-Battery-Low-8  	 5299237	       226 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Battery-Low-Ready-8    	 4798034	       241 ns/op	     112 B/op	       5 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Ready-8    	 1911896	       646 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Riding-8   	 1813723	       613 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Battery-Low-8         	 1944885	       617 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Bounty-8              	 1912375	       622 ns/op	     272 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Collected-8           	 1937920	       611 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateVoidhandlersUser/state-Dropped-8             	 1778307	       618 ns/op	     240 B/op	       9 allocs/op
BenchmarkChangeStateNonVoidhandlersUser/Ready-Unknown-8          	 2808348	       392 ns/op	     128 B/op	       5 allocs/op
BenchmarkChangeStateNonVoidhandlersUser/Ready-Bounty-8           	 2377176	       468 ns/op	     128 B/op	       5 allocs/op
BenchmarkChangeStateNonVoidhandlersUser/Riding-Battery-Low-8     	 3792481	       317 ns/op	     128 B/op	       5 allocs/op

// BenchmarkChangeStateNonVoidhandlersLogic/Ready-Unknown-8         	     100	  12083515 ns/op	    1243 B/op	      29 allocs/op
// BenchmarkChangeStateNonVoidhandlersLogic/Ready-Bounty-8          	     100	  12080946 ns/op	    1178 B/op	      25 allocs/op
// BenchmarkChangeStateNonVoidhandlersLogic/UnknownVSBounty-8       	     100	  12085870 ns/op	    1250 B/op	      29 allocs/op
// BenchmarkChangeStateNonVoidhandlersLogic/Riding-Battery-Low-8    	     100	  12053285 ns/op	     867 B/op	      16 allocs/op

// BenchmarkInitializeVehicle/Ready-8                               	 1000000	      1355 ns/op	     450 B/op	       7 allocs/op
// BenchmarkInitializeVehicle/Riding-8                              	 1000000	      1655 ns/op	     446 B/op	       7 allocs/op
// BenchmarkInitializeVehicle/Battery-Low-8                         	 1000000	      1989 ns/op	     447 B/op	       7 allocs/op
// BenchmarkInitializeVehicle/Bounty-8                              	 1000000	      1514 ns/op	     441 B/op	       6 allocs/op
// BenchmarkInitializeVehicle/Collected-8                           	 1000000	      1462 ns/op	     452 B/op	       6 allocs/op
// BenchmarkInitializeVehicle/Dropped-8                             	 1000000	      1556 ns/op	     434 B/op	       6 allocs/op
// BenchmarkInitializeVehicle/Unknown-8                             	 1000000	      1587 ns/op	     438 B/op	       6 allocs/op

PASS
ok  	gwf/controller/benchmarks	28.050s

*/
