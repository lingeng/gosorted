package gosorted

type Map struct {
	eleMap  map[string]EleInterface
	eleList List
	keys    []string
}

func NewMap(eiMap map[string]EleInterface) *Map {
	m := &Map{
		eleMap:  make(map[string]EleInterface),
		eleList: List{},
		keys:    []string{},
	}
	for k, v := range eiMap {
		m.Set(k, v)
	}
	return m
}

func (km *Map) Set(key string, value EleInterface) {
	km.eleMap[key] = value
	var index int
	index, km.eleList = km.eleList.Insert(value)
	km.keys = append(km.keys, "")
	copy(km.keys[index+1:], km.keys[index:])
	km.keys[index] = key
}

func (km *Map) Len() int {
	return len(km.keys)
}

func (km *Map) Get(key string) EleInterface {
	return km.eleMap[key]
}

func (km *Map) Get2(key string) (EleInterface, bool) {
	value, ok := km.eleMap[key]
	return value, ok
}

func (km *Map) GetByIndex(index int) (string, EleInterface) {
	key := km.keys[index]
	value := km.eleMap[key]
	return key, value
}

func (km *Map) Keys() []string {
	return km.keys
}

func (km *Map) Front() (string, EleInterface) {
	value := km.eleList[0]
	return km.keys[0], value
}

func (km *Map) Back() (string, EleInterface) {
	index := km.Len() - 1
	value := km.eleList[index]
	return km.keys[index], value
}

func (km *Map) Remove(key string) (ei EleInterface) {
	ei = km.eleMap[key]
	delete(km.eleMap, key)
	var index int
	var k string
	for index, k = range km.keys {
		if key == k {
			copy(km.keys[index:], km.keys[index+1:])
			km.keys = km.keys[:len(km.keys)-1]

			copy(km.eleList[index:], km.eleList[index+1:])
			km.eleList = km.eleList[:len(km.eleList)-1]
		}
	}
	return
}
