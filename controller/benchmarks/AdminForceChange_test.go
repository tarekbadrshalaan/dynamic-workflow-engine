package benchmarks

import (
	"context"
	"gwf/controller"
	"testing"
)

//!+bench
// go test -bench=BenchmarkAdminForceChangeState -benchmem

func BenchmarkAdminForceChangeState(b *testing.B) {

	// Initialize only one vehicle to test stress changes
	v, _ := controller.InitializeVehicle(context.Background(), "vec1", "Ready", 99)
	defer v.Terminate()

	tt := []struct {
		testname   string
		adminState string
		usertype   int
	}{
		{
			testname:   "Ready-Bounty",
			adminState: "Bounty",
			usertype:   controller.ADMIN,
		},
		{
			testname:   "Bounty-Riding",
			adminState: "Riding",
			usertype:   controller.ADMIN,
		},
		{
			testname:   "Riding-Battery-Low",
			adminState: "Battery-Low",
			usertype:   controller.ADMIN,
		},
		{
			testname:   "Battery-Low-Unknown",
			adminState: "Unknown",
			usertype:   controller.ADMIN,
		},
		{
			testname:   "Unknown-Bounty",
			adminState: "Bounty",
			usertype:   controller.ADMIN,
		},
		{
			testname:   "Bounty-Dropped",
			adminState: "Dropped",
			usertype:   controller.ADMIN,
		},
		{
			testname:   "Dropped-Battery-Low",
			adminState: "Battery-Low",
			usertype:   controller.ADMIN,
		},
		{
			testname:   "Battery-Low-Ready",
			adminState: "Ready",
			usertype:   controller.ADMIN,
		},
	}
	for _, tc := range tt {
		b.Run(tc.testname, func(b *testing.B) {
			for index := 0; index < b.N; index++ {
				v.AdminForceChangeState(tc.adminState, tc.usertype)
			}
		})
	}
}

/*

go test -bench=BenchmarkAdminForceChangeState -benchmem
goos: linux
goarch: amd64
pkg: gwf/controller/benchmarks
BenchmarkAdminForceChangeState/Ready-Bounty-8         	 4508310	       229 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Bounty-Riding-8        	 5244963	       223 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Riding-Battery-Low-8   	 4951862	       229 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Battery-Low-Unknown-8  	 5198128	       246 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Unknown-Bounty-8       	 4417371	       230 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Bounty-Dropped-8       	 5210500	       225 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Dropped-Battery-Low-8  	 5300404	       239 ns/op	     112 B/op	       5 allocs/op
BenchmarkAdminForceChangeState/Battery-Low-Ready-8    	 5347808	       232 ns/op	     112 B/op	       5 allocs/op
PASS
ok  	gwf/controller/benchmarks	11.262s

*/

//!-bench
