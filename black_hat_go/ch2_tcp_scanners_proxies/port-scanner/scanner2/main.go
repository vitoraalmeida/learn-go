package main

import (
	"fmt"
	"net"
)

func main() {
	//max ports == 65535
	const ports = 1024
	for i := 79; i <= ports; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)

		conn, err := net.Dial("tcp", address)

		if err != nil {
			fmt.Printf("%s closed\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s open\n", address)
	}
}
