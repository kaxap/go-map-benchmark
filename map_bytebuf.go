package main

type MapByteBuf struct {
	N      int
	m      map[int32]int
	buf    []byte
	free   []int
	offset int
}

func NewMapByteBuf(N int) Map {
	it := Item{}
	size := it.Size()
	return &MapByteBuf{buf: make([]byte, N*size),
		m:      make(map[int32]int, N),
		N:      N,
		free:   make([]int, 0, N/8),
		offset: 0,
	}
}

func (m *MapByteBuf) Get(i int32) *Item {
	return DeserializeItem(m.buf[m.m[i]:])
}

func (m *MapByteBuf) Set(i int32, it *Item) {
	if len(m.free) > 0 && len(m.free) >= m.N/8 {
		ind := m.free[len(m.free)-1]
		m.m[i] = ind
		it.Serialize(m.buf[ind:])
		m.free = m.free[:len(m.free)-1]
		return
	}
	m.m[i] = m.offset
	it.Serialize(m.buf[m.offset:])
	m.offset += it.Size()
}

func (m *MapByteBuf) Update(i int32, a int, b int) {
	UpdateItem(m.buf[m.m[i]:], a, b)
}

func (m *MapByteBuf) Delete(i int32) {
	m.free = append(m.free, m.m[i])
	delete(m.m, i)

}
