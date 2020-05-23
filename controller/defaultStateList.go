package controller

// BuildDefaultStateList : get the per defined states list
func BuildDefaultStateList() {
	stateList := map[string]*state{}

	stateList["Ready"] = &state{
		name: "Ready",
		availableStates: map[string]availableState{
			"Riding": availableState{
				name:     "Riding",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{USER: true, HUNTER: true, ADMIN: true, SYSTEM: true},
			},
			"Unknown": availableState{
				name:     "Unknown",
				funcName: "after48H",
				autoRun:  true,
				priority: 0,
				users:    map[int]bool{SYSTEM: true},
			},
			"Bounty": availableState{
				name:     "Bounty",
				funcName: "after930PM",
				autoRun:  true,
				priority: 1,
				users:    map[int]bool{SYSTEM: true},
			},
		},
	}

	stateList["Riding"] = &state{
		name: "Riding",
		availableStates: map[string]availableState{
			"Ready": availableState{
				name:     "Ready",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{USER: true, HUNTER: true, ADMIN: true, SYSTEM: true},
			},
			"Battery-Low": availableState{
				name:     "Battery-Low",
				funcName: "batteryLow",
				autoRun:  true,
				priority: 0,
				users:    map[int]bool{VEHICLE: true, SYSTEM: true},
			},
		},
	}

	stateList["Battery-Low"] = &state{
		name: "Battery-Low",
		availableStates: map[string]availableState{
			"Bounty": availableState{
				name:     "Bounty",
				funcName: "voidHandler",
				autoRun:  true,
				users:    map[int]bool{VEHICLE: true, SYSTEM: true},
			},
		},
	}

	stateList["Bounty"] = &state{
		name: "Bounty",
		availableStates: map[string]availableState{
			"Collected": availableState{
				name:     "Collected",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{HUNTER: true, ADMIN: true, SYSTEM: true},
			},
		},
	}

	stateList["Collected"] = &state{
		name: "Collected",
		availableStates: map[string]availableState{
			"Dropped": availableState{
				name:     "Dropped",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{HUNTER: true, ADMIN: true, SYSTEM: true},
			},
		},
	}

	stateList["Dropped"] = &state{
		name: "Dropped",
		availableStates: map[string]availableState{
			"Ready": availableState{
				name:     "Ready",
				funcName: "voidHandler",
				autoRun:  false,
				users:    map[int]bool{HUNTER: true, ADMIN: true, SYSTEM: true},
			},
		},
	}

	// Unknown does not have Availabe state,
	// it should handled directly from the admin
	stateList["Unknown"] = &state{
		name: "Unknown",
	}

	// set internal states list.
	SetInternalStateList(stateList)
}
