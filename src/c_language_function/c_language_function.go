package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <dlfcn.h>
#include <sys/types.h>
#include <stdint.h>
#include <errno.h>
#include <string.h>

void set_popen_cmd(char* cmd) {

	FILE *fp = popen(cmd, "r");

	do {

		if (fp == NULL) {
			return;
		}

	} while(false);

	if (fp != NULL) {
		pclose(fp);
	}
}

 int add(int a, int b) {
     return a + b;
 }
*/
import "C"

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

printfCgiResult

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析参数，默认是不会解析的
    fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	fmt.Printf("hello world\n")

	var wlan_ap_cmd string = "echo \"321\" > hello ; echo \"1234567\" >> hello"

	C.set_popen_cmd(C.CString(wlan_ap_cmd))
}

func main() {

    http.HandleFunc("/", sayhelloName) //设置访问的路由
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

    a := C.int(1)
    b := C.int(2)
    value := C.add(a, b)
    fmt.Printf("%v\n", value)
    fmt.Printf("%v\n", int(value))
}
