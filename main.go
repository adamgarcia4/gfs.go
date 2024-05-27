package main

import (
	"fmt"
	"sync"
)

// I want to design a safe counter which will
// be used by multiple goroutines.
// 10 different goroutines will increment the counter a single time.
// The counter should be 10 at the end of the program.
type SafeCounter struct {
	mutex sync.Mutex
	data  map[string]int
}

func (counter *SafeCounter) Increment(key string, incrementNum int) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.data[key] += 1
}

func main() {
	fmt.Println("Hi")
}
