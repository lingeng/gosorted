package gosorted

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const numQuantity int = 100

func TestListInsert(t *testing.T) {
	var l List
	rand.Seed(time.Now().Unix())
	var index int
	for i := 0; i < numQuantity; i++ {
		ele := Ele{rand.Intn(100)}
		index, l = l.Insert(ele)
		assert.Truef(t, l[index].Equal(ele), "Wrong index: %d, %v, %v", index, ele, l)
	}
	for i := 0; i < l.Len()-1; i++ {
		assert.Truef(t, (l[i].Less(l[i+1]) || l[i].Equal(l[i+1])), "l[%d] is not less than l[%d]: %d, %d", i, i+1, l[i], l[i+1])
	}
}

func TestListSort(t *testing.T) {
	var l List
	rand.Seed(time.Now().Unix())
	for i := 0; i < numQuantity; i++ {
		l = append(l, Ele{rand.Intn(100)})
	}
	l.Sort()
	for i := 0; i < l.Len()-1; i++ {
		assert.Truef(t, (l[i].Less(l[i+1]) || l[i].Equal(l[i+1])), "l[%d] is not less than l[%d]: %d, %d", i, i+1, l[i], l[i+1])
	}
}
