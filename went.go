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

type StrategyFunc func(string) bool

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		/* Accept exactly one argument */
		fmt.Printf("Usage: %s [dir]\n", os.Args[0])
		return
	}

	if len(args) == 0 {
		/* Without arguments, go to home directory */
		Chdir(os.Getenv("HOME"))
		return
	}

	strategies := make([]StrategyFunc, 0)
	strategies = append(strategies, func (path string) bool {
		if path[0] == '.' {
			switch {
			case path == ".":
				/* FIXME: circle */
				return true

			case all_characters_are(path, '.'):
				cwd, err  := os.Getwd()
				panic_on_error(err)

				ups := len(path) - 1
				parent := cwd
				for i := ups; i > 0  && IsDirExists(parent) ; i-- {
				 	parent = filepath.Dir(parent)
				}

				Chdir(parent)
				return true

			default:
				/* otherwise, use default cd implementation */
				Chdir(path)
				return true
			}
		}
		return false
	})

	path := os.Args[1]
	Some(strategies, path)
}

func Some(functions []StrategyFunc, path string) {
	for _, f := range functions {
		if ret := f(path); ret == true {
			break
		}
	}
}

const dbfilename = "went.csv"
func savePath(path string) {
	basename := filepath.Base(path)
	d := dirmap.LoadDirMap(dbfilename)
	d.Add(basename, path)
	d.Save(dbfilename)
}

func Chdir(path string) {
	fmt.Println(path)
}

func IsDirExists (path string) bool {
	info, err := os.Stat(path)
	return (err == nil && info.IsDir())
}
