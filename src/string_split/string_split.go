package main

import "fmt"
import "strings"


func main() {

	var resolution = "1920x1080"
	var result = strings.Split(resolution, "x")

	var width = result[0]
	var height = result[1]

	fmt.Printf("%s and %s", width, height)
}
