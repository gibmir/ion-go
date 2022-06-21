package core

import (
	"encoding/json"
	"io/ioutil"
	"net"

	api "github.com/gibmir/ion-go/ion-api/core"
	"github.com/gibmir/ion-go/ion-api/dto"
	"github.com/gibmir/ion-go/ion-client/cache"
	client "github.com/gibmir/ion-go/ion-client/core"
	"github.com/sirupsen/logrus"
)

type IonTcpClient struct {
}

type ResponseReader struct {
	channels   *cache.CallbacksCache
	connection net.Conn
}

func (r *ResponseReader) Run() {
	responseBytes, err := ioutil.ReadAll(r.connection)
	if err != nil {
		logrus.Warnf("unable to read from connection [%v]", r.connection.LocalAddr())
	}
	responseMap := make(map[string]interface{})
	err = json.Unmarshal(responseBytes, &responseMap)
	if err != nil {
		logrus.Error(err)
	}
}

type TcpRequest0[R any] struct {
	procedureName  string
	address        string
	connectionPool *ConnectionPool
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

func (r *TcpRequest0[R]) send(id string, connection *net.Conn, pool chan *net.Conn,
	responseError chan error) {
	defer Return(connection, pool)
	request := dto.PositionalRequest{
		Id:     id,
		Method: r.procedureName,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		responseError <- err
	}
	// use prefix with data size
	(*connection).Write(requestBytes)
}

func (r *TcpRequest0[R]) Notification() {
}

func NoArg[R any](tcpClient *IonTcpClient, procedure api.JsonRemoteProcedure0[R]) *client.Request0[R] {
	var request client.Request0[R]
	return &request
}
