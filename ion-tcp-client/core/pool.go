package tcp

import "net"

type ConnectionPool struct {
	poolSize int
	pool     map[string]chan *net.Conn
}

func NewConnectionPool(poolSize int) *ConnectionPool {
	return &ConnectionPool{
		poolSize: poolSize,
		pool:     make(map[string]chan *net.Conn),
	}
}
func (pool *ConnectionPool) Get(addressString string) (chan *net.Conn, chan error) {
	errorChannel := make(chan error)
	connection := pool.pool[addressString]
	if connection == nil {
		connection = pool.createConnection()
		pool.pool[addressString] = connection
	}
	return connection, errorChannel
}

func (pool *ConnectionPool) createConnection() chan *net.Conn {
	return make(chan *net.Conn)
}
