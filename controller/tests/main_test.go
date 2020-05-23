package tests

import (
	"gwf/controller"
	"gwf/logger"
	"os"
	"testing"
)

//!+test
// go test -v
// go test -v -bench . -benchmem

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

//!-tests
