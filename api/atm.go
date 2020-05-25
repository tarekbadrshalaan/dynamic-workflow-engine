package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"dwf/api/helper"
	"dwf/exatm"
	"dwf/logger"

	"github.com/julienschmidt/httprouter"
)

// ATMCards :
var ATMCards = struct {
	mu   sync.Mutex
	list map[string]*exatm.Card
}{
	mu:   sync.Mutex{},
	list: map[string]*exatm.Card{},
}

// CardAPI :
type CardAPI struct {
	ID              string   `json:"id"`
	State           string   `json:"state"`
	AvailableStates []string `json:"available_states"`
	Password        string   `json:"password"`
	InputPin        string   `json:"input_pin"`
	Balance         float64  `json:"Balance"`
	RequestFund     float64  `json:"request_fund"`
}

// ConfigATMRouter :
func ConfigATMRouter() []helper.Route {
	return []helper.Route{
		{Method: "GET", Path: "/atm/all_cards", Handle: allCards},
		{Method: "POST", Path: "/atm/initialize", Handle: initializeATM},
		{Method: "POST", Path: "/atm/change_state", Handle: changeStateATM},
		{Method: "GET", Path: "/atm/available_states/:id", Handle: availableStatesATM},
		{Method: "POST", Path: "/atm/insert_pin_Code", Handle: insertPinCode},
		{Method: "POST", Path: "/atm/request_fund", Handle: requestFund},
	}
}

func allCards(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ATMCards.mu.Lock()
	defer ATMCards.mu.Unlock()

	lsATMCards := []CardAPI{}
	for _, c := range ATMCards.list {
		s, ss := c.AvailableStates()
		lsATMCards = append(lsATMCards, CardAPI{ID: c.ID, State: s, AvailableStates: ss, Balance: c.GetBalance()})
	}

	helper.WriteResponseJSON(w, lsATMCards, http.StatusOK)
}

func initializeATM(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ATMCards.mu.Lock()
	defer ATMCards.mu.Unlock()

	c := &CardAPI{}
	if err := helper.ReadJSON(r, c); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	newCard, err := exatm.InitializeCard(context.Background(), c.ID, c.State, c.Password, c.Balance)
	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusInternalServerError)
		return
	}

	ATMCards.list[newCard.ID] = newCard
	s, ss := newCard.AvailableStates()

	helper.WriteResponseJSON(w, CardAPI{ID: newCard.ID, State: s, AvailableStates: ss, Balance: newCard.GetBalance()}, http.StatusCreated)

}

func availableStatesATM(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ATMCards.mu.Lock()
	defer ATMCards.mu.Unlock()

	requestID := ps.ByName("id")
	c, ok := ATMCards.list[requestID]

	if !ok {
		msg := fmt.Errorf("Can’t find ATMCard with Id (%v)", requestID)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}

	s, ss := c.AvailableStates()
	helper.WriteResponseJSON(w, CardAPI{ID: c.ID, State: s, AvailableStates: ss, Balance: c.GetBalance()}, http.StatusCreated)
}

func changeStateATM(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ATMCards.mu.Lock()
	defer ATMCards.mu.Unlock()

	c := &CardAPI{}
	if err := helper.ReadJSON(r, c); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}
	dbc, ok := ATMCards.list[c.ID]

	if !ok {
		msg := fmt.Errorf("Can’t find card with Id (%v)", c.ID)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}

	err := dbc.ChangeState(c.State, 0)
	if err != nil {
		msg := fmt.Errorf("Can’t change to status (%v); err (%v)", c.State, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}
	s, ss := dbc.AvailableStates()
	helper.WriteResponseJSON(w, CardAPI{ID: dbc.ID, State: s, AvailableStates: ss, Balance: dbc.GetBalance()}, http.StatusCreated)
}

func insertPinCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ATMCards.mu.Lock()
	defer ATMCards.mu.Unlock()

	c := &CardAPI{}
	if err := helper.ReadJSON(r, c); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	dbc, ok := ATMCards.list[c.ID]
	if !ok {
		msg := fmt.Errorf("Can’t find card with Id (%v)", c.ID)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}

	dbc.InputPin = c.Password
	helper.WriteResponseJSON(w, CardAPI{ID: dbc.ID, Balance: dbc.GetBalance()}, http.StatusCreated)
}

func requestFund(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ATMCards.mu.Lock()
	defer ATMCards.mu.Unlock()

	c := &CardAPI{}
	if err := helper.ReadJSON(r, c); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	dbc, ok := ATMCards.list[c.ID]
	if !ok {
		msg := fmt.Errorf("Can’t find card with Id (%v)", c.ID)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}

	dbc.RequestFund = c.RequestFund
	helper.WriteResponseJSON(w, CardAPI{ID: dbc.ID, Balance: dbc.GetBalance()}, http.StatusCreated)
}
