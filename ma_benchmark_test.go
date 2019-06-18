package main

import "testing"

func Set(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{i, i})
	}
}

func Get(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{i, i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Get(int32(i))
	}
}

func Update(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{i, i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Update(int32(i), i, i)
	}
}

func Delete(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{i, i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(int32(i))
	}
}

func SetGet(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{i, i})
		_ = m.Get(int32(i))
	}
}

func SetDelete(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{i, i})
		m.Delete(int32(i))
	}
}

func GetDelete(m Map, b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(int32(i), &Item{i, i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Get(int32(i))
		m.Delete(int32(i))
	}
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
