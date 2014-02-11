package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/trhura/went/dirmap"
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

	path := os.Args[1]
	Some(get_strategies(), path)
}

type StrategyFunc func(string) bool

func get_strategies () []StrategyFunc {
	strategies := make([]StrategyFunc, 0)

	strategies = append(strategies, func (path string) bool {
		if path == "." {
			cwd, err := os.Getwd()
			panic_on_error(err)
			basename := filepath.Base(cwd)

			if recentpath := GetRecentPath(basename); recentpath != "" {
				Chdir(recentpath)

				d := GetDirMap()
				d.ShiftRight(basename)
				d.Save(GetRecentDbPath())
				return true
			}
		}
		return false
	})

	strategies = append(strategies, func (path string) bool {
		if all_characters_are(path, '.') && len(path) > 1 {
			parent, err  := os.Getwd()
			ups := len(path) - 1
			panic_on_error(err)

			for i := ups; i > 0  && IsDirExists(parent) ; i-- {
				parent = filepath.Dir(parent)
			}

			SavePath(parent)
			Chdir(parent)
			return true
		}
		return false
	})

	strategies = append(strategies, func (path string) bool {
		abspath, err := filepath.Abs(path)
		panic_on_error(err)

		if IsDirExists(abspath) {
			SavePath(abspath)
			Chdir(abspath)
			return true
		}
		return false
	})

	strategies = append(strategies, func (path string) bool {
		basename := filepath.Base(path)
		if recentpath := GetRecentPath(basename); recentpath != "" {
			Chdir(recentpath)

			d := GetDirMap()
			d.ShiftRight(basename)
			d.Save(GetRecentDbPath())
			return true
		}

		return false
	})

	strategies = append(strategies, func (path string) bool {
		if IsDirExists(path) { SavePath(path) }
		Chdir(path)
		return true
	})

	return strategies
}

func Chdir(path string) {
	fmt.Println(path)
}

func Some(functions []StrategyFunc, path string) {
	for _, f := range functions {
		if ret := f(path); ret == true {
			break
		}
	}
}

func GetRecentPath(path string) string {
	d := GetDirMap()
	return d.Get(filepath.Base(path))
}

func SavePath(fullpath string) {
	basename := filepath.Base(fullpath)
	d := GetDirMap()
	d.Add(basename, fullpath)
	d.Save(GetRecentDbPath())
}

var _dirmap  *dirmap.DirMap
func GetDirMap() *dirmap.DirMap {
	if _dirmap == nil {
		_dirmap = dirmap.LoadDirMap(GetRecentDbPath())
	}
	return _dirmap
}

func GetRecentDbPath() string {
	return filepath.Join(os.Getenv("HOME"), ".went.recentf")
}

func IsDirExists (path string) bool {
	info, err := os.Stat(path)
	return (err == nil && info.IsDir())
}
