// go comes with multiple type 
// bool => True or False
// Numberic Types int8 , int16 , int32 , int64 , int  
        int8 => represent 8 bit -128 to 127 
        int16 => represent 16 bit -32768 to 32767
        int32 => represent 32 bit -2147483648 to 2147483647
        int64 => represent 64 bit -9223372036854775808 to 9223372036854775807
        int => can use 32 or 64 bit systems  
        You should generally be using int to represent integers unless there is a need to use a specific sized integer.
        
// String 
// float32 , float64
// byte 
// complex 

// un comment code between /* *\ to run it 
/*
func main() {

	a, b := 145.8, 543.9
	fmt.Println("value of a =>", a, "and value of b =>", b)
}

*/

/*
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a, b := 145.8, 543.9
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b
}
*/

/*

func main() {

	a, b := 5.67, 8.97
	fmt.Printf("type of a %T and Type of b %T", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("Sum is ", sum, "Diff is ", diff)
	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

}

*/

/*

func main() {
	first := "Roger"
	lastname := "Steve"
	fmt.Println(first, lastname)
}

*/
