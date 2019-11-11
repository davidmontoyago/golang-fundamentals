package main

import (
	"fmt"

	"github.com/davidmontoyago/golang-fundamentals/fundamentals/packages/mylibpackage"
	foo "github.com/davidmontoyago/golang-fundamentals/fundamentals/packages/mylibpackage"
)

func main() {
	fmt.Println("Every Go program is made up of packages.")
	fmt.Println("Programs start running in package main.")

	mylibpackage.SaySomething("By convention, the package name is the same as the last element of the import path.")
	foo.SaySomething("Package names can also be imported with a custom name.")

}
