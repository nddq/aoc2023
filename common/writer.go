package common

import (
	"io/fs"
	"log"
	"os"
)

// WriteToFile writes content to filename, if the file does not exist, it will be created, if it does, content will be appended to the end of the file
func WriteToFile(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		log.Fatal(err)
	}
}
