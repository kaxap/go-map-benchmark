package main

type MapSliceStruct struct {
	N    int
	m    map[int32]int
	s    []Item
	free map[int]struct{}
}

func NewMapSliceStruct(N int) Map {
	return &MapSliceStruct{N,
		make(map[int32]int, N),
		make([]Item, 0, N),
		make(map[int]struct{}, N/8),
	}
}

func (m *MapSliceStruct) Get(i int32) *Item {
	it := m.s[m.m[i]]
	return &it
}

func (m *MapSliceStruct) Set(i int32, it *Item) {
	if len(m.free) == m.N/8 {
		for ind := range m.free {
			m.m[i] = ind
			m.s[ind] = *it
			delete(m.free, ind)
			return
		}
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
	ind := m.m[i]
	delete(m.m, i)
	m.free[ind] = struct{}{}
	// m.s[ind] = nil
}
