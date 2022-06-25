package pool

import (
	"net/rpc"
)

type RpcCallPool struct {
	poolSize int
	pool     chan chan *rpc.Call
}

func (pool *RpcCallPool) Get() chan *rpc.Call {
	return <-pool.pool
}

func (pool *RpcCallPool) Put(call chan *rpc.Call) {
	pool.pool <- call
}
