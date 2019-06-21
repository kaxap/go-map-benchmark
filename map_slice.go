package main

type MapSlice struct {
	N    int
	m    map[int32]int
	s    []*Item
	free []int
}

func NewMapSlice(N int) Map {
	return &MapSlice{N,
		make(map[int32]int, N),
		make([]*Item, 0, N),
		make([]int, 0, N/8),
	}
}

func (m *MapSlice) Get(i int32) *Item {
	return m.s[m.m[i]]
}

func (m *MapSlice) Set(i int32, it *Item) {
	if len(m.free) > 0 && len(m.free) >= m.N/8 {
		ind := m.free[len(m.free)-1]
		m.m[i] = ind
		m.s[ind] = it
		m.free = m.free[:len(m.free)-1]
		return
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
	m.free = append(m.free, ind)
	m.s[ind] = nil
}
