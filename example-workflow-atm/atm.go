package atm

import (
	"context"
	"dwf/controller"
	"dwf/logger"
	"fmt"
	"sync"
	"time"
)

// Card : representation object for Card.
type Card struct {
	ID            string
	context       context.Context
	ctxCancelFunc context.CancelFunc
	state         *controller.State
	password      string
	balance       float64
	mu            sync.Mutex
}

// InitializeCard new Card
func InitializeCard(ctx context.Context, id, state string, password string, balance float64) (*Card, error) {
	cctx, cancelFunction := context.WithCancel(ctx)
	c := &Card{ID: id, context: cctx, ctxCancelFunc: cancelFunction, password: password, balance: balance}

	st, err := controller.GetState(state)
	if err != nil {
		return nil, err
	}
	c.state = st
	go c.autoStateChangerRunner()
	return c, nil
}

// ChangeState : change Card state
func (c *Card) ChangeState(nextState string, usertype int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	f, ns, err := c.state.ValidatechangeStatus(nextState, usertype)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("change statues from (%v) to (%v) is not valid ERROR:%v", c.state.Name, nextState, err)
	}

	// execute the handler check if allowed to change Card state
	allowed := f(c)
	if !allowed {
		err := fmt.Errorf("not allowed to change statues from (%v) to (%v)", c.state.Name, ns.Name)
		logger.Error(err)
		return err
	}

	oldState := c.state.Name
	c.state = ns
	logger.Infof("Change Card state from (%v) to (%v) by User(%v)", oldState, c.state.Name, controller.UsersType[usertype])
	return nil
}

func (c *Card) autoStateChangerRunner() {
	for {
		select {
		case <-c.context.Done():
			//If context is cancelled, this case is selected
			logger.Info("Card has been terminated ...")
			return
		case <-time.Tick(10 * time.Millisecond):
			for _, s := range c.state.AutoStatesSorted {
				c.ChangeState(s.Name, controller.SYSTEM)
			}
		}
	}
}

// State : get the name of currant state and availabe States
func (c *Card) State() string {
	return c.state.Name
}

// Terminate : Terminate and remove this Card object from the system.
func (c *Card) Terminate() {
	c.ctxCancelFunc()
}

// AvailableStates : get the name of currant state and availabe States
func (c *Card) AvailableStates() (string, []string) {
	stateName := c.state.Name
	availableStates := c.state.AvailableStates
	availableStatesArr := []string{}
	for _, v := range availableStates {
		availableStatesArr = append(availableStatesArr, v.Name)
	}
	return stateName, availableStatesArr
}

// Print :
func (c *Card) Print() bool {
	fmt.Println("============== hello from print status :) =====================")
	return true
}
