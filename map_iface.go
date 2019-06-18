package main

type Map interface {
	Get(int32) *Item
	Set(int32, *Item)
	Update(int32, int, int)
	Delete(int32)
}
