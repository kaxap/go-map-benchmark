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
	// 196/179

	//buf[0] = byte(i.a >> 56)
	//buf[1] = byte(i.a >> 48)
	//buf[2] = byte(i.a >> 40)
	//buf[3] = byte(i.a >> 32)
	//buf[4] = byte(i.a >> 24)
	//buf[5] = byte(i.a >> 16)
	//buf[6] = byte(i.a >> 8)
	//buf[7] = byte(i.a)
	//
	//buf[8] = byte(i.b >> 56)
	//buf[9] = byte(i.b >> 48)
	//buf[10] = byte(i.b >> 40)
	//buf[11] = byte(i.b >> 32)
	//buf[12] = byte(i.b >> 24)
	//buf[13] = byte(i.b >> 16)
	//buf[14] = byte(i.b >> 8)
	//buf[15] = byte(i.b)
	//
	//buf[16] = byte(i.c >> 56)
	//buf[17] = byte(i.c >> 48)
	//buf[18] = byte(i.c >> 40)
	//buf[19] = byte(i.c >> 32)
	//buf[20] = byte(i.c >> 24)
	//buf[21] = byte(i.c >> 16)
	//buf[22] = byte(i.c >> 8)
	//buf[23] = byte(i.c)
	//
	//buf[24] = byte(i.d >> 56)
	//buf[25] = byte(i.d >> 48)
	//buf[26] = byte(i.d >> 40)
	//buf[27] = byte(i.d >> 32)
	//buf[28] = byte(i.d >> 24)
	//buf[29] = byte(i.d >> 16)
	//buf[30] = byte(i.d >> 8)
	//buf[31] = byte(i.d)

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
