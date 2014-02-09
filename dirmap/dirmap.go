package dirmap

type DirMap struct {
	dirmap map[string] OrderedSet
}

func New() *DirMap {
	d := new(DirMap)
	return d
}
