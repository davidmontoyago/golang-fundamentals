package main

import (
	"errors"
	"fmt"
	"time"
)

type ValidationError struct {
	When time.Time
	What string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func main() {
	fmt.Println("In Go, errors are a type and are not throwable.")
	_, err := ValidateOrFail("")
	if err != nil {
		fmt.Printf("Error values must be checked: %v\n", err)
	}

	fmt.Println("Custom error types can be created.")
	var verr *ValidationError
	_, verr = ValidateOrFailWithCustomError("")
	if verr != nil {
		fmt.Printf("Use a custom types to abstract away domain errors or when more details are required: %+v\n", verr)
	}
}

func ValidateOrFail(payload string) (bool, error) {
	if payload == "" {
		return false, errors.New("A payload is required!")
	}
	return false, nil
}

func ValidateOrFailWithCustomError(payload string) (bool, *ValidationError) {
	if payload == "" {
		return false, &ValidationError{
			What: "A payload is required!",
			When: time.Now(),
		}
	}
	return true, nil
}
