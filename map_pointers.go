package main

type MapPointers struct {
	m map[int32]*Item
}

func NewMapPointers(N int) Map {
	return &MapPointers{make(map[int32]*Item, N)}
}

func (m *MapPointers) Get(i int32) *Item {
	return m.m[i]
}

func (m *MapPointers) Set(i int32, it *Item) {
	m.m[i] = it
}

func (m *MapPointers) Update(i int32, a int, b int) {
	it := m.m[i]
	it.a = a
	it.b = b
}

func (m *MapPointers) Delete(i int32) {
	delete(m.m, i)
}
