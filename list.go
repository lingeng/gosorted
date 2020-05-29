package gosorted

import (
	"sort"
)

type List []EleInterface

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	return l[i].Less(l[j])
}

func (l List) Insert(ei EleInterface) (int, List) {
	for i, item := range l {
		if !item.Less(ei) {
			l = append(l, ei)
			copy(l[i+1:], l[i:])
			l[i] = ei
			return i, l
		}
	}
	return len(l), append(l, ei)
}

func (l List) Sort() {
	sort.Sort(l)
}
