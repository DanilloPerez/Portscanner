package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(1024)
	for i := 1; i <= 1024; i++ {
		go PortScan(i, waitgroup)
	}
	waitgroup.Wait()
}
func PortScan(j int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	address := fmt.Sprintf("scanme.nmap.org:%d", j)

	_, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("error occured:%d\n", j)
		return
	}
	fmt.Printf("Port Open: %d\n", j)
}
