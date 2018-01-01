package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/mitchellh/go-homedir"
)

func read(option string, word string) {

	appname := os.Args[0]	// Get the name of the executable file

	// Get user's home directory
	homepath, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	// Check if file exists
	if _, err := os.Stat(homepath + "/" + option + ".txt"); os.IsNotExist(err) {
		fmt.Print("The ", option, ".txt file doesn't exist yet! Use \"", appname, " write ", option, "\" to create it.\n")
		return
	}

	// Open the file
	file, err := os.Open(homepath + "/" + option + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Initialize new buffer
	sc := bufio.NewScanner(file)
	exists := false		// Checks if an entry that contains the provided word exists

	// Iterate through the rows and print out the rows that contain the said word
	for sc.Scan() {
		if (strings.Contains(sc.Text(), word)) {
			exists = true
			fmt.Println(sc.Text())
		}
	}

	// If there is no entry for the given word, print out a message
	if (!exists) {
		fmt.Println("There's no entry containing that word!\n")
	}

}