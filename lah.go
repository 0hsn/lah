package main

import (
	"io/fs"
	"io/ioutil"
)

// Get current directory
func directory() (string, error) {
	currDir := "./"
	return currDir, nil
}

// Get list of directory files
func fileList(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err.Error())
	}

	return files
}

func main() {
	// @todo handle error in future when there is error
	dir, _ := directory()

	// List the file
	files := fileList(dir)

	// Process the output
	// fdis :=
	process(files)

	// Display it
	// output()
}
