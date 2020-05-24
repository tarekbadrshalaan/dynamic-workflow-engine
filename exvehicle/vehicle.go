package exvehicle

import (
	"context"
	"dwf/controller"
	"dwf/logger"
	"fmt"
	"sync"
	"time"
)

// Vehicle : representation object for vehicle.
type Vehicle struct {
	ID                   string
	context              context.Context
	ctxCancelFunc        context.CancelFunc
	state                *controller.State
	batteryPercentage    int
	lastDateStateChanged time.Time
	mu                   sync.Mutex
}

// InitializeVehicle new vehicle
func InitializeVehicle(ctx context.Context, id, state string, batteryPercentage int) (*Vehicle, error) {
	vctx, cancelFunction := context.WithCancel(ctx)
	v := &Vehicle{ID: id, context: vctx, ctxCancelFunc: cancelFunction}

	st, err := controller.GetState(state)
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
	f, ns, err := v.state.ValidatechangeStatus(nextState, usertype)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("change statues from (%v) to (%v) is not valid ERROR:%v", v.state.Name, nextState, err)
	}

	// execute the handler check if allowed to change vehicle state
	allowed := f(v)
	if !allowed {
		err := fmt.Errorf("not allowed to change statues from (%v) to (%v)", v.state.Name, ns.Name)
		// logger.Error(err)
		return err
	}

	oldState := v.state.Name
	v.state = ns
	v.lastDateStateChanged = time.Now()
	logger.Infof("Change Vehicle state from (%v) to (%v) by User(%v)", oldState, v.state.Name, controller.UsersType[usertype])
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
			for _, s := range v.state.AutoStatesSorted {
				v.ChangeState(s.Name, controller.SYSTEM)
			}
		}
	}
}

// State : get the name of currant state and availabe States
func (v *Vehicle) State() string {
	return v.state.Name
}

// Terminate : Terminate and remove this vehicle object from the system.
func (v *Vehicle) Terminate() {
	v.ctxCancelFunc()
}

// AvailableStates : get the name of currant state and availabe States
func (v *Vehicle) AvailableStates() (string, []string) {
	stateName := v.state.Name
	availableStates := v.state.AvailableStates
	availableStatesArr := []string{}
	for _, v := range availableStates {
		availableStatesArr = append(availableStatesArr, v.Name)
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

	if usertype != controller.ADMIN { // check if the next user type available in the system
		err := fmt.Errorf("Non Admin User try to force change vehicle state (%v)", usertype)
		logger.Error(err)
		return err
	}
	logger.Warn("START FORCE VEHICLE to change state")

	ns, err := controller.GetInternalStateList(nextState)
	if err != nil { // check if the next state available in the system
		err := fmt.Errorf("State (%v) is not exist", nextState)
		logger.Error(err)
		return err
	}

	oldState := v.state.Name
	v.state = ns
	v.lastDateStateChanged = time.Now()
	logger.Warnf("END FORCE VEHICLE to change state from (%v) to (%v) by user (%v)", oldState, v.state.Name, controller.UsersType[usertype])
	return nil
}

// Print :
func (v *Vehicle) Print() bool {
	fmt.Println("============== hello from print status :) =====================")
	return true
}

// BatteryPercentage :
func (v *Vehicle) BatteryPercentage() int {
	return v.batteryPercentage
}

// LastDateStateChanged :
func (v *Vehicle) LastDateStateChanged() time.Time {
	return v.lastDateStateChanged
}
