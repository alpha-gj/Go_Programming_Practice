package main

import "fmt"
import "unicode/utf8"

func main() {

	ch := "狙擊 start"

	fmt.Printf("string is %s, data type is %T, and string len is %d by len(%s)\n", ch, ch, len(ch), ch)
	for i := 0; i < len(ch); i++ {
		fmt.Printf("ascii: %c, %d\n", ch[i], ch[i])
	}

	fmt.Printf("it should use utf8.RuneCountInString\n")
	fmt.Printf("string is %s, data type is %T, and the string len is %d by utf8.RuneCountInString\n", ch, ch, utf8.RuneCountInString(ch))

	for  _, s := range ch {
		fmt.Printf("Unicode: %c %d\n", s, s)
	}
}
