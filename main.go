/**
	Go Console Diary 1.0 by Jozsef Sallai
	(c) 2018 All rights reserved.
**/

package main

import (
	"fmt"
	"os"
)

func main() {

	appname := os.Args[0] // The name of the executable file

	// Get and check for arguments
	args := os.Args[1:]
	if len(os.Args) == 1 {
		fmt.Print("No arguments provided. Use \"", appname, " help\".")
		return
	}

	switch (args[0]) {
		case "write":	// If the correct argument was provided, call the write function
			if len(args) == 2 && (args[1] == "dream" || args[1] == "day") {
				write(args[1])
				return
			} else {
				fmt.Println("Usage:", appname, "write dream|day")
				return
			}
		case "read":	// If the correct argument was provided, and all three arguments are present, call the read function
			if len(args) == 3 && (args[1] == "dream" || args[1] == "day") {
				read(args[1], args[2])
				return
			} else {
				fmt.Println("Usage:", appname, "read dream|day MM-DD-YYYY")
				return
			}
		case "help":	// Call the help function
			help();
			return;
	}

}