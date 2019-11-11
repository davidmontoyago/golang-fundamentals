package main

import (
	"fmt"
	"runtime"
)

func main() {
	// If
	if status := true; status {
		fmt.Println("The if statement can start with a short statement to execute before the condition.")
	}

	// For
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d. Go has only one looping construct, the for loop.\n", i)
	}

	sum := 1
	for sum <= 3 {
		fmt.Printf("%d. For is Go's \"while\"\n", sum)
		sum += 1
	}

	// Switch
	fmt.Println("A switch statement is a shorter way to write a sequence of if - else statements.")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on OS X.")
	case "linux":
		fmt.Println("Running on Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("Running on %s.\n", os)
	}

	defer deferredFunction()
}

func deferredFunction() {
	fmt.Println("A defer statement defers the execution of a function until the surrounding function returns.")
}
