package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"io"
	"io/ioutil"
	"strings"
	"github.com/mitchellh/go-homedir"
)

func write(option string) {

	// Get user's home directory
	homepath, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	// Create a welcome string based on the user's choice
	var welcome string
	if (option == "dream") {
		welcome = "What have you dreamed last night?"
	} else {
		welcome = "What happened to you today?"
	}
	fmt.Println(welcome, "\n")

	// Initialize the buffer
	reader := bufio.NewReader(os.Stdin)
	var input, entry string

	// Get current UTC date
	currentTime := time.Now().UTC()
	date := fmt.Sprintf("%d-%d-%d", currentTime.Month(), currentTime.Day(), currentTime.Year())
	entry = date + ": "

	// Read input and add it to the entry string
	input, _ = reader.ReadString('\n')		
	entry += input

	// If file exists, concatenate the contents of the file with the current entry string
	if _, err := os.Stat(homepath + "/" + option + ".txt"); err == nil {
		bytes, err := ioutil.ReadFile(homepath + "/" + option + ".txt")
		if err != nil {
			panic(err)
		}
		entry = string(bytes) + "\n" + entry
	}

	// Open the text file
	file, err := os.Create(homepath + "/" + option + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Copy the entry string to the file
	_, err = io.Copy(file, strings.NewReader(entry))
	if err != nil {
		panic(err)
	}

}