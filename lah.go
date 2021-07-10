package main

import (
	"io/fs"
	"io/ioutil"
)

// Get current directory
func getDirectoryPath() (string, error) {
	currDir := "./"
	return currDir, nil
}

// Get list of directory files
func getListOfDirFile(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err.Error())
	}

	return files
}

func main() {
	// @todo handle error in future when there is error
	dir, _ := getDirectoryPath()

	// List the file
	files := getListOfDirFile(dir)

	// Process the output
	processOutput(files)

	// Display it
	// output()
}
