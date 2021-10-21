package main

import (
	"github.com/withmandala/go-log"
	"os"
	"path"
	"strings"
)
var (
	logger = log.New(os.Stderr)
)
// from https://gobyexample.com/directories
func check(e error) {
	if e != nil {
		logger.Fatal(e)
		os.Exit(1)
	}
}

func line(message string) {
	// I love how bad this looks
	halfTheLine := strings.Repeat("-", 20)
	logger.Info(halfTheLine + message + halfTheLine)
}

func constructPath(directoryName string, fileName string) string {
	return path.Join([]string{ directoryName, fileName }...)
}

func main() {
	logger.Info("1. Create a directory if not exists")
	const directory = "scratch"
	err := os.MkdirAll(directory, os.ModePerm)
	line("Created")

	logger.Info("2. Remove all files from a directory if exists")
	currentDirectory, err := os.Open(directory)
	check(err)
	defer currentDirectory.Close()
	const returnAllFiles = 0
	currentFiles, err := currentDirectory.ReadDir(returnAllFiles)
	check(err)
	for _, file := range currentFiles {
		// To ensure I remember what I did here:
		// []string{directory, file.Name()}...
		//   []string -> Create a new array of type string
		// Instantiate it with two values: directory, file.Name()
		// ... -> Splat the array as arguments to path.Join, as that does not take an array of strings, but an argument
		//   	   list of strings

		filePath := constructPath(directory, file.Name())
		err := os.RemoveAll(filePath)
		check(err)
	}
	line("Directory Purged")

	logger.Info("3. Create a new text file called 'stamp.txt'")
	stampFileName := "stamp.txt"
	contents := []byte ("")
	stampFilePath := constructPath(directory, stampFileName)
	err = os.WriteFile(stampFilePath, contents, os.ModeAppend)
	check(err)
	logger.Info("4. Write the first line 'hello' to the text file.")

	logger.Info("5. Append the second line 'world' to the text file.")

	logger.Info("6. Read the second line 'world' into an input string.")

	logger.Info("7. Print said line to console")

	logger.Info("8. Copy a file to another file")
}