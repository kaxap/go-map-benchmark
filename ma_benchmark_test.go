package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"testing"
	"time"
)

type Stats struct {
	GCTotalTimeBefore time.Duration
	GCPauseTimeBefore time.Duration
	GCTotalTimeAfter  time.Duration
	GCPauseTimeAfter  time.Duration
	MemBefore         uint64
	MemAfter          uint64
}

var StatsMap = make(map[string]Stats)
var printed = false
var prevGCPause = time.Duration(0)

func CollectGC(t *Stats, after bool) {
	start := time.Now()
	runtime.GC()
	total := time.Since(start)
	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	if after {
		t.GCTotalTimeAfter = total
		t.GCPauseTimeAfter = stats.PauseTotal - prevGCPause
		t.MemAfter = memStats.Alloc
	} else {
		t.GCTotalTimeBefore = total
		t.GCPauseTimeBefore = stats.PauseTotal - prevGCPause
		t.MemBefore = memStats.Alloc
	}
	prevGCPause = stats.PauseTotal
}

func Set(m Map, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}
	b.StopTimer()
	key := fmt.Sprintf("A Set %T", m)
	gcTimes := Stats{}
	CollectGC(&gcTimes, false)
	runtime.KeepAlive(m)
	m = nil
	CollectGC(&gcTimes, true)
	StatsMap[key] = gcTimes
}

func Get(m Map, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Get(int32(i))
	}

	b.StopTimer()
	key := fmt.Sprintf("B Get %T", m)
	gcTimes := Stats{}
	CollectGC(&gcTimes, false)
	runtime.KeepAlive(m)
	m = nil
	CollectGC(&gcTimes, true)
	StatsMap[key] = gcTimes
}

func Update(m Map, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Update(int32(i), i, i)
	}

	b.StopTimer()
	key := fmt.Sprintf("C Update %T", m)
	gcTimes := Stats{}
	CollectGC(&gcTimes, false)
	runtime.KeepAlive(m)
	m = nil
	CollectGC(&gcTimes, true)
	StatsMap[key] = gcTimes
}

func Delete(m Map, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(int32(i))
	}

	b.StopTimer()
	key := fmt.Sprintf("D Delete %T", m)
	gcTimes := Stats{}
	CollectGC(&gcTimes, false)
	runtime.KeepAlive(m)
	m = nil
	CollectGC(&gcTimes, true)
	StatsMap[key] = gcTimes
}

func SetGet(m Map, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
		_ = m.Get(int32(i))
	}

	b.StopTimer()
	key := fmt.Sprintf("E SetGet %T", m)
	gcTimes := Stats{}
	CollectGC(&gcTimes, false)
	runtime.KeepAlive(m)
	m = nil
	CollectGC(&gcTimes, true)
	StatsMap[key] = gcTimes
}

func SetDelete(m Map, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
		m.Delete(int32(i))
	}

	b.StopTimer()
	key := fmt.Sprintf("F SetDelete %T", m)
	gcTimes := Stats{}
	CollectGC(&gcTimes, false)
	runtime.KeepAlive(m)
	m = nil
	CollectGC(&gcTimes, true)
	StatsMap[key] = gcTimes
}

func GetDelete(m Map, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{a: i, b: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Get(int32(i))
		m.Delete(int32(i))
	}

	b.StopTimer()
	key := fmt.Sprintf("G GetDelete %T", m)
	gcTimes := Stats{}
	CollectGC(&gcTimes, false)
	runtime.KeepAlive(m)
	m = nil
	CollectGC(&gcTimes, true)
	StatsMap[key] = gcTimes
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
func BenchmarkMapSliceSlice_Set(b *testing.B) {
	Set(NewMapSliceSlice(b.N), b)
}
func BenchmarkMapByteBuf_Set(b *testing.B) {
	Set(NewMapByteBuf(b.N), b)
}
func BenchmarkMapSliceStruct_Set(b *testing.B) {
	Set(NewMapSliceStruct(b.N), b)
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
func BenchmarkMapSliceSlice_Get(b *testing.B) {
	Get(NewMapSliceSlice(b.N), b)
}
func BenchmarkMapByteBuf_Get(b *testing.B) {
	Get(NewMapByteBuf(b.N), b)
}
func BenchmarkMapSliceStruct_Get(b *testing.B) {
	Get(NewMapSliceStruct(b.N), b)
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
func BenchmarkMapSliceSlice_Update(b *testing.B) {
	Update(NewMapSliceSlice(b.N), b)
}
func BenchmarkMapByteBuf_Update(b *testing.B) {
	Update(NewMapByteBuf(b.N), b)
}
func BenchmarkMapSliceStruct_Update(b *testing.B) {
	Update(NewMapSliceStruct(b.N), b)
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
func BenchmarkMapSliceSlice_Delete(b *testing.B) {
	Delete(NewMapSliceSlice(b.N), b)
}
func BenchmarkMapByteBuf_Delete(b *testing.B) {
	Delete(NewMapByteBuf(b.N), b)
}
func BenchmarkMapSliceStruct_Delete(b *testing.B) {
	Delete(NewMapSliceStruct(b.N), b)
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
func BenchmarkMapSliceSlice_SetGet(b *testing.B) {
	SetGet(NewMapSliceSlice(b.N), b)
}
func BenchmarkMapByteBuf_SetGet(b *testing.B) {
	SetGet(NewMapByteBuf(b.N), b)
}
func BenchmarkNewMapSliceStruct(b *testing.B) {
	SetGet(NewMapSliceStruct(b.N), b)
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
func BenchmarkMapSliceSlice_SetDelete(b *testing.B) {
	SetDelete(NewMapSliceSlice(b.N), b)
}
func BenchmarkMapByteBuf_SetDelete(b *testing.B) {
	SetDelete(NewMapByteBuf(b.N), b)
}
func BenchmarkMapSliceStruct_SetDelete(b *testing.B) {
	SetDelete(NewMapSliceStruct(b.N), b)
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
func BenchmarkMapSliceSlice_GetDelete(b *testing.B) {
	GetDelete(NewMapSliceSlice(b.N), b)
}
func BenchmarkMapByteBuf_GetDelete(b *testing.B) {
	GetDelete(NewMapByteBuf(b.N), b)
}
func BenchmarkMapSliceStruct_GetDelete(b *testing.B) {
	GetDelete(NewMapSliceStruct(b.N), b)
}

// print GC times
func BenchmarkPrintGCTimes(b *testing.B) {
	if printed {
		return
	}
	printed = true
	keys := make([]string, 0, len(StatsMap))
	for k := range StatsMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Println("|Name|Total ms before sweep|Pause ms before sweep|Total ms after sweep|Pause after ms sweep|Mem GB before sweep|Mem GB after sweep|")
	fmt.Println("|----|---------------------|---------------------|--------------------|--------------------|")
	for _, k := range keys {

		fmt.Printf("|%s|%.1f|%.3f|%1.f|%.3f|%.1f|%.1f|\n",
			strings.Replace(k, "*main.", "", 1),
			float64(StatsMap[k].GCTotalTimeBefore.Nanoseconds())/1000000.0,
			float64(StatsMap[k].GCPauseTimeBefore.Nanoseconds())/1000000.0,
			float64(StatsMap[k].GCTotalTimeAfter.Nanoseconds())/1000000.0,
			float64(StatsMap[k].GCPauseTimeAfter.Nanoseconds())/1000000.0,
			float64(StatsMap[k].MemBefore)/(1024*1024*1024),
			float64(StatsMap[k].MemAfter)/(1024*1024*1024),
		)
	}
}
