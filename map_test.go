package gosorted

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const letters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Ele struct {
	I int
}

func (e Ele) Less(ei EleInterface) bool {
	e2, _ := ei.(Ele)
	return e.I < e2.I
}

func (e Ele) Equal(ei EleInterface) bool {
	e2, _ := ei.(Ele)
	return e.I == e2.I
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randNewMap(length int) *Map {
	nm := NewMap(nil)
	for i := 0; i < length; i++ {
		key := RandString(5)
		value := Ele{rand.Intn(10)}
		nm.Set(key, value)
	}
	return nm
}

func TestMapNewMap(t *testing.T) {
	m := make(map[string]EleInterface)
	m["foo"] = Ele{}
	m["bar"] = Ele{1}
	m[""] = Ele{2}

	nm := NewMap(m)
	for k, v := range m {
		assert.Equalf(t, v, nm.Get(k), "Value is wrong: %v, %v", nm.Get(k), v)
	}

	nm = NewMap(nil)
	assert.Zerof(t, nm.Len(), "Wrong item quantity: %d, should be 0", nm.Len())

	nm.Set("foo", m["foo"])
	assert.Equalf(t, nm.Get("foo"), m["foo"], "Value is wrong: %v, %v\n", nm.Get("foo"), m["foo"])
}

func TestMapSet(t *testing.T) {
	nm := NewMap(nil)
	testMap := make(map[string]EleInterface)
	for i := 0; i < 100; i++ {
		key := RandString(5)
		value := Ele{rand.Intn(10)}
		testMap[key] = value
		nm.Set(key, value)
	}

	assert.Equalf(t, len(testMap), nm.Len(), "map length is %d, but it should be %d", nm.Len(), len(testMap))

	var currKey string
	var currEle EleInterface
	for i := 0; i < nm.Len(); i++ {
		if i == 0 {
			_, currEle = nm.GetByIndex(i)
			continue
		}
		k, v := nm.GetByIndex(i)
		assert.Truef(t, (currEle.Less(v) || currEle.Equal(v)), "%s:%v should be less than %s:%v at index %d", currKey, currEle, k, v, i)
	}
}

func TestMapGet(t *testing.T) {
	nm := NewMap(nil)
	testMap := make(map[string]EleInterface)
	for i := 0; i < 100; i++ {
		key := RandString(5)
		value := Ele{rand.Intn(10)}
		testMap[key] = value
		nm.Set(key, value)
	}

	for k, v := range testMap {
		assert.Truef(t, v.Equal(nm.Get(k)), "value is not equal for the same key: %s, %v, %v", k, v, nm.Get(k))
	}
}

func TestMapLen(t *testing.T) {
	m := make(map[string]EleInterface)
	m["foo"] = Ele{}
	m["bar"] = Ele{1}
	m[""] = Ele{2}

	nm := NewMap(m)
	assert.Equalf(t, 3, nm.Len(), "map length is %d, should be %d", nm.Len(), 3)
}

func TestMapKeys(t *testing.T) {
	nm := NewMap(nil)
	keyMap := make(map[string]struct{})
	for i := 0; i < 50; i++ {
		key := RandString(5)
		value := Ele{rand.Intn(10)}
		nm.Set(key, value)
		keyMap[key] = struct{}{}
	}

	keys := nm.Keys()
	assert.Equalf(t, len(keys), nm.Len(), "length of keys is %d, should be %d", len(keys), nm.Len())
	assert.Equalf(t, len(keys), len(keyMap), "length of keys is %d, should be %d", len(keys), nm.Len())

	for _, v := range keys {
		_, ok := keyMap[v]
		assert.Truef(t, ok, "%s should exist", v)
	}
}

func TestMapFront(t *testing.T) {
	nm := randNewMap(50)
	key, value := nm.Front()

	fk, fv := nm.GetByIndex(0)
	assert.Equalf(t, fk, key, "front key should be %s, not %s", fk, key)
	assert.Equalf(t, fv, value, "front value should be %v, not %v", fv, value)

	length := nm.Len()
	for i := 0; i < length; i++ {
		k, v := nm.GetByIndex(i)
		assert.Truef(t, (value.Less(v) || value.Equal(v)), "%s, %v should less or equal than %s, %v", key, value, k, v)
	}
}

func TestMapBack(t *testing.T) {
	nm := randNewMap(50)
	key, value := nm.Back()

	length := nm.Len()
	bk, bv := nm.GetByIndex(length - 1)
	assert.Equalf(t, bk, key, "back key should be %s, not %s", bk, key)
	assert.Equalf(t, bv, value, "back value should be %v, not %v", bv, value)

	for i := 0; i < length; i++ {
		k, v := nm.GetByIndex(i)
		assert.Truef(t, (v.Less(value) || v.Equal(value)), "%s, %v should less or equal than %s, %v", k, v, key, value)
	}
}

func TestMapRemove(t *testing.T) {
	nm := randNewMap(5)

	length := nm.Len()
	for i := 0; i < length; i++ {
		k, v := nm.GetByIndex(0)
		rv := nm.Remove(k)
		assert.Equal(t, rv, v, "should be equal")
		_, ok := nm.Get2(k)
		assert.True(t, !ok, "should be false")
	}

	assert.Zero(t, nm.Len())
}
