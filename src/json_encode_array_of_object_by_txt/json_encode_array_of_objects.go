package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os/exec"
    "os"
	"strings"
	"strconv"
	"bufio"
//	"io"
//	"io/ioutil"
)

func execSystem(cmd string) string {

		command := exec.Command("sh", "-c", cmd)

		buf, err := command.Output()
		if err != nil {
			fmt.Println(err)
		}

//		result := strings.Replace(string(buf), "\n", "", -1)
//		return result
		return string(buf)
}

func MustAtoi(s string) int {
    result_int, _ := strconv.Atoi(s)
    return result_int               
}

func GetChannelByFreq (freq string) int {

    switch freq {
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
            return -1
    }
}

func GetSignalBydbm (signal_dbm string) (Rssi_Quality int) {

	Rssi := MustAtoi(signal_dbm)

	if Rssi >= -50 {
		Rssi_Quality = 100
	} else if Rssi >= -80    {
		Rssi_Quality = (24 + ((Rssi + 80) * 26) / 10)	 // between -50 ~ -80dbm
	} else if Rssi >= -90 {
		Rssi_Quality = (((Rssi + 90) * 26) / 10) 		// between -80 ~ -90dbm 
	} else { 
		Rssi_Quality = 0 								// < -84 dbm
	}

    return Rssi_Quality
}

func GetAuthEncByFlags (Flags string) (Auth string, Enc string) {

	// Auth
	if strings.Contains(Flags, "WPA2") {
		Auth = "WPA2-PSK"
	} else if strings.Contains(Flags, "WPA")  {
		Auth = "WPA-PSK"
	} else {
		Auth = "open"
	}

	// Enc
	if strings.Contains(Flags, "CCMP") {
		Enc = "AES"
	} else if strings.Contains(Flags, "TKIP")  {
		Enc = "TKIP"
	} else {
		Enc = "none"
	}

	return Auth, Enc
}


type Sitesurvery_info struct {
    Ssid       string	`json:"ssid"`
    Signal     int      `json:"signal"`
    Mode       string   `json:"mode"`
    Channel    int      `json:"channel"`   
    Auth       string   `json:"auth"`   
    Encryption string   `json:"encryption"` 
//    Encryption string   `json:"encryption,omitempty"` // if it empty, it will not show it
}   

type Site_list struct {
	Sites []Sitesurvery_info `json:"sites"`
}


func main() {
	
//	sitesurvey_result := execSystem("cat ./test | tail -n +3 | head -n 1")
//	fmt.Println(sitesurvey_result)

	fileHandle, _ := os.Open("./test")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var count int = 0 
	const start_line = 2
	var res Site_list
	for fileScanner.Scan() {
		if count >= start_line {
			// fmt.Println(fileScanner.Text())
			s := strings.Fields(fileScanner.Text())

			ssid 				:= s[4]
			signal 				:= GetSignalBydbm(s[2])
			mode				:= "infrastructure"
			channel				:= GetChannelByFreq(s[1])
			auth, encryption	:= GetAuthEncByFlags(s[3])

			res.Sites = append(res.Sites, Sitesurvery_info {
				Ssid 		: ssid,
				Signal 		: signal,
				Mode		: mode,
				Channel 	: channel,
				Auth		: auth,
				Encryption 	: encryption,
			})

		} else {
			count++
		}
	}

    pagesJson, err := json.Marshal(res)
    if err != nil {
        log.Fatal("Cannot encode to JSON ", err)
    }
    fmt.Fprintf(os.Stdout, "%s", pagesJson)
}

//	s := strings.Fields(sitesurvey_result)
//
//	var res output
//	res.Sites = append(res.Sites, Sitesurvery_info {
//		Ssid 		: s[4],
//		Signal 		: MustAtoi(s[2]),
//		Mode		: "infrastructure",
//		Channel 	: MustAtoi(s[1]),
//		Auth		: s[3],
//		Encryption 	: "",
//	})
//
	
// How to loop it?
//	var res output
//	res.Sites = append(res.Sites, Sitesurvery_info {
//		Ssid 		: "gj_test1",
//		Signal 		: 123,
//		Mode		: "infrastructure",
//		Channel 	: 1,
//		Auth		: "WPA2-PSK",
//		Encryption 	: "",
//	})
//
//	res.Sites = append(res.Sites, Sitesurvery_info {
//		Ssid 		: "gj_test2",
//		Signal 		: 123,
//		Mode		: "infrastructure",
//		Channel 	: 1,
//		Auth		: "WPA2-PSK",
//		Encryption 	: "AES",
//	})
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
