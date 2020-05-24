package controller

import (
	"dwf/controller/handlers"
)

// Handler :representation function for generic vehicle state handler
type Handler func(obj interface{}) bool

// handlersList : the per defined handlers list
// NOTE: we didn't use GO reflection because it might be insecure
// to expose internal memory to configuration environment.
var handlersList = map[string]Handler{
	"voidHandler":               handlers.VoidHandler,
	"batteryLow":                handlers.BatteryLowHandler,
	"after930PM":                handlers.After930PMHandler,
	"after48H":                  handlers.After48HHandler,
	"isValidCardHandler":        handlers.IsValidCardHandler,
	"notValidHandler":           handlers.NotValidHandler,
	"validCardHandler":          handlers.ValidCardHandler,
	"printNotValidHandler":      handlers.PrintNotValidHandler,
	"showPinScreenHandler":      handlers.ShowPinScreenHandler,
	"validatePinHandler":        handlers.ValidatePinHandler,
	"validPinHandler":           handlers.ValidPinHandler,
	"inValidPinHandler":         handlers.InValidPinHandler,
	"chooseActionHandler":       handlers.ChooseActionHandler,
	"chooseMoneyHandler":        handlers.ChooseMoneyHandler,
	"showAccountBalanceHandler": handlers.ShowAccountBalanceHandler,
	"showMoneyScreenHandler":    handlers.ShowMoneyScreenHandler,
	"sufficientFundHandler":     handlers.SufficientFundHandler,
	"relaseMoneyHandler":        handlers.RelaseMoneyHandler,
	"endHandler":                handlers.EndHandler,
}
