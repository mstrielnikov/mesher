package main

type Peer struct {
	Name    string
	Address string
	Alias   string
	Hash    string
}

type Mesh struct {
	Connections []Peer
}

func (m *Mesh) AddPeer(peer Peer) Mesh {
	m.Connections = append(m.Connections, peer)
	return *m
}

func (m *Mesh) AddPeers(peers []Peer) {
	for _, peer := range peers {
		m.AddPeer(peer)
	}
}

func DiscoverPeers(address, port string, enabled bool) []Peer {
	if !enabled {
		return nil
	}
	return []Peer{}
}

func ListenPeers(address, port string, enabled bool) []Peer {
	if !enabled {
		return nil
	}
	return []Peer{}
}

func NotifyPeers(address, port string, enabled bool) error {
	if !enabled {
		return nil
	}
	var err error
	return err
}

func SharePeers(address, port string, enabled bool) error {
	if !enabled {
		return nil
	}
	var err error
	return err
}
