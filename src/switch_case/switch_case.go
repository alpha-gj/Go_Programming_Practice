package main

import "fmt"
import "strconv"

func GetChannelByFreq (freq string) int {

	switch freq{
		case "2412":
			return 1
		case "2417":
			return 2
		case "2422":
			return 3
		case "2427":
			return 4
		case "2432":
			return 5
		case "2437":
			return 6
		case "2442":
			return 7
		case "2447":
			return 8
		case "2452":
			return 9
		case "2457":
			return 10
		case "2462":
			return 11
		case "2467":
			return 12
		case "2472":
			return 13
		default:
			return 6
	}

}

func main() {
	fmt.Println("vim-go")
	for i := 2412; i <= 2472; i+=5 {
		fmt.Printf("The channel is %d\n", GetChannelByFreq(""+strconv.Itoa(i)+""))
	}
}
