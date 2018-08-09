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

#define BUF_SIZE 128
printfCgiResult(char* cgiItemName, char* dbGroupName, char* dbItemName) {

//	char getDBCmd[BUF_SIZE] = {};
//	snprintf(getDBCmd, strlen(getDBCmd), "getprop persist.%s.%s", dbGroupName, dbItemName);

	char* getDBCmd = "echo hi";
	FILE *fp = popen(getDBCmd, "r");
//	FILE *fp = popen("echo on", "r");
	char Buf[BUF_SIZE] = {};

	do {

		if (fp == NULL) {
			printf("fp is error\n");
			return;
		}

		fgets(Buf, BUF_SIZE, fp);
//		while(fgets(Buf, BUF_SIZE, fp) != NULL) {
//		}

	} while(false);

	if (fp != NULL) {
		pclose(fp);
	}

	printf("%s=%s\n", cgiItemName, Buf);
	printf("%s, %s, %s\n", cgiItemName, dbGroupName, dbItemName);
}
*/
import "C"

import (
    "fmt"
//    "net/http"
//    "strings"
//    "log"
)

func main() {

// var Admin string
// Admin = "Administrator"                                        
// w.Header().Set(HTTPHeader.ContentType, HTTPHeader.ContentValue)
// w.Header().Set("Content-Type", "text/plain")

	C.printfCgiResult(C.CString("enable"), 
					  C.CString("Wireless"),
					  C.CString("STAEnable"))

	fmt.Printf("hello\n");
	
}
