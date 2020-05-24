package main

import (
	"dwf/api"
	"dwf/config"
	"dwf/controller"
	"dwf/logger"
	"fmt"
	"log"
	"net/http"
)

func main() {
	/* logger initialize start */
	mylogger := logger.NewZapLogger()
	logger.InitializeLogger(&mylogger)
	defer logger.Close()
	/* logger initialize end */

	/* controller initialize start */
	c := &controller.Configuration{}
	// config.Configuration("exatm/controller_config.json", c)
	config.Configuration("exvehicle/controller_config.json", c)
	err := controller.BuildStates(c)
	if err != nil {
		logger.Fatal(err)
	}
	/* controller initialize end */

	/* initialize webserver start */
	r := api.NewRouter()
	addr := fmt.Sprintf("%v:%d", c.WebAddress, c.WebPort)
	log.Fatal(http.ListenAndServe(addr, r))
	/* initialize webserver end */
}
