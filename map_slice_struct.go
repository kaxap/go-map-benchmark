package main

type MapSliceStruct struct {
	N    int
	m    map[int32]int
	s    []Item
	free []int
}

func NewMapSliceStruct(N int) Map {
	return &MapSliceStruct{N,
		make(map[int32]int, N),
		make([]Item, 0, N),
		make([]int, 0, N/8),
	}
}

func (m *MapSliceStruct) Get(i int32) *Item {
	it := m.s[m.m[i]]
	return &it
}

func (m *MapSliceStruct) Set(i int32, it *Item) {
	if len(m.free) > 0 && len(m.free) >= m.N/8 {
		ind := m.free[len(m.free)-1]
		m.m[i] = ind
		m.s[ind] = *it
		m.free = m.free[:len(m.free)-1]
		return
	}
	m.m[i] = len(m.s)
	m.s = append(m.s, *it)
}

func (m *MapSliceStruct) Update(i int32, a int, b int) {
	it := m.s[m.m[i]]
	it.a = a
	it.b = b
}

func (m *MapSliceStruct) Delete(i int32) {
	m.free = append(m.free, m.m[i])
	delete(m.m, i)
}
