package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	host := "scanme.nmap.org"
	//max ports == 65535
	lastPort := 80

	for i := 70; i <= lastPort; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", host, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}

	wg.Wait()
}
