package controller

const (
	// USER : End­ users ­ Regular app­users / riders.
	USER = iota
	// HUNTER : End users who have signed up to be chargers of vehicles and are responsible for picking up low battery vehicles.
	HUNTER
	// ADMIN : Super users who can do everything
	ADMIN
	// VEHICLE : User to allow vehicle create state changes
	VEHICLE
	// SYSTEM : User to allow the backend system create state changes
	SYSTEM
	// add more user type here ...
)

// UsersType : all users types represented in string value
var UsersType = map[int]string{
	USER:    "User",
	HUNTER:  "Hunter",
	ADMIN:   "Admin",
	VEHICLE: "Vehicle",
	SYSTEM:  "System",
}

// UsersTypeStr :
var UsersTypeStr = map[string]int{
	"User":    USER,
	"Hunter":  HUNTER,
	"Admin":   ADMIN,
	"Vehicle": VEHICLE,
	"System":  SYSTEM,
}
