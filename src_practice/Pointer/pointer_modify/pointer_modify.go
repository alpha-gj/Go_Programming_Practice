package main

import "fmt"


func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
func main() {

	x, y := 1, 2
	fmt.Printf("x is %d, y is %d\n", x ,y)


	fmt.Printf("swap!\n")
	swap(&x, &y)

	fmt.Printf("x is %d, y is %d\n", x ,y)
}
