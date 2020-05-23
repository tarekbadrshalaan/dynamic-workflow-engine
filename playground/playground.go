package playground

import (
	"context"
	"dwf/controller"
	"fmt"
	"time"
)

// PlayFullVehicleStateCycle : ..
func PlayFullVehicleStateCycle() (*controller.Vehicle, error) {
	fmt.Println("========================= PlayFullVehicleStateCycle ===============================")
	myAwesomeVehicle, err := controller.InitializeVehicle(context.Background(), "num1", "Ready", 100)
	if err != nil {
		return nil, fmt.Errorf("Error while initialize new vehicles : %v", err)
	}

	err = myAwesomeVehicle.ChangeState("Riding", controller.USER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}
	time.Sleep(500 * time.Millisecond)
	myAwesomeVehicle.Terminate()
	time.Sleep(500 * time.Millisecond)

	err = myAwesomeVehicle.ChangeState("Ready", controller.USER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Riding", controller.USER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	myAwesomeVehicle.SetBatteryPercentage(10)
	time.Sleep(1 * time.Second)
	err = myAwesomeVehicle.ChangeState("Collected", controller.HUNTER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Dropped", controller.HUNTER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Ready", controller.HUNTER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	return myAwesomeVehicle, nil
}

// PlayReadyBounty : ..
func PlayReadyBounty() (*controller.Vehicle, error) {
	fmt.Println("========================= PlayReadyBounty ===============================")
	myAwesomeVehicle, err := controller.InitializeVehicle(context.Background(), "num1", "Ready", 50)
	if err != nil {
		return nil, fmt.Errorf("Error while initialize new vehicles : %v", err)
	}

	err = myAwesomeVehicle.AdminForceChangeState("Bounty", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Collected", controller.HUNTER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Dropped", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Ready", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	return myAwesomeVehicle, err
}

// PlayReadyUnknown : ..
func PlayReadyUnknown() (*controller.Vehicle, error) {
	fmt.Println("========================= PlayReadyUnknown ===============================")

	myAwesomeVehicle, err := controller.InitializeVehicle(context.Background(), "num1", "Ready", 21)
	if err != nil {
		return nil, fmt.Errorf("Error while initialize new vehicles : %v", err)
	}

	err = myAwesomeVehicle.ChangeState("Riding", controller.USER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Ready", controller.USER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Riding", controller.USER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Ready", controller.USER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Unknown", controller.SYSTEM)
	if err != nil {
		fmt.Println(err)
	}
	err = myAwesomeVehicle.AdminForceChangeState("Unknown", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	return myAwesomeVehicle, err
}

// PlayInvalidUser : ..
func PlayInvalidUser() (*controller.Vehicle, error) {
	fmt.Println("========================= PlayInvalidUser ===============================")
	myAwesomeVehicle, err := controller.InitializeVehicle(context.Background(), "num1", "Riding", 40)
	if err != nil {
		return nil, fmt.Errorf("Error while initialize new vehicles : %v", err)
	}

	err = myAwesomeVehicle.ChangeState("Battery-Low", controller.USER)
	if err != nil {
		fmt.Println(err)
	}

	err = myAwesomeVehicle.ChangeState("Ready", controller.HUNTER)
	if err != nil {
		fmt.Println(err)
		return myAwesomeVehicle, err
	}

	err = myAwesomeVehicle.ChangeState("Unknown", controller.HUNTER)
	if err != nil {
		fmt.Println(err)
	}

	return myAwesomeVehicle, err
}

// Dynamic : ..
func Dynamic() (*controller.Vehicle, error) {
	fmt.Println("========================= PlayInvalidUser ===============================")
	myAwesomeVehicle, err := controller.InitializeVehicle(context.Background(), "num1", "Ready", 40)
	if err != nil {
		return nil, fmt.Errorf("Error while initialize new vehicles : %v", err)
	}

	err = myAwesomeVehicle.ChangeState("print", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
	}

	return myAwesomeVehicle, err
}
