package dirmap

type OrderedCapSet struct {
	slice []interface{}
}

func NewOrderedCapSet() (r *OrderedCapSet) {
	r = new(OrderedCapSet)
	r.slice = make([]interface{}, 0)
	return
}

func (r *OrderedCapSet) Len() int {
	return len(r.slice)
}

func (r *OrderedCapSet) IndexOf (elem interface{}) int {
	for i, item := range r.slice {
		if (item == elem) {
			return i
		}
	}
	return -1
}

func (r *OrderedCapSet) Get(i int) (elem interface{}) {
	if (i >= 0 && i < len(r.slice)) {
		elem = r.slice[i]
	}
	return
}

func (r *OrderedCapSet) GetAll() []interface{} {
	return r.slice
}

func (r *OrderedCapSet) RemoveAt(i int) {
	if (i >= 0) {
		before := r.slice[:i]
		after := r.slice[i+1:]
		r.slice = append(before, after...)
	}
}

const sizelimit = 3
func (r *OrderedCapSet) resize() {
	if(r.Len() > sizelimit) {
		r.slice = r.slice[:sizelimit]
	}
}

func (r *OrderedCapSet) ShiftRight () {
	leftmost := r.slice[0]
	r.RemoveAt(0)
	r.slice = append(r.slice, leftmost)
}

func (r *OrderedCapSet) Prepend(elem interface{}) {
	newslice := []interface{} {elem}
	r.slice = append(newslice, r.slice...)
	r.resize()
}

func (r *OrderedCapSet) Push(elem interface{}) {
	r.RemoveAt(r.IndexOf(elem))
	r.Prepend(elem)
}
