package main

import "fmt"

func main() {

	var a int = 100
	var b int = 200

	a, b = b, a	// No need the 'temp' variable to keep the temp data

	fmt.Println(a, b)
}
