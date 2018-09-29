package main

import "fmt"

func main() {
	thmem := "生活大爆炸 start"
	fmt.Println("Using the ascii")
	for i := 0; i < len(thmem); i++ {
		fmt.Printf("ascii: %c %d\n", thmem[i], thmem[i])
	}

	fmt.Println("")
	for index, s := range thmem {
		fmt.Printf("Unicode: %c %d\n", s, s)
		fmt.Println(index)
	}

	for index, s := range thmem {
		if index == 0 {
			fmt.Printf("The index of %d is %c\n", index, s)
		}
		if index == 3 {
			fmt.Printf("The index of %d is %c\n", index, s)
		}
		if index == 6 {
			fmt.Printf("The index of %d is %c\n", index, s)
		}
		if index == 9 {
			fmt.Printf("The index of %d is %c\n", index, s)
		}
		if index == 12 {
			fmt.Printf("The index of %d is %c\n", index, s)
		}
	}
}
