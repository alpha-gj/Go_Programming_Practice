package main

import "fmt"
import "os/exec"
import "strconv"

/* DB get function */
func getDbData(data string) string {

		command := exec.Command("echo", "-n", data)

		buf, err := command.Output()
		if err != nil {
			fmt.Println(err)
		}

		return string(buf)
}

func main() {

//		enable := []string {"off", "on"}
		
//		index := getDbData("0")
//		value, err := strconv.Atoi(string(index))
		
		value, err := strconv.Atoi(getDbData("0"))
		if err != nil {
			fmt.Println(err)
		}
        fmt.Printf("enable=%s, %d\n", "0", value)
}
