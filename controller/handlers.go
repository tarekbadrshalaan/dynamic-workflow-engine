package controller

import (
	"dwf/logger"
	"time"
)

// Handler :representation function for generic vehicle state handler
type Handler func(obj interface{}) bool

// handlersList : the per defined handlers list
// NOTE: we didn't use GO reflection because it might be insecure
// to expose internal memory to configuration environment.
var handlersList = map[string]Handler{
	"voidHandler":  voidHandler,
	"batteryLow":   batteryLowHandler,
	"after930PM":   after930PMHandler,
	"after48H":     after48HHandler,
	"printHandler": printHandler,
}

// voidHandler : handler accept any chanage
func voidHandler(obj interface{}) bool {
	return true
}

// IBatteryPercentage :
type IBatteryPercentage interface {
	BatteryPercentage() int
}

// batteryLowHandler : handler check if battery low
func batteryLowHandler(obj interface{}) bool {
	v, ok := obj.(IBatteryPercentage)
	if !ok {
		return false
	}
	if v.BatteryPercentage() < 20 {
		logger.Info("Change State batteryLow handler")
		return true
	}
	return false
}

// after930PMHandler : handler directly fire after 9:30 pm
func after930PMHandler(obj interface{}) bool {
	now := time.Now()
	t930PM := time.Date(now.Year(), now.Month(), now.Day(), 21, 30, 0, 0, time.UTC)
	if now.After(t930PM) {
		logger.Info("Change State After 9:30 pm handler")
		return true
	}
	return false
}

// IlastDateStateChanged :
type IlastDateStateChanged interface {
	LastDateStateChanged() time.Time
}

// after48HHandler : handler directly fire after 48 Hours without a state change
func after48HHandler(obj interface{}) bool {
	v, ok := obj.(IlastDateStateChanged)
	if !ok {
		return false
	}
	after48h := v.LastDateStateChanged().Add(time.Hour * time.Duration(48))
	now := time.Now()
	if now.After(after48h) {
		logger.Info("Change State from After 48 Hours handler")
		return true
	}
	return false
}

// IPrintExecute :
type IPrintExecute interface {
	Print() bool
}

// PtHandler :
func printHandler(obj interface{}) bool {
	p, ok := obj.(IPrintExecute)
	if !ok {
		return false
	}
	return p.Print()
}
