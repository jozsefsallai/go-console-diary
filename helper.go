package main

import (
	"fmt"
	"os"
	"github.com/mitchellh/go-homedir"
)

func help() {

	// Get user's home directory
	homepath, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	appname := os.Args[0]	// Get the name of the executable file

	// Generate help string
	helpstring := "Console Diary 1.0 - Help\n\n"
	helpstring += "Available commands:\n"
	helpstring += "(*) write - writes dream or diary entry to file\n"
	helpstring += "    Usage: " + appname + " write dream|day\n"
	helpstring += "(*) read - print out an entry from a specific date\n"
	helpstring += "    Usage: " + appname + " read dream|day word\n"
	helpstring += "           search entries for a specific day: " + appname + " read dream|day M-D-YYYY\n\n"
	helpstring += "The files are being saved in your home directory (" + homepath + ") and are called \"dream.txt\" and \"day.txt\".\n"
	
	// Print help string
	fmt.Println(helpstring)

}