package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky123456789", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}
