package main

import "fmt"

func main() {
	const maxProfileId int = 10

	sum := 0
	for index := 0; index < maxProfileId; index++ {
		sum += index
	}

	fmt.Println("sum is equal to ", sum)
}
