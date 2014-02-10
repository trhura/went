package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/trhura/went/dirmap"
)

func check_error(err error) {
	if (err == nil) { return }
	panic(err)
}

func main() {
	args := os.Args[1:]

	/* Accept exactly one argument */
	if (len(args) != 1) {
		fmt.Printf("Usage: %s Path\n", os.Args[0])
		os.Exit(255)
	}

	path :=  os.Args[1]
	/* If the path is relative, get absolute path */
	if (!filepath.IsAbs(path)) {
		pwd, err := os.Getwd()
		check_error (err)
		path = filepath.Join(pwd, path)
	}

	/* Check whether  path exists */
	fileinfo, err := os.Stat(path)
	if (os.IsNotExist(err)) {
		os.Exit(255)
	}

	/* Check whether path is a directory */
	if (!fileinfo.IsDir()) {
		os.Exit(255)
	}

	filename := "went.csv"
	basename := filepath.Base(path)

	d := dirmap.LoadDirMap(filename)
	d.Add(basename, path)
	d.Save(filename)
}
