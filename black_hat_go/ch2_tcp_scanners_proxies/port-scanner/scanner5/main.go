package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func main() {
	//max ports == 65535
	maxPorts := 65535
	host := "scanme.nmap.org"
	ports := make(chan int, 2000)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < 1000; i++ {
		go worker(ports, results, host)
	}

	go func() {
		for i := 1; i <= maxPorts; i++ {
			ports <- i
		}
	}()

	for i := 0; i < maxPorts; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}

func worker(ports, results chan int, host string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.DialTimeout("tcp", address, 1000*time.Millisecond)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		fmt.Println("open: ", p)
		results <- p
	}
}
