package main

import "encoding/binary"

type Item struct {
	a int
	b int
	c int
	d int
}

func (i *Item) Serialize(buf []byte) {
	binary.BigEndian.PutUint64(buf, uint64(i.a))
	binary.BigEndian.PutUint64(buf[8:], uint64(i.b))
	binary.BigEndian.PutUint64(buf[16:], uint64(i.c))
	binary.BigEndian.PutUint64(buf[24:], uint64(i.d))
}

func DeserializeItem(buf []byte) *Item {
	return &Item{
		a: int(binary.BigEndian.Uint64(buf)),
		b: int(binary.BigEndian.Uint64(buf[8:])),
		c: int(binary.BigEndian.Uint64(buf[16:])),
		d: int(binary.BigEndian.Uint64(buf[24:])),
	}
}

func (i *Item) Size() int {
	return 4 * 8
}

func UpdateItem(buf []byte, a, b int) {
	binary.BigEndian.PutUint64(buf, uint64(a))
	binary.BigEndian.PutUint64(buf[8:], uint64(b))
}
