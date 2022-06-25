package core

import (
	"encoding/json"
	"io/ioutil"
	"net"

	api "github.com/gibmir/ion-go/api/core"
	"github.com/gibmir/ion-go/api/dto"
	"github.com/gibmir/ion-go/client/cache"
	client "github.com/gibmir/ion-go/client/core"
	"github.com/gibmir/ion-go/tcp-client/pool"
	"github.com/sirupsen/logrus"
)

type IonTcpClient struct {
	connectionPool *pool.ConnectionPool
	callbacks      *cache.CallbacksCache
}

type TcpRequest0[R any] struct {
	procedureName  string
	address        string
	connectionPool *pool.ConnectionPool
	callbacks      *cache.CallbacksCache
}

func (r *TcpRequest0[R]) Call(id string) (chan *R, chan error) {
	response := make(chan *R)
	responseError := make(chan error)
	go func() {
		r.callbacks.Append(id, &cache.Callback{
			Response: response,
			Err:      responseError,
		})
		pool, errorChannel := r.connectionPool.Get(r.address)
		select {
		case connection := <-pool:
			r.send(id, connection, pool, responseError)
		case err := <-errorChannel:
			responseError <- err
		}
	}()
	return response, responseError
}

func (r *TcpRequest0[R]) send(id string, connection *net.Conn, connections chan *net.Conn,
	responseError chan error) {
	defer pool.Return(connection, connections)
	request := dto.PositionalRequest{
		Id:     id,
		Method: r.procedureName,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		responseError <- err
		return
	}
	// use prefix with data size
	(*connection).Write(requestBytes)

	responseBytes, err := ioutil.ReadAll(*connection)
	if err != nil {
		responseError <- err
	}

	responseBytes

	logrus.Infof("request with id [%s] was send", id)
}

func (r *TcpRequest0[R]) Notification() {
}

func NoArg[R any](tcpClient *IonTcpClient, procedure api.JsonRemoteProcedure0[R]) *client.Request0[R] {
	var request client.Request0[R]
	return &request
}
