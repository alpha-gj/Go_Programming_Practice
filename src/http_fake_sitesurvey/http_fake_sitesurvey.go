package main

import "fmt"
import "net/http"
import "log"

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<?xml version=\"1.0\" encoding=\"utf-8\"?>")
	fmt.Fprintln(w, "<?xml-stylesheet type=\"text/xsl\" href=\"adv_survey.xsl\"?>")
	fmt.Fprintln(w, "<site>\n")
	fmt.Fprintln(w, "<ssid>GJ_DB</ssid>")
	fmt.Fprintln(w, "<bssid>c8:be:19:69:d2:df</bssid>")
	fmt.Fprintln(w, "<signal>-43</signal>")
	fmt.Fprintln(w, "<rssi>0</rssi>")
	fmt.Fprintln(w, "<type>Infra</type>")
	fmt.Fprintln(w, "<channel>1</channel>")
	fmt.Fprintln(w, "<auth>WPA2-PSK</auth>")
	fmt.Fprintln(w, "<encrypt>AES</encrypt>")
	fmt.Fprintln(w, "</site>")

	defer r.Body.Close()
}

func main() {
    http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":2000", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

