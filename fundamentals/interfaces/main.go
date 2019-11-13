package main

import "fmt"

type Persona struct {
}

type Authenticator interface {
	Authenticate(Persona) bool
}

type AWSAuthenticator struct {
}

func (a AWSAuthenticator) Authenticate(p Persona) bool {
	fmt.Println("To implement an interface in Go, we just need to implement all the methods in the interface. ")
	return true
}

func main() {
	var authenticator Authenticator
	authenticator = AWSAuthenticator{}

	authenticator.Authenticate(Persona{})
}
