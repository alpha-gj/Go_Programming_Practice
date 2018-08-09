package main

import "fmt"

func main() {
	var flag = false

	if !flag {
		fmt.Println("false\n")
	}

	flag = true

	
	if flag {
		fmt.Println("true\n")
	}

	fmt.Println("vim-go")
}
