# Dynamic Workflow Engine (DWE) 🚀

## Introduction
Dynamic Workflow Engine (DWE) is a Proof-of-Concept (PoC) implemented in Go, designed to provide a highly configurable workflow engine. You can easily tailor DWE for various use-cases, as illustrated in the `exATM` and `exVehicle` directories.

## Features 🌟
- **Controller**: The core package for workflow orchestration.
- **API**: Exposes functionalities for easy integration.
- **Config**: Handles configuration via JSON files.
- **Logger**: Provides three types of logging implementations.
- **Examples**: Includes example implementations like `exATM` and `exVehicle`.

## Getting Started 🛠️

### Installation
1. Clone the repository.
2. Navigate to the project directory.
3. Run `go install` to install the necessary packages.

### Usage
1. Import the `controller` package into your soft real-time API solution.
2. Initialize the workflow with your desired configurations.

```go
confStateList := &controller.Configuration{}
config.Configuration("controller_config.json", confStateList)
err := controller.BuildStates(confStateList)
if err != nil {
    logger.Fatal(err)
}
```

## Workflow States 🔄

### Defining States
Create a JSON file following the schema below:

```json
{
    "states_list": [
        {
            "name": "state1", // 'state1' is an example name.
            "available_states": [ // Array of next possible states.
                {
                    "name": "state2", // 'state2' is an example name.
                    "function": "HANDLER-NAME", // This value should be one of the options provided by the system.
                    "auto_run": true, // Boolean to indicate whether this state is automatically executed by the system or needs user intervention.
                    "priority": 0, // Value to sort next states; useful if multiple states can be set simultaneously.
                    "users": [ // Defines the Authority Condition: which types of users can transition to this state.
                        "User", // The system has only 5 hardcoded user types.
                        "Hunter",
                        "Admin",
                        "System",
                        "Vehicle"
                    ]
                },
                {
                    // ... Add another available state.
                }
            ]
        },
        // Add another status...
        {
            "name": "state2",
            "available_states": [
                {
                    // ...
                }
            ]
        }
    ]
}
```

### Status Flow Strategy
DWF is built to be dynamic, allowing you to modify the workflow without changing the codebase.

#### Conditions
- **Authority Condition**: User roles determine the accessible states.
- **Functional Condition**: Custom logic can further restrict state transitions.

## Handlers 🎛️
1. **voidHandler**: No conditions.
2. **batteryLowHandler**: Battery < 20%.
3. **after930PM**: After 9:30 PM local time.
4. **after48H**: No state change for 48 hours.

## User Roles 👥
*Note: User roles are hardcoded and require a code update to modify.*

1. **User**
2. **Hunter**
3. **Admin**
4. **Vehicle**
5. **System**

## Examples 📚
Refer to the `exvehicle` package for a detailed example involving a vehicle workflow.

### Vehicle Struct
```go
type Vehicle struct {
    ID string
}
```

### Public Methods
- `InitializeVehicle(id, state string, batteryPercentage int) (*Vehicle, error)`
- `SetBatteryPercentage(newBatteryPercentage int)`
- `State() string`
- `AvailableStates() (string, []string)`
- `AdminForceChangeState(nextState string, userType int) error`
- `ChangeState(nextState string, userType int) error`

## Contributions 🤝
Feel free to open issues or submit PRs for additional handlers or improvements.
