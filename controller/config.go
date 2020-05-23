package controller

import (
	"dwf/logger"
	"fmt"
	"sort"
)

// internalStateList :
var internalStateList map[string]*State

// Configuration : controller configuration
type Configuration struct {
	StatesList []StateConfig `json:"states_list"`
}

// StateConfig :
type StateConfig struct {
	Name                 string                 `json:"name"`
	AvailableStateConfig []AvailableStateConfig `json:"available_states"`
}

// AvailableStateConfig :
type AvailableStateConfig struct {
	Name     string   `json:"name"`
	FuncName string   `json:"function"`
	AutoRun  bool     `json:"auto_run"`
	Priority int      `json:"priority"`
	Users    []string `json:"users"`
}

// BuildStates : build states from configuration
func BuildStates(conf *Configuration) error {
	states := map[string]*State{}

	for _, s := range conf.StatesList {
		states[s.Name] = &State{
			Name:            s.Name,
			AvailableStates: map[string]availableState{},
		}
		for _, a := range s.AvailableStateConfig {
			avs := availableState{
				Name:     a.Name,
				autoRun:  a.AutoRun,
				priority: a.Priority,
				users:    map[int]bool{},
			}
			_, ok := handlersList[a.FuncName]
			if !ok {
				err := fmt.Errorf("configration Function (%v) is not exist", a.FuncName)
				logger.Error(err)
				return err
			}
			avs.funcName = a.FuncName

			for _, u := range a.Users {
				inuser, ok := UsersTypeStr[u]
				if !ok {
					err := fmt.Errorf("configration User (%v) is not exist", u)
					logger.Error(err)
					return err
				}
				avs.users[inuser] = true
			}

			states[s.Name].AvailableStates[a.Name] = avs
		}
	}
	// set internal states list.
	SetInternalStateList(states)
	return nil
}

// SetInternalStateList :
func SetInternalStateList(states map[string]*State) {

	for _, s := range states {
		autoStates := []availableState{}
		for _, s := range s.AvailableStates {
			if s.autoRun {
				autoStates = append(autoStates, s)
			}
		}
		sort.SliceStable(autoStates, func(i, j int) bool {
			return autoStates[i].priority < autoStates[j].priority
		})
		s.AutoStatesSorted = autoStates
	}

	// set internal states list.
	internalStateList = states
}
