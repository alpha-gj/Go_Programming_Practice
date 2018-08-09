package main

import "fmt"
import "os/exec"
import "strings"
import "encoding/json"

//fmt.Fprintln(w, "<site>\n")
//fmt.Fprintln(w, "<ssid>GJ_DB</ssid>\n")
//fmt.Fprintln(w, "<bssid>c8:be:19:69:d2:df</bssid>\n")
//fmt.Fprintln(w, "<signal>-43</signal>\n")
//fmt.Fprintln(w, "<rssi>0</rssi>\n")
//fmt.Fprintln(w, "<type>Infra</type>\n")
//fmt.Fprintln(w, "<channel>1</channel>\n")
//fmt.Fprintln(w, "<auth>WPA2-PSK</auth>\n")
//fmt.Fprintln(w, "<encrypt>AES</encrypt>\n")

type Sitesurvery_info struct {
	Ssid        string  `json:"ssid"`
	Signal      int 	`json:"signal"`
	Mode		string  `json:"mode"`
	Channel     int		`json:"channel"` 
	Auth		string  `json:"auth"`	
	Encryption	string  `json:"encryption"` 
}

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

func main() {


	sitesurvey_result := execSystem("cat ./test | tail -n +3")
    fmt.Println(sitesurvey_result)
	
//    s := strings.Split(sitesurvey_result, " ")
//    bssid, frequency, signal_level, flags, ssid := s[0], s[7], s[11], s[16], s[17]
//    fmt.Println(bssid, frequency, signal_level, flags, ssid)

	a, _ := json.Marshal(n)
	fmt.Println(string(a))

}
