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

	// Create cmd
	char getDBCmd[BUF_SIZE] = {};
	snprintf(getDBCmd, sizeof(getDBCmd), "getprop persist.%s.%s", dbGroupName, dbItemName);

	//Use the cmd to get data from db
	FILE *fp = popen(getDBCmd, "r");
	char getResult[BUF_SIZE] = {};

	do {

		if (fp == NULL) {
			printf("fp is error\n");
			return;
		}

		fgets(getResult, BUF_SIZE, fp);

	} while(false);

	if (fp != NULL) {
		pclose(fp);
	}

	//Print data, cgiItemName=GetResult
	printf("%s=%s\n", cgiItemName, getResult);
}
*/
import "C"

func main() {
	C.printfCgiResult(C.CString("enable"), 
					  C.CString("Wireless"),
					  C.CString("STAEnable"))
}
