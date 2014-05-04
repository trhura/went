package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"github.com/trhura/went/dirmap"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		/* Without arguments, go to home directory */
		ShellBuiltinCd(os.Getenv("HOME"))
		return
	}

	/* Join the args with space */
	dir := strings.Join(args[0:], " ")
	TryStrategies(GetCdStrategies(), dir)

	d := GetRecentlyVisitedDb()
	d.Save(GetRecentlyVisitedDbPath())
}

/**
 * A func which take a relative/absolute path, and try to change cwd
 * return true if succesfully changed cwd, otherwise false
 */
type CdStrategyFunc func(string) bool

/**
 * return a list of strategies to change current working directory
 */
func GetCdStrategies() []CdStrategyFunc {
	strategies := make([]CdStrategyFunc, 0)

	strategies = append(strategies, func (path string) bool {
		/* If the passed param is `.` */
		if path == "." {
			cwd, err := os.Getwd()
			PanicOnError(err)
			basename := filepath.Base(cwd)

			if recentpath := GetNextRecentlyVisitedPath(basename); recentpath != "" {
				ShellBuiltinCd(recentpath)
				return true
			}
		}
		return false
	})

	strategies = append(strategies, func (path string) bool {
		/* If the passed path contains all `.`, chdir into
		 * corresponding parent directory.
		 * `..` => parent folder
		 * `...` => parent's parent fold and so on
		 */
		if AllCharsAre(path, '.') && len(path) > 1 {
			parent, err  := os.Getwd()
			PanicOnError(err)

			parentcount := len(path) - 1
			for i := parentcount; i > 0  && IsDirExists(parent) ; i-- {
				parent = filepath.Dir(parent)
			}

			AddPathToRecentlyVisited(parent)
			ShellBuiltinCd(parent)
			return true
		}
		return false
	})

	strategies = append(strategies, func (path string) bool {
		/**
		 * If the given path exists and is an absolute path, or
		 * If the given dir exists in current directory,
		 */
		abspath, err := filepath.Abs(path)
		PanicOnError(err)

		if IsDirExists(abspath) {
			AddPathToRecentlyVisited(abspath)
			ShellBuiltinCd(abspath)
			return true
		}
		return false
	})

	strategies = append(strategies, func (path string) bool {
		basename := filepath.Base(path)
		if recentpath := GetRecentlyVisitedPath(basename); recentpath != "" {
			ShellBuiltinCd(recentpath)
			return true
		}

		return false
	})

	strategies = append(strategies, func (path string) bool {
		/**
		 * if no other strategy works,
		 * just use the shell's builtin cd
		 */
		ShellBuiltinCd(path)
		return true
	})

	return strategies
}

/**
 * Print out the path, which will be piped
 * into shell builtin cd
 */
func ShellBuiltinCd(path string) {
	fmt.Println(path)
}

/**
 * Take a list of strategy funcs, iterate and evaluate each
 * function in order, until one of the return `true`
 */
func TryStrategies(functions []CdStrategyFunc, path string) {
	for _, f := range functions {
		if ret := f(path); ret == true {
			break
		}
	}
}

/**
 * Query the recently visited dir, and return the next path
 */
func GetNextRecentlyVisitedPath(basename string) string {
	d := GetRecentlyVisitedDb()
	d.ShiftRight(basename)
	return GetRecentlyVisitedPath(basename)
}

/**
 * Query the recently visited dir, and return the path
 */
func GetRecentlyVisitedPath(basename string) string {
	d := GetRecentlyVisitedDb()
	recentpath := d.Get(basename)

	for !IsDirExists(recentpath) && recentpath != "" {
		d.Remove(basename)
		recentpath = d.Get(basename)
	}

	return recentpath
}

/**
 * Saved the path in recently visited db, using its
 * basename as key
 */
func AddPathToRecentlyVisited(fullpath string) {
	basename := filepath.Base(fullpath)
	d := GetRecentlyVisitedDb()
	d.Add(basename, fullpath)
}

/**
 * Cache and return a DirMap of recently visited dirs
 */
var _dirmap  *dirmap.DirMap
func GetRecentlyVisitedDb() *dirmap.DirMap {
	if _dirmap == nil {
		_dirmap = dirmap.LoadDirMap(GetRecentlyVisitedDbPath())
	}
	return _dirmap
}

/**
 * return path to storing the recently visited paths db
 */
func GetRecentlyVisitedDbPath() string {
	return filepath.Join(os.Getenv("HOME"), ".went.recentf")
}

/**
 * return true if path exists and is a directory, otherwise false
 */
func IsDirExists(path string) bool {
	info, err := os.Stat(path)
	return (err == nil && info.IsDir())
}

/**
 * helper function to panic on error
 */
func PanicOnError(err error) {
	if err == nil {
		return
	}
	panic(err)
}

/**
 * return true if all characters in `s` are `c`, otherwise false
 */
func AllCharsAre(s string, c rune) bool {
	for _, sc := range s {
		if sc != c {
			return false
		}
	}
	return true
}
