package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	CACHELINE_SIZE = 64
	MAX_LOOP       = 100000000
)


// No false sharing
type ValNoShare struct {
	x int
	_ [CACHELINE_SIZE]byte // Prevent false sharing
	y int
}

// With false sharing
type ValShare struct {
	x int
	y int
}

func run(x *int, y *int) {
	var wg sync.WaitGroup
	now := time.Now()

	wg.Add(1)
	go func(c *int) {
		defer wg.Done()
		for i := 0; i < MAX_LOOP; i++ {
			*c += i
		}
	}(x)

	wg.Add(1)
	go func(c *int) {
		defer wg.Done()
		for i := 0; i < MAX_LOOP; i++ {
			*c += i
		}
	}(y)

	wg.Wait()

	fmt.Printf("It took: %v\n", time.Now().Sub(now))
}

func main() {
	valShare := ValShare{}
	fmt.Println("Sharing cacheline")
	run(&valShare.x, &valShare.y)

	fmt.Println()

	noshare := ValNoShare{}
	fmt.Println("No sharing cacheline")
	run(&noshare.x, &noshare.y)

}
