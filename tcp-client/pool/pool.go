package pool

import (
	"net"
	"sync"
)

type ConnectionPool struct {
	poolSize int
	mutex    sync.RWMutex
	pool     map[string]chan *net.Conn
}

func NewConnectionPool(poolSize int) *ConnectionPool {
	return &ConnectionPool{
		poolSize: poolSize,
		pool:     make(map[string]chan *net.Conn),
		mutex:    sync.RWMutex{},
	}
}

func (pool *ConnectionPool) Get(address string) (chan *net.Conn, chan error) {
	defer pool.mutex.RUnlock()
	defer pool.mutex.Unlock()
	errorChannel := make(chan error)
	pool.mutex.RLock()
	connection := pool.pool[address]
	if connection == nil {
		connection, errorChannel = pool.createConnectionPool(address)
		pool.mutex.Lock()
		pool.pool[address] = connection
	}
	return connection, errorChannel
}

func Return(connection *net.Conn, pool chan *net.Conn) {
	pool <- connection
}

func (pool *ConnectionPool) createConnectionPool(address string) (chan *net.Conn, chan error) {
	connectionPool := make(chan *net.Conn, pool.poolSize)
	errorChannel := make(chan error)
	for i := 0; i < pool.poolSize; i++ {
		connection, err := net.Dial("tcp", address)
		if err != nil {
			errorChannel <- err
		}
		connectionPool <- &connection

	}
	return connectionPool, errorChannel
}
