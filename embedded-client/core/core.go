package core

import (
	"net/rpc"

	"github.com/gibmir/ion-go/embedded-client/pool"
)

type EmbeddedRequest0[R any] struct {
	client        *rpc.Client
	procedureName string
	rpcCallPool   *pool.RpcCallPool
}

func (r *EmbeddedRequest0[R]) Call(id string) (chan *R, chan error) {
	responseChannel := make(chan *R)
	errorChannel := make(chan error)
	go r.send(responseChannel, errorChannel)
	return responseChannel, errorChannel

}

func (r *EmbeddedRequest0[R]) send(responseChannel chan *R, errorChannel chan error) {
	var response R
	done := r.rpcCallPool.Get()
	defer r.rpcCallPool.Put(done)
	r.client.Go(r.procedureName, nil, response, done)
	result := <-done
	if resultError := result.Error; resultError != nil {
		errorChannel <- resultError
		return
	}
	responseChannel <- result.Reply.(*R)
}
