package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i<= 1024; i++ {
		wg.Add(1)
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		_, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Printf("Connection Successful at port: %d\n", i)
		} else {
			fmt.Printf("There was a problem connecting to host at port: %d\n", i)
		}
	}

}
