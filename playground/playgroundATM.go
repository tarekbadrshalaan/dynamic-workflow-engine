package playground

import (
	"context"
	"dwf/controller"
	"dwf/exatm"
	"fmt"
	"time"
)

// PlayFullATMStateCycle : ..
func PlayFullATMStateCycle() (*exatm.Card, error) {
	fmt.Println("========================= PlayFullVehicleStateCycle ===============================")
	myAwesomeCard, err := exatm.InitializeCard(context.Background(), "num1", "insertCard", "123", 100)
	if err != nil {
		return nil, fmt.Errorf("Error while initialize new Cards : %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	err = myAwesomeCard.ChangeState("validatePin", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
		return myAwesomeCard, err
	}
	time.Sleep(100 * time.Millisecond)

	err = myAwesomeCard.ChangeState("chooseMoney", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
		return myAwesomeCard, err
	}
	time.Sleep(100 * time.Millisecond)

	err = myAwesomeCard.ChangeState("sufficientFund", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
		return myAwesomeCard, err
	}
	time.Sleep(100 * time.Millisecond)

	err = myAwesomeCard.ChangeState("end", controller.ADMIN)
	if err != nil {
		fmt.Println(err)
		return myAwesomeCard, err
	}
	time.Sleep(100 * time.Millisecond)

	return myAwesomeCard, nil
}
