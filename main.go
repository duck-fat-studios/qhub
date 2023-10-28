package main

import (
	"fmt"
	"net"
	"sync"

	 "./devices"

	d "github.com/duck-fat-studios/dfs-go-networking/udp"
)

var (
	wg sync.WaitGroup
)

type UDPHandler struct{}

func (h *UDPHandler) Handle(data []byte, addr *net.UDPAddr) {
	Logger.Debug(fmt.Sprintf("%s : %v \n", addr, data))

}

func main() {

	setupLogger()

	Logger.Debug("Started, bitches!", true)

	handler := &UDPHandler{}
	localIP := net.ParseIP("0.0.0.0")

	genDev := 


	genDev.Init()

	server, err := d.NewUDPComms(localIP, handler, 1234)  // removed unnecessary parameters
	if err != nil {
		panic(err)
	}	

	go func() {
	wg.Add(1)
    if err := server.Run(); err != nil {
        panic(err)
    }
}()

	weiglAddress, err := net.ResolveUDPAddr("udp", "10.0.3.99:5555")

	if err !=nil {
		fmt.Println(err)
	}

	server.Send( []byte("!gva1#"), weiglAddress)

	wg.Wait()
}
