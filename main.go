package main

import (
	"fmt"
	"net"
	"time"
)

var (
	Peers        = make(map[string]net.Conn)
	BufferRead   = make([]byte, 1024)
	LocalhostIp4 = "localhost:0"
	MulticastIp4 = "224.0.0.1:0"
)

func main() {
	for {
		//Resolve multicast addr
		multicastAddr, err := ResolveMulticast(MulticastIp4)
		ThenTerminate(LogError(err))
		//Create multicast socket
		socketMulticast, err := net.DialUDP("udp", nil, multicastAddr)
		ThenTerminate(LogError(err))

		for {
			_, err := socketMulticast.Write([]byte("Ping"))
			ThenTerminate(LogError(err))
			time.Sleep(1 * time.Second)
		}
	}

	for {
		//Resolve multicast addr
		multicastAddr, err := ResolveMulticast(MulticastIp4)
		ThenTerminate(LogError(err))
		//Create multicast socket
		socketMulticast, err := net.ListenMulticastUDP("udp", nil, multicastAddr)
		ThenTerminate(LogError(err))
		err = socketMulticast.SetReadBuffer(1024)
		ThenTerminate(LogError(err))

		for {
			_, src, err := socketMulticast.ReadFromUDP(BufferRead)
			ThenTerminate(LogError(err))
			fmt.Println(src.IP.String())
			//Implement: check if peer is commected already1
			//trim := bytes.Trim(BufferRead, "\x00")
			//peerAddress := src.IP.String() + string(trim[5+64:])

		}
	}

	//Create listening socket
	// socketListener, err := net.Listen("tcp4", LocalhostIp4)
	// ThenTerminate(LogError(err))

	// defer socketMulticast.Close()
	// defer socketListener.Close()
}
