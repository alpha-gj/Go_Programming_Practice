package main

import "fmt"
import "strings"

//using strings.Index to find substring

func main() {
	//NOTE there are two blanks after comma(,)
	tracer := "死神來了,  死神bye bye"

	comma := strings.Index(tracer, ",")
	fmt.Printf("comma is %d\n", comma)

	pos := strings.Index(tracer[comma:], "死神")
	fmt.Printf("pos is %d\n", pos)

	fmt.Println(comma, pos, tracer[comma+pos:])
	fmt.Println(comma, pos, tracer[comma:])
}
