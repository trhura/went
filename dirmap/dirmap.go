package dirmap

type DirMap struct {
	dirmap map[string] *OrderedSet
}

func New() *DirMap {
	d := new(DirMap)
	return d
}

func (d *DirMap) Add(basename string, fullpath string) {
	if _, exists := d.dirmap[basename]; exists == false {
		d.dirmap[basename] = NewOrderedSet()
	}

	d.dirmap[basename].Push(fullpath)
}

func (d *DirMap) Has(basename string) bool {
	_, exists := d.dirmap[basename]
	return exists
}

func (d *DirMap) Get(basename string) (path string) {
	if _, exists := d.dirmap[basename]; exists == true {
		//return d.dirmap[basename].Get(0)
	}
	return path
}
