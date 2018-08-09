package main

import (
       "fmt"
       "encoding/json"
	   "os/exec"
)

func get_result_int() int {
	return 100000
}

func get_result_string() string {
	command := exec.Command("echo", "-n", "30")

	buf, err := command.Output()
	if err != nil {
		fmt.Println(err)
	}                                                       

	return string(buf)

//	var buf string = "30"
//	return string(buf)
}

type Video_setting struct {
		Vprofileformat	string  `json:"vprofileformat"`
		Profileid 	 	int		`json:"profileid"`
		Codec 			string  `json:"codec"`
		Resolution		string  `json:"resolution"`
		Framerate 		int 	`json:"framerate"`
		Qualitymode		string 	`json:"qualitymode"`
		Bitrate  		int		`json:"bitrate"`
		Goplength 		string	`json:"goplength"`
}

func main() {

// OK
//	n := Video_setting{ "0.9", 1, "H264", "1920x1080", 30, "CBR", 4096, 30 }

// OK
	n := Video_setting{ "0.9",
	1,
	"H264",
	"1920x1080",
	get_result_int(),
	"CBR", 
	4096, 
	get_result_string(),
}
	   a, _ := json.Marshal(n)
       fmt.Println(string(a))
}

