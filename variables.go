// Variable is the name given to a memory location to store a value of a specific type. There are various syntaxes to declare variables in Go. Let's look at them one by one.
// How to declare variable ?
// var name type as example => var age int
// if int is not assigned with value by default it will be zero to test it uncomment /*   *\

/* package main

import "fmt"

func main() {
	var age int // variable declaration
	fmt.Println("My age is", age)
}

*/

// A variable can be assigned to any value of its type. In the above program, age can be assigned any integer value.

/*

func main() {

	var age int // variable declaration
	fmt.Println("My age is", age)
	age = 29
	fmt.Println("My age is", age)
	age = 40
	fmt.Println("My new age is", age)
}

*/

// Multiple variable declaration
/*

func main() {
	var width, height = 100, 50
	fmt.Println("My Width", width, "My Height", height)
}

*/ 
// Multiple variable 

/*
func main() {  
    var (
        name   = "naveen"
        age    = 29
        height int
    )
    fmt.Println("my name is", name, ", age is", age, "and height is", height)
}
*/


// Short hand declaration you can use this for first time on variable declaration after that it will cause errors 
/* 
func main() {
	count := 10
	fmt.Println(count)
}
*/ 

// multiple variables in a single line using short hand syntax.

/*
func main() {

	name, age := "Roger", 27
	fmt.Println("My name is ", name, "And my age is ", age)
}

*/

/*
package main

import (
	"fmt"
	"math"
)

func main() {

	a, b := 145.8, 543.9
	c := math.Min(a, b)
	fmt.Println(c)
}

*/
