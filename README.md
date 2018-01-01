# go-console-diary

This is a small Go project I made to create a "log" of my dreams and also use it as a diary. It creates two files in the user's home directory: `dream.txt` (for the dream log) and `day.txt` (for the diary). 

The program uses the [mitchellh/go-homedir](https://github.com/mitchellh/go-homedir) package.

**Available commands:**
 - **write dream|day** - write entry to the log
 - **read dream|day word** - search entries that contain a specific word (this can also be used to search for entries from a certain day)
 - **help** - display whatever I just said here.
