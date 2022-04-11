package main

import (
	"fmt"
	"net"
)

var (
	BufferRead = make([]byte, 1024)
)

func GetAddr(listener net.Listener) string {
	return listener.Addr().(*net.TCPAddr).IP.String()
}

func GetPort(listener net.Listener) int {
	return listener.Addr().(*net.TCPAddr).Port
}

//Resolve multicast addr
func ResolveMulticast(addr string) (*net.UDPAddr, error) {
	multicastAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	if !multicastAddr.IP.IsMulticast() {
		return nil, fmt.Errorf("the addr %s is not multicast", multicastAddr.IP.String())
	}
	return multicastAddr, nil
}

func MulticastEcho(msg string, socket *net.UDPConn) {
	_, err := socket.Write([]byte(msg))
	Ok(Log(err))
}

// func MulticastListen(socket *net.UDPConn, buffer []byte) {
// 	_, src, err := socket.ReadFromUDP(buffer)
// 	Terminate(LogError(err))
// }
