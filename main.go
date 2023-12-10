package main

import (
	"fmt"
	"sync"
	"time"
	// "unsafe"
)

const (
	CACHELINE_SIZE = 64
	MAX_LOOP       = 100000000
)

type noCache struct {
	x int
	y int
}

type withCache struct {
	x   int
	arr []int
	y   int
}

func main() {
	var wg sync.WaitGroup
	noC := noCache{0, 0}  // No cacheline

	// ========================================================================================
	// Uncomment to enable exploiting cacheline
	// var foo int
	// withC := withCache{0, make([]int, (CACHELINE_SIZE/unsafe.Sizeof(foo))-unsafe.Sizeof(foo)), 0}  // With Cacheline
	// =========================================================================================


	now := time.Now()

	wg.Add(1)
	go func(c *int) {
		defer wg.Done()
		for i := 0; i < MAX_LOOP; i++ {
			*c += i
		}
	}(&noC.x)

	wg.Add(1)
	go func(c *int) {
		defer wg.Done()
		for i := 0; i < MAX_LOOP; i++ {
			*c += i
		}
	}(&noC.y)

	wg.Wait()

	fmt.Printf("It took: %v\n", time.Now().Sub(now))
}
