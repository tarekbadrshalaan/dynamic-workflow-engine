package handlers

import (
	"dwf/logger"
	"time"
)

// After930PMHandler : handler directly fire after 9:30 pm
func After930PMHandler(obj interface{}) bool {
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

// After48HHandler : handler directly fire after 48 Hours without a state change
func After48HHandler(obj interface{}) bool {
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
