package controller

import (
	"fmt"
	"gwf/logger"
)

// state : representation object for state
type state struct {
	name             string
	availableStates  map[string]availableState
	autoStatesSorted []availableState
}

// AvailableState : one available state to move on from previous state.
type availableState struct {
	name     string
	funcName string
	autoRun  bool
	priority int
	users    map[int]bool
}

// validatechangeStatus : check if the currant state can move to next_state with user_type
func (s *state) validatechangeStatus(nextState string, usertype int) (handler, *state, error) {

	ns, ok := internalStateList[nextState]
	if !ok { // check if the next state available in the system
		err := fmt.Errorf("State (%v) is not exist", nextState)
		logger.Error(err)
		return nil, nil, err
	}
	user, ok := UsersType[usertype]
	if !ok { // check if the next user type available in the system
		err := fmt.Errorf("User type (%v) is not exist", usertype)
		logger.Error(err)
		return nil, nil, err
	}

	availableState, ok := s.availableStates[ns.name]
	if !ok {
		err := fmt.Errorf("nextState(%v) is not exist", ns.name)
		logger.Error(err)
		return nil, nil, err
	}

	if !availableState.users[usertype] {
		err := fmt.Errorf("Change vehicle state from (%v) to (%v) is not available for (%v user)", s.name, ns.name, user)
		logger.Error(err)
		return nil, nil, err
	}

	f, ok := handlersList[availableState.funcName]
	if !ok {
		err := fmt.Errorf("FuncName(%v) is not exist", availableState.funcName)
		logger.Error(err)
		return nil, nil, err
	}
	return f, ns, nil
}

// getState : get state by name
func getState(state string) (*state, error) {
	s, ok := internalStateList[state]
	if !ok {
		err := fmt.Errorf("State (%v) is not exist", state)
		logger.Error(err)
		return nil, err
	}
	return s, nil
}
