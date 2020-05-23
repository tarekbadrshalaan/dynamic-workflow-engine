package controller

import (
	"dwf/logger"
	"fmt"
)

// State : representation object for State
type State struct {
	Name             string
	AvailableStates  map[string]availableState
	AutoStatesSorted []availableState
}

// AvailableState : one available state to move on from previous state.
type availableState struct {
	Name     string
	funcName string
	autoRun  bool
	priority int
	users    map[int]bool
}

// ValidatechangeStatus : check if the currant state can move to next_state with user_type
func (s *State) ValidatechangeStatus(nextState string, usertype int) (Handler, *State, error) {

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

	availableState, ok := s.AvailableStates[ns.Name]
	if !ok {
		err := fmt.Errorf("nextState(%v) is not exist", ns.Name)
		logger.Error(err)
		return nil, nil, err
	}

	if !availableState.users[usertype] {
		err := fmt.Errorf("Change vehicle state from (%v) to (%v) is not available for (%v user)", s.Name, ns.Name, user)
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

// GetState : get state by name
func GetState(state string) (*State, error) {
	s, ok := internalStateList[state]
	if !ok {
		err := fmt.Errorf("State (%v) is not exist", state)
		logger.Error(err)
		return nil, err
	}
	return s, nil
}
