package dirmap

type OrderedSet struct {
	slice []interface{}
}

func New() (r *OrderedSet) {
	r = new(OrderedSet)
	r.slice = make([]interface{}, 0)
	return
}

func (r *OrderedSet) Len() int {
	return len(r.slice)
}

func (r *OrderedSet) IndexOf (elem interface{}) int {
	for i, item := range r.slice {
		if (item == elem) {
			return i
		}
	}
	return -1
}

func (r *OrderedSet) Get(i int) (elem interface{}) {
	if (i >= 0 && i < len(r.slice)) {
		elem = r.slice[i]
	}
	return
}

func (r *OrderedSet) RemoveAt(i int) {
	if (i >= 0) {
		before := r.slice[:i]
		after := r.slice[i+1:]
		r.slice = append(before, after...)
	}
}

func (r *OrderedSet) Prepend(elem interface{}) {
	newslice := make([]interface{}, 1)
	newslice[0] = elem
	r.slice = append(newslice, r.slice...)
}

func (r *OrderedSet) Push(elem interface{}) {
	i := r.IndexOf(elem)
	r.RemoveAt(i)
	r.Prepend(elem)
}
