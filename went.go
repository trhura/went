package main

import (
	"fmt"
	"github.com/trhura/went/dirmap"
	"os"
	"path/filepath"
)

func panic_on_error(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func all_characters_are (s string, c rune) bool {
	for _, sc := range s {
		if sc != c {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	/* Accept exactly one argument */
	if len(args) != 1 {
		fmt.Printf("Usage: %s Path\n", os.Args[0])
		os.Exit(255)
	}

	path := os.Args[1]
	/* If the path is absolute, save it in database  and exit*/
	if filepath.IsAbs(path) {
		savePath(path)
		return
	}

	pwd, err := os.Getwd()
	panic_on_error(err)

	fullpath := filepath.Join(pwd, path)
	fileinfo, err := os.Stat(fullpath)

	switch {
	case path == ".":
		/* FIXME: circle */
		break

	case all_characters_are(path, '.'):
		/* FIXME: go parent */
		upcount := len(path) - 1
		for i := upcount; i > 0; i-- {

		}
		break

	case (!os.IsNotExist(err)  && fileinfo.IsDir()):
		/* if the folder exists in current directory */
		os.Chdir(fullpath)
		break

	default:
		/* FIXME: query db */
		break
	}
}

const dbfilename = "went.csv"

func savePath(path string) {
	/* Check whether  path exists */
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Exit(-1)
	}

	/* Check whether path is a directory */
	if !fileinfo.IsDir() {
		os.Exit(-1)
	}

	basename := filepath.Base(path)
	d := dirmap.LoadDirMap(dbfilename)
	d.Add(basename, path)
	d.Save(dbfilename)
}
