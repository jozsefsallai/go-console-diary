package main

import (
	"fmt"
	"os"
)

func help() {

	appname := os.Args[0]	// Get the name of the executable file

	// Generate help string
	helpstring := "Console Diary 1.0 - Help\n\n"
	helpstring += "Available commands:\n"
	helpstring += "(*) write - writes dream or diary entry to file\n"
	helpstring += "    Usage: " + appname + " write dream|day\n"
	helpstring += "(*) read - print out an entry from a specific date\n"
	helpstring += "    Usage: " + appname + " read dream|day MM-DD-YYYY"
	
	// Print help string
	fmt.Println(helpstring)

}