package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"dwf/api/helper"
	"dwf/exvehicle"
	"dwf/logger"

	"github.com/julienschmidt/httprouter"
)

// Vehicles :
var Vehicles = struct {
	mu   sync.Mutex
	list map[string]*exvehicle.Vehicle
}{
	mu:   sync.Mutex{},
	list: map[string]*exvehicle.Vehicle{},
}

// VehicleAPI :
type VehicleAPI struct {
	ID                string   `json:"id"`
	State             string   `json:"state"`
	BatteryPercentage int      `json:"battery_percentage"`
	User              int      `json:"user"`
	AvailableStates   []string `json:"available_states"`
}

// ConfigVehicleRouter :
func ConfigVehicleRouter() []helper.Route {
	return []helper.Route{
		{Method: "GET", Path: "/vehicle/all_vehicles", Handle: allVehicles},
		{Method: "POST", Path: "/vehicle/initialize", Handle: initializeVehicle},
		{Method: "POST", Path: "/vehicle/change_state", Handle: changeStateVehicle},
		{Method: "POST", Path: "/vehicle/change_battery", Handle: changeBatteryVehicle},
		{Method: "GET", Path: "/vehicle/available_states/:id", Handle: availableStatesVehicle},
	}
}

func initializeVehicle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Vehicles.mu.Lock()
	defer Vehicles.mu.Unlock()

	v := &VehicleAPI{}
	if err := helper.ReadJSON(r, v); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	newVec, err := exvehicle.InitializeVehicle(context.Background(), v.ID, v.State, v.BatteryPercentage)
	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusInternalServerError)
		return
	}

	Vehicles.list[newVec.ID] = newVec
	s, ss := newVec.AvailableStates()

	helper.WriteResponseJSON(w, VehicleAPI{ID: newVec.ID, State: s, AvailableStates: ss, BatteryPercentage: newVec.BatteryPercentage()}, http.StatusCreated)
}

func allVehicles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Vehicles.mu.Lock()
	defer Vehicles.mu.Unlock()

	lsVehicles := []VehicleAPI{}
	for _, v := range Vehicles.list {
		s, ss := v.AvailableStates()
		lsVehicles = append(lsVehicles, VehicleAPI{ID: v.ID, State: s, AvailableStates: ss, BatteryPercentage: v.BatteryPercentage()})
	}

	helper.WriteResponseJSON(w, lsVehicles, http.StatusOK)
}

func changeStateVehicle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Vehicles.mu.Lock()
	defer Vehicles.mu.Unlock()

	v := &VehicleAPI{}
	if err := helper.ReadJSON(r, v); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}
	dbv, ok := Vehicles.list[v.ID]

	if !ok {
		msg := fmt.Errorf("Can’t find Vehicle with Id (%v)", v.ID)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}

	err := dbv.ChangeState(v.State, v.User)
	if err != nil {
		msg := fmt.Errorf("Can’t change to status (%v); err (%v)", v.State, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}
	s, ss := dbv.AvailableStates()
	helper.WriteResponseJSON(w, VehicleAPI{ID: dbv.ID, State: s, AvailableStates: ss, BatteryPercentage: dbv.BatteryPercentage()}, http.StatusOK)
}

func changeBatteryVehicle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Vehicles.mu.Lock()
	defer Vehicles.mu.Unlock()

	v := &VehicleAPI{}
	if err := helper.ReadJSON(r, v); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}
	dbv, ok := Vehicles.list[v.ID]

	if !ok {
		msg := fmt.Errorf("Can’t find Vehicle with Id (%v)", v.ID)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}

	dbv.SetBatteryPercentage(v.BatteryPercentage)
	helper.WriteResponseJSON(w, VehicleAPI{ID: dbv.ID, BatteryPercentage: dbv.BatteryPercentage()}, http.StatusOK)
}

func availableStatesVehicle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Vehicles.mu.Lock()
	defer Vehicles.mu.Unlock()

	requestID := ps.ByName("id")
	v, ok := Vehicles.list[requestID]

	if !ok {
		msg := fmt.Errorf("Can’t find Vehicle with Id (%v)", requestID)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}

	s, ss := v.AvailableStates()
	helper.WriteResponseJSON(w, VehicleAPI{ID: v.ID, State: s, AvailableStates: ss, BatteryPercentage: v.BatteryPercentage()}, http.StatusOK)
}
