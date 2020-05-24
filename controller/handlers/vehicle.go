package handlers

// IBatteryPercentage :
type IBatteryPercentage interface {
	BatteryPercentage() int
}

// BatteryLowHandler : handler check if battery low
func BatteryLowHandler(obj interface{}) bool {
	v, ok := obj.(IBatteryPercentage)
	if !ok {
		return false
	}
	if v.BatteryPercentage() < 20 {
		// logger.Info("Change State batteryLow handler")
		return true
	}
	return false
}
