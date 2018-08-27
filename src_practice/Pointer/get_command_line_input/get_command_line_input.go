package main

import "fmt"
import "flag"

//paramter, default value, help context 
var cmd = flag.String("mode", "hi default", "process mode")

func main() {

	flag.Parse()
	fmt.Println(*cmd)
	fmt.Printf("data type of cmd is %T\n", cmd)
}
