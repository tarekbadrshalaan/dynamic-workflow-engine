package api

import (
	"net/http"

	"dwf/api/helper"

	"github.com/julienschmidt/httprouter"
)

// ConfigATMRouter :
func ConfigATMRouter() []helper.Route {
	return []helper.Route{
		{Method: "POST", Path: "/atm/initialize", Handle: initializeATM},
		{Method: "POST", Path: "/atm/change_state", Handle: changeStateATM},
		{Method: "GET", Path: "/atm/available_states/:id", Handle: availableStatesATM},
	}
}

func initializeATM(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func availableStatesATM(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	requestID := ps.ByName("id")
	helper.WriteResponseJSON(w, requestID, http.StatusOK)
}

func changeStateATM(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
