// Variable is the name given to a memory location to store a value of a specific type. There are various syntaxes to declare variables in Go. Let's look at them one by one.
// How to declare variable ?
// var name type as example => var age int
// if int is not assigned with value by default it will be zero

package main

import "fmt"

func main() {
	var age int // variable declaration
	fmt.Println("My age is", age)
}
