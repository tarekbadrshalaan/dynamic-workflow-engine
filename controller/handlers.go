package controller

import (
	"gwf/logger"
	"time"
)

// handler :representation function for generic vehicle state handler
type handler func(v *Vehicle) bool

// handlersList : the per defined handlers list
// NOTE: we didn't use GO reflection because it might be insecure
// to expose internal memory to configuration environment.
var handlersList = map[string]handler{
	"voidHandler": voidHandler,
	"batteryLow":  batteryLowHandler,
	"after930PM":  after930PMHandler,
	"after48H":    after48HHandler,
}

// voidHandler : handler accept any chanage
func voidHandler(v *Vehicle) bool {
	return true
}

// batteryLowHandler : handler check if battery low
func batteryLowHandler(v *Vehicle) bool {
	if v.batteryPercentage < 20 {
		logger.Info("Change State batteryLow handler")
		return true
	}
	return false
}

// after930PMHandler : handler directly fire after 9:30 pm
func after930PMHandler(v *Vehicle) bool {
	now := time.Now()
	t930PM := time.Date(now.Year(), now.Month(), now.Day(), 21, 30, 0, 0, time.UTC)
	if now.After(t930PM) {
		logger.Info("Change State After 9:30 pm handler")
		return true
	}
	return false
}

// after48HHandler : handler directly fire after 48 Hours without a state change
func after48HHandler(v *Vehicle) bool {
	after48h := v.lastDateStateChanged.Add(time.Hour * time.Duration(48))
	now := time.Now()
	if now.After(after48h) {
		logger.Info("Change State from After 48 Hours handler")
		return true
	}
	return false
}

// TODO :

// IPrintExecute :
type IPrintExecute interface {
	print() bool
}

// printHandler :
func printHandler(obj interface{}) bool {
	p, ok := obj.(IPrintExecute)
	if !ok {
		return false
	}
	return p.print()
}
