package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func read(option string, date string) {

	appname := os.Args[0]	// Get the name of the executable file

	// Check if file exists
	if _, err := os.Stat(option + ".txt"); os.IsNotExist(err) {
		fmt.Print("The ", option, ".txt file doesn't exist yet! Use \"", appname, " write ", option, "\" to create it.")
		return
	}

	// Open the file
	file, err := os.Open(option + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Initialize new buffer
	sc := bufio.NewScanner(file)
	exists := false		// Checks if an entry for the given date exists

	// Iterate through the rows and print out the rows that contain the said date
	for sc.Scan() {
		if (strings.Contains(sc.Text(), date)) {
			exists = true
			fmt.Println(sc.Text())
		}
	}

	// If there is no entry for the given date, print out a message
	if (!exists) {
		fmt.Println("There's no entry for that date!")
	}

}