package handlers

// VoidHandler : handler accept any chanage
func VoidHandler(obj interface{}) bool {
	return true
}

// IPrint :
type IPrint interface {
	Print() bool
}

// PrintHandler :
func PrintHandler(obj interface{}) bool {
	p, ok := obj.(IPrint)
	if !ok {
		return false
	}
	return p.Print()
}
