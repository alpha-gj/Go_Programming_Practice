package main

import "fmt"
import "unicode/utf8"

func main() {

	ch := "生活大爆炸"
	en := "life"


	fmt.Printf("string is %s, data type is %T, and string len is %d by len(%s)\n", en, ch, len(en), en)
	fmt.Printf("string is %s, data type is %T, and string len is %d by len(%s)\n", ch, ch, len(ch), ch)
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("Using the len() to cal Unicode is not ok, beacasue one chinese word allocates 3 byte\n")
	fmt.Printf("it should use utf8.RuneCountInString\n")
	fmt.Printf("string is %s, data type is %T, and the string len is %d by utf8.RuneCountInString\n", ch, ch, utf8.RuneCountInString(ch))


	ch_v2 := "大爆炸"
	fmt.Printf("string is %s, data type is %T, and counting string len is %d by using utf8.RuneCountInString \n", ch_v2, ch_v2, utf8.RuneCountInString(ch_v2));
	


}
