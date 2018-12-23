package main

import "fmt"
import "strings"
import "unicode/utf8"

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

	//sol. 
	//死    神    來    了      ,        死       神       b  y  e  b  y  e
	//0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26
	fmt.Printf("len of tracer is %d\n", utf8.RuneCountInString(tracer))
}
