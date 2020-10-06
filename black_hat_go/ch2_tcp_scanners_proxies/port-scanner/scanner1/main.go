package main

import (
	"fmt"
	"net"
)

func main() {
	//       connection kind  host: "address:port"
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}
