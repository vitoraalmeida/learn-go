package main

import (
	"fmt"
	"sync"
)

func main() {
	//max ports == 65535
	maxPorts := 1024
	//host := "127.0.0.1"
	ports := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 0; i <= cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i := 0; i <= maxPorts; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
}

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}
