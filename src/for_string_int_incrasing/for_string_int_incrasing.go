package main

import "fmt"
import "strconv"

func main() {

	var id_index int = 0

	for {

		fmt.Printf("Bitrate=%s\n", strconv.Itoa(id_index)) //OK
		fmt.Printf("Bitrate="+strconv.Itoa(id_index)+"\n") //OK

		id_index++

		if id_index > 10 {
			break;
		}

	}
}

