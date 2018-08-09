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
	
    // 执行系统命令
    // 第一个参数是命令名称
    // 后面参数可以有多个，命令参数
    cmd := exec.Command("cat", "./test")
    // 获取输出对象，可以从该对象中读取输出结果
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }
    // 保证关闭输出流
    defer stdout.Close()
    // 运行命令
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }

	fileScanner := bufio.NewScanner(stdout)

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
