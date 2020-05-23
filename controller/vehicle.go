package controller

import (
	"context"
	"fmt"
	"gwf/logger"
	"sync"
	"time"
)

// Vehicle : representation object for vehicle.
type Vehicle struct {
	ID                   string
	context              context.Context
	ctxCancelFunc        context.CancelFunc
	state                *state
	batteryPercentage    int
	lastDateStateChanged time.Time
	mu                   sync.Mutex
}

// InitializeVehicle new vehicle
func InitializeVehicle(ctx context.Context, id, state string, batteryPercentage int) (*Vehicle, error) {
	vctx, cancelFunction := context.WithCancel(ctx)
	v := &Vehicle{ID: id, context: vctx, ctxCancelFunc: cancelFunction}

	st, err := getState(state)
	if err != nil {
		return nil, err
	}
	v.state = st
	v.batteryPercentage = batteryPercentage
	v.lastDateStateChanged = time.Now()
	go v.autoStateChangerRunner()
	return v, nil
}

// ChangeState : change vehicle state
func (v *Vehicle) ChangeState(nextState string, usertype int) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	f, ns, err := v.state.validatechangeStatus(nextState, usertype)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("change statues from (%v) to (%v) is not valid ERROR:%v", v.state.name, nextState, err)
	}

	// execute the handler check if allowed to change vehicle state
	allowed := f(v)
	if !allowed {
		err := fmt.Errorf("not allowed to change statues from (%v) to (%v)", v.state.name, ns.name)
		logger.Error(err)
		return err
	}

	oldState := v.state.name
	v.state = ns
	v.lastDateStateChanged = time.Now()
	logger.Infof("Change Vehicle state from (%v) to (%v) by User(%v)", oldState, v.state.name, UsersType[usertype])
	return nil
}

func (v *Vehicle) autoStateChangerRunner() {
	for {
		select {
		case <-v.context.Done():
			//If context is cancelled, this case is selected
			logger.Info("Vehicle has been terminated ...")
			return
		case <-time.Tick(10 * time.Millisecond):
			for _, s := range v.state.autoStatesSorted {
				v.ChangeState(s.name, SYSTEM)
			}
		}
	}
}

// State : get the name of currant state and availabe States
func (v *Vehicle) State() string {
	return v.state.name
}

// Terminate : Terminate and remove this vehicle object from the system.
func (v *Vehicle) Terminate() {
	v.ctxCancelFunc()
}

// AvailableStates : get the name of currant state and availabe States
func (v *Vehicle) AvailableStates() (string, []string) {
	stateName := v.state.name
	availableStates := v.state.availableStates
	availableStatesArr := []string{}
	for _, v := range availableStates {
		availableStatesArr = append(availableStatesArr, v.name)
	}
	return stateName, availableStatesArr
}

// SetBatteryPercentage : set a new BatteryPercentage
func (v *Vehicle) SetBatteryPercentage(btr int) {
	v.batteryPercentage = btr
}

// AdminForceChangeState : function to force vehicle to change state.
// WARNING: this method will not check the user (authorization/session)
// the permission check should be done before call this method
func (v *Vehicle) AdminForceChangeState(nextState string, usertype int) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	if usertype != ADMIN { // check if the next user type available in the system
		err := fmt.Errorf("Non Admin User try to force change vehicle state (%v)", usertype)
		logger.Error(err)
		return err
	}
	logger.Warn("START FORCE VEHICLE to change state")

	ns, ok := internalStateList[nextState]
	if !ok { // check if the next state available in the system
		err := fmt.Errorf("State (%v) is not exist", nextState)
		logger.Error(err)
		return err
	}

	oldState := v.state.name
	v.state = ns
	v.lastDateStateChanged = time.Now()
	logger.Warnf("END FORCE VEHICLE to change state from (%v) to (%v) by user (%v)", oldState, v.state.name, UsersType[usertype])
	return nil
}
