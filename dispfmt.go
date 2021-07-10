package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
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
	fsize  string // file size
	fsunit string // file size unit
	fcount string // file count in dir
}

func processOutput(files []fs.FileInfo) *[]fileDisplayInfo {
	var fdis []fileDisplayInfo

	for _, file := range files {
		fdi := new(fileDisplayInfo)

		loadName(file, fdi)
		loadPerm(file, fdi)
		loadUser(file, fdi)
		loadGroup(file, fdi)
		loadSize(file, fdi)
		loadFileCount(file, fdi)

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

// load file size and unit
func loadSize(file fs.FileInfo, fdi *fileDisplayInfo) {
	fsize := file.Sys().(*syscall.Stat_t).Size
	fsizef32 := float32(fsize)

	sz := [4]string{"B", "KB", "MB", "GB"}

	var i int
	for i = 0; i <= len(sz); i++ {
		if fsizef32 > 1000 {
			fsizef32 /= 1000
			continue
		}
		break
	}

	if fsizef32 == float32(int32(fsizef32)) {
		fdi.fsize = fmt.Sprint(int32(fsizef32))
	} else {
		fdi.fsize = fmt.Sprintf("%.1f", fsizef32)
	}

	fdi.fsunit = sz[i]
}

// load file count if dir
func loadFileCount(file fs.FileInfo, fdi *fileDisplayInfo) {
	if !file.IsDir() {
		fdi.fcount = "0"
		return
	}

	files, _ := ioutil.ReadDir(file.Name())
	fdi.fcount = fmt.Sprint(len(files))
}
