package dirmap

type DirMap struct {
	dirmap map[string]*OrderedCapSet
}

func NewDirMap() *DirMap {
	d := new(DirMap)
	d.dirmap = make(map[string]*OrderedCapSet)
	return d
}

func (d *DirMap) Add(basename string, fullpath string) {
	if _, exists := d.dirmap[basename]; exists == false {
		d.dirmap[basename] = NewOrderedCapSet()
	}

	d.dirmap[basename].Push(fullpath)
}

func (d *DirMap) Has(basename string) bool {
	_, exists := d.dirmap[basename]
	return exists
}

func (d *DirMap) Get(basename string) (path string) {
	if d.Has(basename) == true {
		return d.dirmap[basename].Get(0).(string)
	}
	return path
}

func (d *DirMap) GetAll(basename string) []string {
	if d.Has(basename) == false {
		return make([]string, 0)
	}

	set := d.dirmap[basename]
	all := make([]string, set.Len())
	for i, item := range set.GetAll() {
		all[i] = item.(string)
	}

	return all
}
