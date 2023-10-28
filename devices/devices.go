package qhubdevice

import (
	// d "github.com/duck-fat-studios/dfs-go-networking/udp"
	"net"

	k "github.com/ericboxer/go-kindling"
)

type Device interface {
	Init(ip string, port int, logger k.Logger)
	Passthrough(data []byte)
	Send(data []byte)

}


type GenericUDPDevice struct {
	udpAddress net.UDPAddr
	logger k.Logger
}


func (g *GenericUDPDevice) Init(ip string, port int, logger k.Logger) {
	
	g.udpAddress = net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
	
	
	g.logger = logger

}

func (g *GenericUDPDevice) Send(data []byte) {
	g.logger.Debug(string(data))
}
