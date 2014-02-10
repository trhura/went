package dirmap

import "os"
import "encoding/csv"

type DirMap struct {
	dirmap map[string]*OrderedCapSet
}

func NewDirMap() *DirMap {
	d := new(DirMap)
	d.dirmap = make(map[string]*OrderedCapSet)
	return d
}

func (d *DirMap) Add(basename string, fullpaths ...string) {
	if _, exists := d.dirmap[basename]; exists == false {
		d.dirmap[basename] = NewOrderedCapSet()
	}

	last := len(fullpaths) - 1
	for i, _ := range fullpaths {
		path := fullpaths[last-i]
		if path != "" {
	 		d.dirmap[basename].Push(path)
		}
	}
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

func check_error(err error) {
	if (err != nil) {
		panic(err)
	}
}

func LoadDirMap (filename string) *DirMap {
	d := NewDirMap()

	f, err := os.Open(filename)
	if (err != nil) { return NewDirMap() }
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	check_error(err)

	for _, record := range records {
		basename := record[0]
		paths := record[1:]
		d.Add(basename, paths...)
	}

	return d
}

func (d *DirMap) Len () int {
	return len(d.dirmap)
}

func (d *DirMap) Save (filename string) {
	f, err := os.Create(filename)
	check_error(err)
	defer f.Close()

	records := make([][]string, 0)
	for k, _ := range d.dirmap {
		record := make([]string, 4)
		record[0] = k

		for i := 0; i < 3; i++ {
			var str string
			item := d.dirmap[k].Get(i)
			if (item == nil) {
				str  = ""
			} else {
				str = item.(string)
			}

			record[i+1] = str
		}

		records = append(records, record)
	}

	writer := csv.NewWriter(f)
	writer.WriteAll(records)
	defer writer.Flush()
}
