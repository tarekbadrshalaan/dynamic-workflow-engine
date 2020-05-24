package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewRouter :creates a new router instance and iterate through all the Routes to get each’s
// Route’s Method, Pattern and Handle and registers a new request handle.
func NewRouter() http.Handler {
	router := httprouter.New()

	for _, r := range ConfigATMRouter() {
		router.Handle(r.Method, r.Path, r.Handler())
	}

	for _, r := range ConfigVehicleRouter() {
		router.Handle(r.Method, r.Path, r.Handler())
	}

	return router
}
