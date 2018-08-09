package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type Sitesurvery_info struct {
    Ssid       string	`json:"ssid"`
    Signal     int      `json:"signal"`
    Mode       string   `json:"mode"`
    Channel    int      `json:"channel"`   
    Auth       string   `json:"auth"`   
    Encryption string   `json:"encryption,omitempty"` // if it empty, it will not show it
}   

type output struct {
	Sites []Sitesurvery_info `json:"sites"`
}

func main() {
	
// How to loop it?
	var res output
	res.Sites = append(res.Sites, Sitesurvery_info {
		Ssid 		: "gj_test1",
		Signal 		: 123,
		Mode		: "infrastructure",
		Channel 	: 1,
		Auth		: "WPA2-PSK",
		Encryption 	: "",
	})

	res.Sites = append(res.Sites, Sitesurvery_info {
		Ssid 		: "gj_test2",
		Signal 		: 123,
		Mode		: "infrastructure",
		Channel 	: 1,
		Auth		: "WPA2-PSK",
		Encryption 	: "AES",
	})
//	for i := 0; i < 10; i++ {
//		res.Sites[].Ssid = "hi"
//		res.Sites[].Signal = i
//		res.Sites[].Mode = "hi~"
//		res.Sites[].Channel = i
//		res.Sites[].Auth = "hi"
//		res.Sites[].Encryption = "hi"
//	}
//	fmt.Printf("%s %d %s %d %s %s\n", 
//	res.Sites[0].Ssid, 
//	res.Sites[0].Signal,
//	res.Sites[0].Mode,
//	res.Sites[0].Channel,
//	res.Sites[0].Auth,
//	res.Sites[0].Encryption,
//)
// Way2
//	res := output { 
//		[]Sitesurvery_info{
//			Sitesurvery_info {
//				Ssid 		: "gj_test1",
//				Signal 		: 123,
//				Mode		: "infrastructure",
//				Channel 	: 1,
//				Auth		: "WPA2-PSK",
//				Encryption 	: "",
//			},
//			Sitesurvery_info {
//				Ssid 		: "gj_test2",
//				Signal 		: 123,
//				Mode		: "infrastructure",
//				Channel 	: 1,
//				Auth		: "WPA2-PSK",
//				Encryption 	: "AES",
//			},
//		},
//	}

    pagesJson, err := json.Marshal(res)
    if err != nil {
        log.Fatal("Cannot encode to JSON ", err)
    }
    fmt.Fprintf(os.Stdout, "%s", pagesJson)
}
