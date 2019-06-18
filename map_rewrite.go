package main

type MapRewrite struct {
	m map[int32]Item
}

func NewMapRewrite(N int) Map {
	return &MapRewrite{make(map[int32]Item, N)}
}

func (m *MapRewrite) Get(i int32) *Item {
	r := m.m[i]
	return &r
}

func (m *MapRewrite) Set(i int32, it *Item) {
	m.m[i] = *it
}

func (m *MapRewrite) Update(i int32, a int, b int) {
	it := m.m[i]
	it.a = a
	it.b = b
	m.m[i] = it
}

func (m *MapRewrite) Delete(i int32) {
	delete(m.m, i)
}
