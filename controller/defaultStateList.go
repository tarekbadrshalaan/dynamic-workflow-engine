package controller

// BuildDefaultStateList : get the per defined states list
func BuildDefaultStateList() {
	stateList := map[string]*State{}

	stateList["Ready"] = &State{
		Name: "Ready",
		AvailableStates: map[string]availableState{
			"Riding": availableState{
				Name:     "Riding",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{USER: true, HUNTER: true, ADMIN: true, SYSTEM: true},
			},
			"Unknown": availableState{
				Name:     "Unknown",
				funcName: "after48H",
				autoRun:  true,
				priority: 0,
				users:    map[int]bool{SYSTEM: true},
			},
			"Bounty": availableState{
				Name:     "Bounty",
				funcName: "after930PM",
				autoRun:  true,
				priority: 1,
				users:    map[int]bool{SYSTEM: true},
			},
		},
	}

	stateList["Riding"] = &State{
		Name: "Riding",
		AvailableStates: map[string]availableState{
			"Ready": availableState{
				Name:     "Ready",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{USER: true, HUNTER: true, ADMIN: true, SYSTEM: true},
			},
			"Battery-Low": availableState{
				Name:     "Battery-Low",
				funcName: "batteryLow",
				autoRun:  true,
				priority: 0,
				users:    map[int]bool{VEHICLE: true, SYSTEM: true},
			},
		},
	}

	stateList["Battery-Low"] = &State{
		Name: "Battery-Low",
		AvailableStates: map[string]availableState{
			"Bounty": availableState{
				Name:     "Bounty",
				funcName: "voidHandler",
				autoRun:  true,
				users:    map[int]bool{VEHICLE: true, SYSTEM: true},
			},
		},
	}

	stateList["Bounty"] = &State{
		Name: "Bounty",
		AvailableStates: map[string]availableState{
			"Collected": availableState{
				Name:     "Collected",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{HUNTER: true, ADMIN: true, SYSTEM: true},
			},
		},
	}

	stateList["Collected"] = &State{
		Name: "Collected",
		AvailableStates: map[string]availableState{
			"Dropped": availableState{
				Name:     "Dropped",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{HUNTER: true, ADMIN: true, SYSTEM: true},
			},
		},
	}

	stateList["Dropped"] = &State{
		Name: "Dropped",
		AvailableStates: map[string]availableState{
			"Ready": availableState{
				Name:     "Ready",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{HUNTER: true, ADMIN: true, SYSTEM: true},
			},
		},
	}

	// Unknown does not have Availabe state,
	// it should handled directly from the admin
	stateList["Unknown"] = &State{
		Name: "Unknown",
	}

	// set internal states list.
	SetInternalStateList(stateList)
}
