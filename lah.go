package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

const (
	havingPathNo = 1
	havingPath   = 2
)

// Get current directory
func directory() string {
	var curDir string
	count := len(os.Args)

	switch {
	case count == havingPath:
		curDir = os.Args[1]
	case count == havingPathNo:
		curDir = "./"
	default:
		displayUsage()
		os.Exit(0)
	}
	return curDir
}

func displayUsage() {
	usage := `Usage: lah [FILE]

List information about the FILEs (the current directory by default).

[FILE] Directory you want to list files for`
	fmt.Println(usage)
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
	// get the directory
	dir := directory()

	// List the file
	files := fileList(dir)

	// Process the output
	fdis := process(files)

	// Display it
	display(*fdis)
}
