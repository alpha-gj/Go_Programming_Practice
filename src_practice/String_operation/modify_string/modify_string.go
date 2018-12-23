package main

import "fmt"
//import "unicode/utf8"

func main() {
	angel := "Heros never die"

	angleBytes := []byte(angel)

	for i := 6; i <= 10; i++ {
		angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes))
	fmt.Println(angleBytes)

	result_string := string(angleBytes)
	//result_string[0] = 'h'			/* ./modify_string.go:18: cannot assign to result_string[0] */
	fmt.Println(result_string)
}
