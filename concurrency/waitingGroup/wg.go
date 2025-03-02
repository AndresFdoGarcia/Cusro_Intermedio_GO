package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting... %d\n", i)
	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}
