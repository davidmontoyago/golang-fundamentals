package mylibpackage

import "fmt"

func ThisHasPublicScope() {
	fmt.Println("ThisHasPublicScope() can be called from anywhere.")
	thisHasPrivateScope()
}

func thisHasPrivateScope() {
	fmt.Println("thisHasPrivateScope() can only be called from within mylibpackage.")
}

func FunctionsCanReturnMultipleValues() (string, bool) {
	return "A function can return any number of results.", true
}

func FunctionsCanHaveNamedReturnValues() (msg string, status bool) {
	msg = "The return values of a function can be named allowing you to have a \"naked\" return."
	status = true
	return
}
