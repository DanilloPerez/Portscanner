package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	address, err := ValidateUrl(os.Args[1])
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(1024)
	for i := 1; i <= 1024; i++ {
		go PortScan(i, waitgroup, address+":"+strconv.Itoa(i))
	}
	waitgroup.Wait()
}
func PortScan(port int, waitgroup *sync.WaitGroup, address string) {
	defer waitgroup.Done()
	_, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("error occured:%d\n", port)
		return
	}
	fmt.Printf("Port Open: %d\n", port)
}
func ValidateUrl(address string) (returnAddress string, returnerr error) {
	_, err := url.Parse(address)
	if err != nil {
		return "", fmt.Errorf("invalid URL")
	}
	return address, nil
}
