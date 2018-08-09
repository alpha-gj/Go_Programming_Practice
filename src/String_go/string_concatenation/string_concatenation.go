package main

import (
       "fmt"
       "encoding/json"
	   "os/exec"
	   "strings"
	   "strconv"
)


func get_result_string() string {
	command := exec.Command("echo", "30")

	buf, err := command.Output()
	if err != nil {
		fmt.Println(err)
	}                                                       

	output := strings.Replace(string(buf), "\n", "", -1)

	return output

}

func get_result_width() string {
	command := exec.Command("echo", "200")

	buf, err := command.Output()
	if err != nil {
		fmt.Println(err)
	}                                                       

	output := strings.Replace(string(buf), "\n", "", -1)

	return output

}

func get_result_hight() string {
	command := exec.Command("echo", "100")

	buf, err := command.Output()
	if err != nil {
		fmt.Println(err)
	}                                                       

	output := strings.Replace(string(buf), "\n", "", -1)

	return output

}

func get_result_int_by_string() string {
	command := exec.Command("echo", "10000")

	buf, err := command.Output()
	if err != nil {
		fmt.Println(err)
	}                                                       

	output := strings.Replace(string(buf), "\n", "", -1)

	return output

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

func MustAtoi(s string) int {
    result_int, err := strconv.Atoi(s)
    if err == nil {
        fmt.Println(result_int)
    }
    return result_int
}                                                

func main() {

// OK
//	n := Video_setting{ "0.9", 1, "H264", "1920x1080", 30, "CBR", 4096, 30 }

// OK
	
	//bitrate, _ := strconv.Atoi(string(get_result_int_by_string()))
	n := Video_setting{ 
		"0.9",
		1,
		"H264",
		fmt.Sprintf("%sx%s", get_result_hight(), get_result_width()),
		30,
		"CBR", 
		MustAtoi(get_result_int_by_string()),
		get_result_string(),
	}

	a, _ := json.Marshal(n)
	fmt.Println(string(a))
}

