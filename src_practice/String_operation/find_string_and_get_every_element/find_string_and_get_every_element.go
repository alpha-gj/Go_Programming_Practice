package main

import "fmt"
import "unicode/utf8"

func main() {

	thmem := "生活大爆炸 start"
	fmt.Printf("Printf the theme is %s, and printf every char for it\n", thmem);

	// Using
	fmt.Printf("# Using index\n");
	for i := 0; i < len(thmem); i++ {
		fmt.Printf("ascii: %c %d\n", thmem[i], thmem[i]);
	}
	fmt.Printf("len of thmem is %d (it is wrong len), by using len(thmem)\n", len(thmem));

	fmt.Printf("\n\n");

	fmt.Printf("# Using range\n");
	for i, s := range thmem {
		fmt.Printf("i is %d, s is %c %d\n", i, s, s);
	}
	fmt.Printf("len of thmem is %d (it is correct), by using utf8.RuneCountInString(thmem)\n", utf8.RuneCountInString(thmem));

	for i, c := range "go" {
        fmt.Println(i, c)
		fmt.Printf("c is %c\n", c);
    }
}
