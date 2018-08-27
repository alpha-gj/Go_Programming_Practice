package main

import "fmt"

func main() {

	// Allocate the mem which is data type of string for str
	// var str *string = new(string)  // use it is ok
	str := new(string)


	// Let dereference of str and assign the string to it
	*str = "ninja"

	fmt.Printf("dereference of str is %s\n", *str)
	fmt.Printf("data type of str is %T\n", str)

}
