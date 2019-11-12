package main

import "fmt"

type AdmissionReview struct {
	Kind    string
	Spec    string
	Version string
}

func (r *AdmissionReview) DoSomething() {
	fmt.Println("just did something.")
}

func main() {
	review := AdmissionReview{
		Kind:    "Deployment",
		Version: "betav1",
	}
	fmt.Printf("A struct is a collection of fields like %+v\n", review)

	fmt.Println("A struct type can have behavior attached like AdmissionReview.DoSomething()")
	review.DoSomething()

	fmt.Println("In Go, pointers are a type.")
	thisIsAPointer := &AdmissionReview{
		Kind:    "Job",
		Version: "alphav1",
	}
	thisIsAValue := AdmissionReview{
		Kind:    "Job",
		Version: "alphav1",
	}
	fmt.Printf("this is a pointer: %p\n", thisIsAPointer)
	fmt.Printf("this is a value %v\n", thisIsAValue)
	fmt.Printf("with * you can get the value of the pointer (dereference): %v\n", *thisIsAPointer)
	fmt.Printf("with & you can get the pointer of the value: %p\n", &thisIsAPointer)

	callback := func() {
		fmt.Println("Functions are a type too!")
	}
	callback()
}
