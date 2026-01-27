package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu1, mu2 sync.Mutex

	wg.Add(2)
	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("goroutine 1 acquired lock 1")
		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("goroutine 1 acquired lock 2")
	}()
	go func() {
		defer wg.Done()
		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("goroutine 2 acquired lock 2")
		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("goroutine 2 acquired lock 1")
	}()
	wg.Wait()
}
