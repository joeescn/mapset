package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"

	"github.com/joeescn/mapset"
)

func main() {
	runtime.GOMAXPROCS(2)
	ints := rand.Perm(1000)

	s := mapset.NewSet[int]()
	lock := sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := 0; i < len(ints); i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			s.Add(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for _, i := range ints {
		if !s.Contains(i) {
			log.Fatalf("Set is missing element: %v", i)
		}
	}
}
