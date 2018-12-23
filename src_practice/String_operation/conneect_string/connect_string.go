package main

import "fmt"
import "bytes"

func main() {

	hammer := "吃我一錘"
	sickle := "死吧"

	// Using bytes.Buffer
	var stringBuilder bytes.Buffer

	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)

	fmt.Println(stringBuilder.String())


	// Using "+"
	test_result := hammer + sickle
	fmt.Println(test_result)

}
