package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileHandle, _ := os.Open("test")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
}
