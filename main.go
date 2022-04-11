package main

const (
	LevelDB = "./leveldb"
)

func main() {

	storage := NewStorage(LevelDB)

	// for {
	// 	//Resolve multicast addr
	// 	multicastAddr, err := ResolveMulticast(MulticastIp4)
	// 	Terminate(LogError(err))
	// 	//Create multicast socket
	// 	socketMulticast, err := net.DialUDP("udp", nil, multicastAddr)
	// 	Terminate(LogError(err))
	// 	for {
	// 		_, err := socketMulticast.Write([]byte("Ping"))
	// 		Terminate(LogError(err))
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }
	// for {
	// 	//Resolve multicast addr
	// 	multicastAddr, err := ResolveMulticast(MulticastIp4)
	// 	Terminate(LogError(err))
	// 	//Create multicast socket
	// 	socketMulticast, err := net.ListenMulticastUDP("udp", nil, multicastAddr)
	// 	Terminate(LogError(err))
	// 	err = socketMulticast.SetReadBuffer(1024)
	// 	Terminate(LogError(err))
	// 	for {
	// 		_, src, err := socketMulticast.ReadFromUDP(BufferRead)
	// 		Terminate(LogError(err))
	// 		fmt.Println(src.IP.String())
	// 		//Implement: check if peer is commected already1
	// 		trim := bytes.Trim(BufferRead, "\x00")
	// 		peerAddress := src.IP.String() + string(trim[5+64:])
	// 	}
	// 	//Create listening socket
	// 	socketListener, err := net.Listen("tcp4", LocalhostIp4)
	// 	ThenTerminate(LogError(err))
	// 	defer socketMulticast.Close()
	// 	defer socketListener.Close()
	// }

	defer storage.DB.Close()
}
