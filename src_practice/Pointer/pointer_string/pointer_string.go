package main

import "fmt"

func main() {

	var cat int = 1
	fmt.Printf("address of cat is %p\n", &cat)

	var str string = "banana"
	fmt.Printf("address of str is %p\n", &str)
}
