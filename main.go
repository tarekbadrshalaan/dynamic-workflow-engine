package main

import (
	"dwf/config"
	"dwf/controller"
	"dwf/logger"
	"dwf/playground"
)

func main() {
	/* logger initialize start */
	mylogger := logger.NewZapLogger()
	logger.InitializeLogger(&mylogger)
	defer logger.Close()
	/* logger initialize end */

	/* controller initialize start */

	// a> build state list from predefined one
	// controller.BuildDefaultStateList()

	// b> use json configuration file to build state list
	confstateList := &controller.Configuration{}
	config.Configuration("controller_config.json", confstateList)
	err := controller.BuildStates(confstateList)
	if err != nil {
		logger.Fatal(err)
	}
	/* controller initialize end */

	vi, err := playground.Dynamic()
	if err != nil {
		logger.Error(err)
	}
	viState := vi.State()
	logger.Infof("Vehicle id (%v) the final state (%v)", vi.ID, viState)

	// vi, err = playground.PlayFullVehicleStateCycle()
	// if err != nil {
	// 	logger.Error(err)
	// }
	// viState = vi.State()
	// logger.Infof("Vehicle id (%v) the final state (%v)", vi.ID, viState)

	// vi, err = playground.PlayReadyBounty()
	// if err != nil {
	// 	logger.Error(err)
	// }
	// viState = vi.State()
	// logger.Infof("Vehicle id (%v) the final state (%v)", vi.ID, viState)

	// vi, err = playground.PlayReadyUnknown()
	// if err != nil {
	// 	logger.Error(err)
	// }
	// viState = vi.State()
	// logger.Infof("Vehicle id (%v) the final state (%v)", vi.ID, viState)

	// vi, err = playground.PlayInvalidUser()
	// if err != nil {
	// 	logger.Error(err)
	// }
	// viState = vi.State()
	// logger.Infof("Vehicle id (%v) the final state (%v)", vi.ID, viState)

}
