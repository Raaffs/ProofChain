package rpc

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ErrNoAvailableRPC = errors.New("no healthy RPCs available")
type ClientManager struct {
	mu      sync.Mutex
	clients []*ethclient.Client
	urls    []string
	current int
}

func NewClientManager(urls []string) (*ClientManager, error) {
	var clients []*ethclient.Client
	for _, url := range urls {
		c, err := ethclient.Dial(url)
		if err != nil {
			continue // skip dead endpoints
		}
		clients = append(clients, c)
	}

	if len(clients) == 0 {
		return nil, ErrNoAvailableRPC
	}

	return &ClientManager{
		clients: clients,
		urls:    urls,
	}, nil
}

func (m *ClientManager) Current() *ethclient.Client {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.clients[m.current]
}

func (m *ClientManager) Rotate() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.current = (m.current + 1) % len(m.clients)
}