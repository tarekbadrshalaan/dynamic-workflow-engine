package config

import (
	"github.com/tarekbadrshalaan/goStuff/configuration"
)

// Configuration : get configuration for any type, it should be json file
func Configuration(path string, config interface{}) {
	err := configuration.JSON(path, config)
	if err != nil {
		panic(err)
	}
}
