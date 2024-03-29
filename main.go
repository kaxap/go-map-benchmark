package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
	"unsafe"
)

// from https://github.com/golang/go/issues/9477#issuecomment-68500643 by github.com/mish15

// Results of this program on my machine: (linux amd64, go 1.4):
//
// for t in 1 2 3 4; do go run maps.go $t; done
// With map[int32]*int32, GC took 2.831259536s
// With map[int32]int32, GC took 541.091193ms
// With map shards ([]map[int32]int32), GC took 325.312192ms
// With a plain slice ([]main.t), GC took 69.359µs
//a
// Higher parallelism does help, to some extent:
//
// for t in 1 2 3 4; do GOMAXPROCS=4 go run maps.go $t; done
// With map[int32]*int32, GC took 2.347938525s
// With map[int32]int32, GC took 362.263322ms
// With map shards ([]map[int32]int32), GC took 89.884292ms
// With a plain slice ([]main.t), GC took 312.583µs

func main() {
	const N = 30e6

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s [1 2 3 4]\n(number selects the test)\n", os.Args[0])
		return
	}

	switch os.Args[1] {
	case "1":
		// Big map with a pointer in the value
		m := make(map[int32]*int32)
		for i := 0; i < N; i++ {
			n := int32(i)
			m[n] = &n
		}
		runtime.GC()
		fmt.Printf("With %T, GC took %s\n", m, timeGC())
		_ = m[0] // Preserve m until here, hopefully

	case "1sa":
		// Big map with a pointer in the value
		m := make(map[int32]int, N)
		s := make([]*int32, 0, N)
		for i := 0; i < N; i++ {
			n := int32(i)
			s = append(s, &n)
			m[n] = len(s)
		}
		runtime.GC()
		fmt.Printf("With %T, GC took %s\n", m, timeGC())
		_ = m[0] // Preserve m until here, hopefully
		_ = s[0]

	case "1p":
		// Big map with a pointer in the value
		m := make(map[int32]unsafe.Pointer)
		s := make([]*int32, 0, N)
		for i := 0; i < N; i++ {
			n := int32(i)
			s = append(s, &n)
			m[n] = unsafe.Pointer(&n)
		}
		runtime.GC()
		fmt.Printf("With %T, GC took %s\n", m, timeGC())
		_ = m[0] // Preserve m until here, hopefully
		_ = s[0]

	case "1a":
		// Big map (preallocated) with a pointer in the value
		m := make(map[int32]*int32, N)
		for i := 0; i < N; i++ {
			n := int32(i)
			m[n] = &n
		}
		runtime.GC()
		fmt.Printf("With %T, GC took %s\n", m, timeGC())
		_ = m[0] // Preserve m until here, hopefully

	case "2":
		// Big map, no pointer in the value
		m := make(map[int32]int32)
		for i := 0; i < N; i++ {
			n := int32(i)
			m[n] = n
		}
		runtime.GC()
		fmt.Printf("With %T, GC took %s\n", m, timeGC())
		_ = m[0]

	case "2a":
		// Big map, no pointer in the value, map fully pre-allocated
		m := make(map[int32]int32, N)
		for i := 0; i < N; i++ {
			m[int32(i)] = int32(i)
		}
		runtime.GC()
		fmt.Printf("With %T, GC took %s\n", m, timeGC())
		_ = m[0]

	case "3":
		// Split the map into 100 shards
		shards := make([]map[int32]int32, 100)
		for i := range shards {
			shards[i] = make(map[int32]int32)
		}
		for i := 0; i < N; i++ {
			n := int32(i)
			shards[i%100][n] = n
		}
		runtime.GC()
		fmt.Printf("With map shards (%T), GC took %s\n", shards, timeGC())
		_ = shards[0][0]

	case "3a":
		// Split the map into 100 shards
		shards := make([]map[int32]int32, 100)
		for i := range shards {
			shards[i] = make(map[int32]int32, N/100)
		}
		for i := 0; i < N; i++ {
			shards[i%100][int32(i)] = int32(i)
		}
		runtime.GC()
		fmt.Printf("With map shards (%T), GC took %s\n", shards, timeGC())
		_ = shards[0][0]

	case "4":
		// A slice, just for comparison to show that
		// merely holding onto millions of int32s is fine
		// if they're in a slice.
		var s []*int32
		for i := 0; i < N; i++ {
			n := int32(i)
			s = append(s, &n)
		}
		runtime.GC()
		fmt.Printf("With a plain slice (%T), GC took %s\n", s, timeGC())
		_ = s[0]

	case "4a":
		// A slice, just for comparison to show that
		// merely holding onto millions of int32s is fine
		// if they're in a pre-allocated slice.
		s := make([]*int32, 0, N)
		for i := 0; i < N; i++ {
			n := int32(i)
			s = append(s, &n)
		}
		runtime.GC()
		fmt.Printf("With a plain slice (%T), GC took %s\n", s, timeGC())
		_ = s[0]
	}
}

func timeGC() time.Duration {
	start := time.Now()
	runtime.GC()
	return time.Since(start)
}
