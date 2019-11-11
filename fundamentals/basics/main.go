package main

import (
	"fmt"

	"github.com/davidmontoyago/golang-fundamentals/fundamentals/basics/mylibpackage"
)

var (
	packageScopedVariable = "A var statement can be at package or function level."
)

const (
	CONSTANT = "Constants are declared like variables, but with the const keyword."
)

func main() {
	var declared string
	declared = "A variable can declared and initialized separately (Long form)."
	fmt.Println(declared)

	implicitlyDeclared := "A variable can be declared and initialized in a single statement (Short form, implicit type)."
	fmt.Println(implicitlyDeclared)

	fmt.Println(packageScopedVariable)

	mylibpackage.ThisHasPublicScope()
	// mylibpackage.thisHasPrivateScope()

	msg, status := mylibpackage.FunctionsCanReturnMultipleValues()
	if status {
		fmt.Println(msg)
	}

	msg, status = mylibpackage.FunctionsCanHaveNamedReturnValues()
	if status {
		fmt.Println(msg)
	}

	fmt.Println(CONSTANT)

	var i int
	var boolz bool
	var unsetMsg string
	fmt.Println("Variables declared without an explicit initial value are given their \"zero value\".")
	fmt.Printf("i=%d, boolz=%v, unsetMsg=%s", i, boolz, unsetMsg)
}
