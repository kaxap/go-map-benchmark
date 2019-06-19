package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"testing"
	"time"
)

var GCTimes = make(map[string]time.Duration)

func Set(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}
	b.StopTimer()
	runtime.GC()
	GCTimes[fmt.Sprintf("A Set %T", m)] = timeGC()
	runtime.KeepAlive(m)
}

func Get(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Get(int32(i))
	}

	b.StopTimer()
	runtime.GC()
	GCTimes[fmt.Sprintf("B Get %T", m)] = timeGC()
	runtime.KeepAlive(m)
}

func Update(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Update(int32(i), i, i)
	}

	b.StopTimer()
	runtime.GC()
	GCTimes[fmt.Sprintf("C Update %T", m)] = timeGC()
	runtime.KeepAlive(m)
}

func Delete(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(int32(i))
	}

	b.StopTimer()
	runtime.GC()
	GCTimes[fmt.Sprintf("D Delete %T", m)] = timeGC()
	runtime.KeepAlive(m)
}

func SetGet(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
		_ = m.Get(int32(i))
	}

	b.StopTimer()
	runtime.GC()
	GCTimes[fmt.Sprintf("E SetGet %T", m)] = timeGC()
	runtime.KeepAlive(m)
}

func SetDelete(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
		m.Delete(int32(i))
	}

	b.StopTimer()
	runtime.GC()
	GCTimes[fmt.Sprintf("F SetDelete %T", m)] = timeGC()
	runtime.KeepAlive(m)
}

func GetDelete(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Get(int32(i))
		m.Delete(int32(i))
	}

	b.StopTimer()
	runtime.GC()
	GCTimes[fmt.Sprintf("G GetDelete %T", m)] = timeGC()
	runtime.KeepAlive(m)
}

// Set
func BenchmarkMapPointers_Set(b *testing.B) {
	Set(NewMapPointers(b.N), b)
}
func BenchmarkMapRewrite_Set(b *testing.B) {
	Set(NewMapRewrite(b.N), b)
}
func BenchmarkMapSlice_Set(b *testing.B) {
	Set(NewMapSlice(b.N), b)
}
func BenchmarkMapByteBuf_Set(b *testing.B) {
	Set(NewMapByteBuf(b.N), b)
}

// Get
func BenchmarkMapPointers_Get(b *testing.B) {
	Get(NewMapPointers(b.N), b)
}
func BenchmarkMapRewrite_Get(b *testing.B) {
	Get(NewMapRewrite(b.N), b)
}
func BenchmarkMapSlice_Get(b *testing.B) {
	Get(NewMapSlice(b.N), b)
}
func BenchmarkMapByteBuf_Get(b *testing.B) {
	Get(NewMapByteBuf(b.N), b)
}

// Update
func BenchmarkMapPointers_Update(b *testing.B) {
	Update(NewMapPointers(b.N), b)
}
func BenchmarkMapRewrite_Update(b *testing.B) {
	Update(NewMapRewrite(b.N), b)
}
func BenchmarkMapSlice_Update(b *testing.B) {
	Update(NewMapSlice(b.N), b)
}
func BenchmarkMapByteBuf_Update(b *testing.B) {
	Update(NewMapByteBuf(b.N), b)
}

// Delete
func BenchmarkMapPointers_Delete(b *testing.B) {
	Delete(NewMapPointers(b.N), b)
}
func BenchmarkMapRewrite_Delete(b *testing.B) {
	Delete(NewMapRewrite(b.N), b)
}
func BenchmarkMapSlice_Delete(b *testing.B) {
	Delete(NewMapSlice(b.N), b)
}
func BenchmarkMapByteBuf_Delete(b *testing.B) {
	Delete(NewMapByteBuf(b.N), b)
}

// SetGet
func BenchmarkMapPointers_SetGet(b *testing.B) {
	SetGet(NewMapPointers(b.N), b)
}
func BenchmarkMapRewrite_SetGet(b *testing.B) {
	SetGet(NewMapRewrite(b.N), b)
}
func BenchmarkMapSlice_SetGet(b *testing.B) {
	SetGet(NewMapSlice(b.N), b)
}
func BenchmarkMapByteBuf_SetGet(b *testing.B) {
	SetGet(NewMapByteBuf(b.N), b)
}

// SetDelete
func BenchmarkMapPointers_SetDelete(b *testing.B) {
	SetDelete(NewMapPointers(b.N), b)
}
func BenchmarkMapRewrite_SetDelete(b *testing.B) {
	SetDelete(NewMapRewrite(b.N), b)
}
func BenchmarkMapSlice_SetDelete(b *testing.B) {
	SetDelete(NewMapSlice(b.N), b)
}
func BenchmarkMapByteBuf_SetDelete(b *testing.B) {
	SetDelete(NewMapByteBuf(b.N), b)
}

// GetDelete
func BenchmarkMapPointers_GetDelete(b *testing.B) {
	GetDelete(NewMapPointers(b.N), b)
}
func BenchmarkMapRewrite_GetDelete(b *testing.B) {
	GetDelete(NewMapRewrite(b.N), b)
}
func BenchmarkMapSlice_GetDelete(b *testing.B) {
	GetDelete(NewMapSlice(b.N), b)
}
func BenchmarkMapByteBuf_GetDelete(b *testing.B) {
	GetDelete(NewMapByteBuf(b.N), b)
}

// print GC times
func BenchmarkPrintGCTimes(b *testing.B) {
	keys := make([]string, 0, len(GCTimes))
	for k := range GCTimes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		log.Printf("With %s GC took %.1f ms\n", k, float64(GCTimes[k].Nanoseconds())/1000000.0)
	}
}
