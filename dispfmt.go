package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"syscall"
)

type fileDisplayInfo struct {
	name   string // name
	ftype  string // file type
	perm   string // perm
	user   string // owner
	group  string // owner group
	fsize  int    // file size
	fsunit string // file size unit
	fcount int    // file count in dir
}

func processOutput(files []fs.FileInfo) *[]fileDisplayInfo {
	var fdis []fileDisplayInfo

	for _, file := range files {
		fdi := new(fileDisplayInfo)

		loadName(file, fdi)
		loadPerm(file, fdi)
		loadUser(file, fdi)
		loadGroup(file, fdi)

		fdis = append(fdis, *fdi)
	}

	fmt.Println(fdis)
	os.Exit(0)

	return &fdis
}

/*
 * formatter utility
 */

// load file name
func loadName(file fs.FileInfo, fdi *fileDisplayInfo) {
	fdi.name = file.Name()
}

// load permission as octal string
func loadPerm(file fs.FileInfo, fdi *fileDisplayInfo) {
	fdi.perm = fmt.Sprintf("%04o", file.Mode().Perm())
}

// load os user name
func loadUser(file fs.FileInfo, fdi *fileDisplayInfo) {
	uid := file.Sys().(*syscall.Stat_t).Uid
	usr, _ := user.LookupId(fmt.Sprint(uid))

	fdi.user = usr.Username
}

// load os group name
func loadGroup(file fs.FileInfo, fdi *fileDisplayInfo) {
	gid := file.Sys().(*syscall.Stat_t).Gid
	grp, _ := user.LookupGroupId(fmt.Sprint(gid))

	fdi.group = grp.Name
}
