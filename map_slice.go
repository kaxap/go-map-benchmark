package main

type MapSlice struct {
	N    int
	m    map[int32]int
	s    []*Item
	free map[int]struct{}
}

func NewMapSlice(N int) Map {
	return &MapSlice{N,
		make(map[int32]int, N),
		make([]*Item, 0, N),
		make(map[int]struct{}, N/8),
	}
}

func (m *MapSlice) Get(i int32) *Item {
	return m.s[m.m[i]]
}

func (m *MapSlice) Set(i int32, it *Item) {
	if len(m.free) == m.N/8 {
		for ind := range m.free {
			m.m[i] = ind
			m.s[ind] = it
			delete(m.free, ind)
			return
		}
	}
	m.m[i] = len(m.s)
	m.s = append(m.s, it)
}

func (m *MapSlice) Update(i int32, a int, b int) {
	it := m.s[m.m[i]]
	it.a = a
	it.b = b
}

func (m *MapSlice) Delete(i int32) {
	ind := m.m[i]
	delete(m.m, i)
	m.free[ind] = struct{}{}
	m.s[ind] = nil
}
