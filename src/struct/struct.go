package main

import "fmt"

type Wireless_setting struct {                      
	Enable          string  `json:"enable"`
	Mode            string  `json:"mode"`
	Essid           string  `json:"essid"`
	Chpatterns      string  `json:"chpatterns"`
	Channel         int     `json:"channel"`
	Auth            string  `json:"auth"`
	Encryption      string  `json:"encryption"`
	Passphrase      string  `json:"passphrase"`
}

func main() {
	fmt.Println("vim-go")
}
