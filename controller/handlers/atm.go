package handlers

// IIsValidCardHandler :
type IIsValidCardHandler interface {
	IsValidCardHandler() bool
}

// IsValidCardHandler :
func IsValidCardHandler(obj interface{}) bool {
	v, ok := obj.(IIsValidCardHandler)
	if !ok {
		return false
	}
	return v.IsValidCardHandler()
}

// INotValidHandler :
type INotValidHandler interface {
	NotValidHandler() bool
}

// NotValidHandler :
func NotValidHandler(obj interface{}) bool {
	v, ok := obj.(INotValidHandler)
	if !ok {
		return false
	}
	return v.NotValidHandler()
}

// IValidCardHandler :
type IValidCardHandler interface {
	ValidCardHandler() bool
}

// ValidCardHandler :
func ValidCardHandler(obj interface{}) bool {
	v, ok := obj.(IValidCardHandler)
	if !ok {
		return false
	}
	return v.ValidCardHandler()
}

// IPrintNotValidHandler :
type IPrintNotValidHandler interface {
	PrintNotValidHandler() bool
}

// PrintNotValidHandler :
func PrintNotValidHandler(obj interface{}) bool {
	v, ok := obj.(IPrintNotValidHandler)
	if !ok {
		return false
	}
	return v.PrintNotValidHandler()
}

// IShowPinScreenHandler :
type IShowPinScreenHandler interface {
	ShowPinScreenHandler() bool
}

// ShowPinScreenHandler :
func ShowPinScreenHandler(obj interface{}) bool {
	v, ok := obj.(IShowPinScreenHandler)
	if !ok {
		return false
	}
	return v.ShowPinScreenHandler()
}

// IValidatePinHandler :
type IValidatePinHandler interface {
	ValidatePinHandler() bool
}

// ValidatePinHandler :
func ValidatePinHandler(obj interface{}) bool {
	v, ok := obj.(IValidatePinHandler)
	if !ok {
		return false
	}
	return v.ValidatePinHandler()
}

// IValidPinHandler :
type IValidPinHandler interface {
	ValidPinHandler() bool
}

// ValidPinHandler :
func ValidPinHandler(obj interface{}) bool {
	v, ok := obj.(IValidPinHandler)
	if !ok {
		return false
	}
	return v.ValidPinHandler()
}

// IInValidPinHandler :
type IInValidPinHandler interface {
	InValidPinHandler() bool
}

// InValidPinHandler :
func InValidPinHandler(obj interface{}) bool {
	v, ok := obj.(IInValidPinHandler)
	if !ok {
		return false
	}
	return v.InValidPinHandler()
}

// IChooseActionHandler :
type IChooseActionHandler interface {
	ChooseActionHandler() bool
}

// ChooseActionHandler :
func ChooseActionHandler(obj interface{}) bool {
	v, ok := obj.(IChooseActionHandler)
	if !ok {
		return false
	}
	return v.ChooseActionHandler()
}

// IChooseMoneyHandler :
type IChooseMoneyHandler interface {
	ChooseMoneyHandler() bool
}

// ChooseMoneyHandler :
func ChooseMoneyHandler(obj interface{}) bool {
	v, ok := obj.(IChooseMoneyHandler)
	if !ok {
		return false
	}
	return v.ChooseMoneyHandler()
}

// IShowAccountBalanceHandler :
type IShowAccountBalanceHandler interface {
	ShowAccountBalanceHandler() bool
}

// ShowAccountBalanceHandler :
func ShowAccountBalanceHandler(obj interface{}) bool {
	v, ok := obj.(IShowAccountBalanceHandler)
	if !ok {
		return false
	}
	return v.ShowAccountBalanceHandler()
}

// IShowMoneyScreenHandler :
type IShowMoneyScreenHandler interface {
	ShowMoneyScreenHandler() bool
}

// ShowMoneyScreenHandler :
func ShowMoneyScreenHandler(obj interface{}) bool {
	v, ok := obj.(IShowMoneyScreenHandler)
	if !ok {
		return false
	}
	return v.ShowMoneyScreenHandler()
}

// ISufficientFundHandler :
type ISufficientFundHandler interface {
	SufficientFundHandler() bool
}

// SufficientFundHandler :
func SufficientFundHandler(obj interface{}) bool {
	v, ok := obj.(ISufficientFundHandler)
	if !ok {
		return false
	}
	return v.SufficientFundHandler()
}

// IRelaseMoneyHandler :
type IRelaseMoneyHandler interface {
	RelaseMoneyHandler() bool
}

// RelaseMoneyHandler :
func RelaseMoneyHandler(obj interface{}) bool {
	v, ok := obj.(IRelaseMoneyHandler)
	if !ok {
		return false
	}
	return v.RelaseMoneyHandler()
}

// IEndHandler :
type IEndHandler interface {
	EndHandler() bool
}

// EndHandler :
func EndHandler(obj interface{}) bool {
	v, ok := obj.(IEndHandler)
	if !ok {
		return false
	}
	return v.EndHandler()
}
