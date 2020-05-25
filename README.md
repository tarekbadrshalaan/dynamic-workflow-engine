# dynamic-workflow-engine (DWE)
dynamic workflow engine

```
DWE is a POC of creating dynamic workflow engine in golang,
basically DWE could implement any workflow see (exatm, exvehicle  directories).
```

## Components 

- **controller** tha main package to handle workflow and test package.
- **exatm** example using atm system.
- **exvehicle** example using vehicle system.
- **api** API interface to use the engine.
- **config** simple package to handle read configuration from json file.
- **logger** logging service include 3 type of logger implementation.
- **playground** external package to use package controller, and main package use *playground* to run.


## How to use **DWE**

DWE provide you with nice *controller* package to control your workflow.

so basically you need to extract *controller* package and use it in your soft real­time API solution.


### **Status flow** strategy
```
dwf status flow built to be dynamic and configurable, that's mean your work flow can be increase, decrease ,change status position without touch the code.
```
- `Status` is the base unit to the flow system.
- Each `status` contains it's name and next available status.
- To move from `status` to next available status, you should go throw two conditions 
    - **Authority Condition** : each user in the system has some status he can go throw them, and this Authorization roles is configurable throw status configuration.
    - **Function Condition** : to go from state to available state the authorized user go throw condition function which simply allow or disallow change state to happen.

- to set `status` flow, you have to path
    - build your own flow in a json file and pass the file path to `BuildStates` e.g.
        ```
        	confstateList := &controller.Configuration{}
            config.Configuration("controller_config.json", confstateList)
            err := controller.BuildStates(confstateList)
            if err != nil {
                logger.Fatal(err)
            }
        ```


### build your own **status flow**

- create json file with this schema
```
{
    "states_list": [
        {
            "name": "state1", // 'state1' is an example name 
            "available_states": [ // array of next states
                {
                    "name": "state2", // 'state2' is an example name 
                    "function": "HANDLER-NAME", // this value should be one of the choices provided by the system.
                    "auto_run": true, // boolen to set if this state running by the system or should be executed by the user. 
                    "priority": 0, // value to sort next states, if there is many available states can be setted in the same time
                    "users": [ // set the Authority condition: which type of user can change to this state.
                        "User",// there is only 5 hard coded type in the system.
                        "Hunter",
                        "Admin",
                        "System",
                        "Vehicle"

                    ]
                },
                {
                    //... add another available state.
                }
            ]
        },
        // add another status ... 
        {

            "name": "state2",
            "available_states": [
                {
                    // ...
                }
            ]
        }, 
    ]
}
```
### Available **Handlers** in the system 
1. `voidHandler` : basic handler accept to go to the next state without conditions

2. `batteryLowHandler` : accept to go to the next state if the battery is less then 20%

3. `after930PM` : accept to go to the next state if the local time after 9:30 pm

4. `after48H` : accept to go to the next state if there is no change in vehicle state 48 hour ago 

#### Note: feel free to ask for other handler/logic the dynamic architecture allow us to change in a very simple way  

### Available **Users** in the system 
```
NOTE: the currant users system is hard coded so it need developer/code updated to add or remove user type
```
1. `User`: basic user, and has no exception features

2. `Hunter`: basic user, and has no exception features

3. `Admin`: has the basic features, plus able to use `AdminForceChangeState` method

4. `Vehicle`: basic user, and has no exception features, we created it to mention requests from vehicle itself 

5. `System`: basic user, we created it to mention requests from automation system.

## Examples : 

### **exvehicle** package

*exvehicle* package provide you with **Vehicle** struct which is our base unit to build the workflow

- Vehicle struct contain only one public property which is `ID`
    ```
    type Vehicle struct {
    	ID                   string
    }
    ```
- to Initialize new Vehicle use `InitializeVehicle` function with input about currant `state` of the vehicle and `battery Percentage` in return you will get new `Vehicle` instances or `error`
    ```
    InitializeVehicle(id, state string, batteryPercentage int) (*Vehicle, error)
    ```
- `Vehicle` instances provide you 5 public methods to control your instance
    
    - `SetBatteryPercentage` to reset battery percentage of your vehicle with no return e.g.
    ```
    yourvehicle.SetBatteryPercentage(newBatteryPercentage int)
    ```
    
    - `State` to get the currant state of your vehicle e.g.
    ```
    yourvehicle.State() currant_state string
    ```
    
    - `AvailableStates` to get the currant state and available next state of your vehicle e.g.
    ```
    yourvehicle.AvailableStates()(currant_state string, AvailableStates []string)
    ```
    
    - `AdminForceChangeState` to be able to force state change from admin with error in return e.g.
    ```
    // WARNING: this method will not check the user (authorization/session)
    // the permission check should be done before call this method

    yourvehicle.AdminForceChangeState(nextState string, usertype int) error
    ```
    
    - `ChangeState` to change your vehicle state, but you have to follow th role that has been set in the workflow core **StateList** e.g.
    ```
    yourvehicle.ChangeState(nextState string, usertype int) error
    ```

each vehicle contain **state** and this state controlled by workflow has been set before in the system
