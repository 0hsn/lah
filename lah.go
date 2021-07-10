package main

import (
	"fmt"
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

// Display the file displayable infos
func display(fdis []fileDisplayInfo) {
	for _, fdi := range fdis {
		fmt.Printf("%3s %4s %2s %3s%2s %s:%s %s\n",
			fdi.ftype, fdi.perm, fdi.fcount, fdi.fsize, fdi.fsunit, fdi.user, fdi.group, fdi.name)
	}
}

func main() {
	// @todo handle error in future when there is error
	dir, _ := directory()

	// List the file
	files := fileList(dir)

	// Process the output
	fdis := process(files)

	// Display it
	display(*fdis)
}
