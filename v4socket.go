package main

import (
	"fmt"
	"net"
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
		return nil, fmt.Errorf("The addr %s is not multicast", multicastAddr.IP.String())
	}
	return multicastAddr, nil
}

func MulticastEcho(msg string, socket *net.UDPConn) {
	_, err := socket.Write([]byte(msg))
	ThenTerminate(LogError(err))
}

// func MulticastListen(socket *net.UDPConn, buffer []byte) {
// 	_, src, err := socket.ReadFromUDP(buffer)
// 	ThenTerminate(LogError(err))
// }
