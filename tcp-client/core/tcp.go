package core

import (
	"encoding/json"
	"fmt"
	"net"

	api "github.com/gibmir/ion-go/api/core"
	"github.com/gibmir/ion-go/api/dto"
	client "github.com/gibmir/ion-go/client/core"
	"github.com/gibmir/ion-go/tcp-client/pool"
	"github.com/gibmir/ion-go/tcp-common/length"
	"github.com/sirupsen/logrus"
)

type IonTcpClient struct {
	connectionPool *pool.ConnectionPool
}


type TcpRequest0[R any] struct {
	procedureName      string
	address            string
	lengthFieldService *length.LengthFieldService
	bufferPool         *pool.BufferPool
	connectionPool     *pool.ConnectionPool
}

func (r *TcpRequest0[R]) Call(id string) (chan *R, chan error) {
	//todo pooling
	response := make(chan *R)
	responseError := make(chan error)
	go func() {
		pool, errorChannel := r.connectionPool.Get(r.address)
		select {
		case connection := <-pool:
			r.send(id, connection, pool, response, responseError)
		case err := <-errorChannel:
			responseError <- err
		}
	}()
	return response, responseError
}

func (r *TcpRequest0[R]) send(id string, connection *net.Conn, connections chan *net.Conn, responseChannel chan *R,
	responseError chan error) {
	defer pool.Return(connection, connections)
	request := dto.PositionalRequest{
		Id:     id,
		Method: r.procedureName,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		err = fmt.Errorf("unable to marshal request with id [%s]. %w", id, err)
		logrus.Debug(err)
		responseError <- err
		close(responseChannel)
		return
	}

	buffer := r.bufferPool.Get()
	defer r.bufferPool.Return(buffer)

	length := len(requestBytes)
	err = r.lengthFieldService.WriteLengthField(length, buffer)
	if err != nil {
		err = fmt.Errorf("unable to write length field to request with id [%s]. %w",
			id, err)
		logrus.Debug(err)
		responseError <- err
		close(responseChannel)
		return
	}
	_, err = buffer.Write(requestBytes)
	if err != nil {
		err = fmt.Errorf("unable to write request with id [%s] in buffer. %w", id, err)
		logrus.Debug(err)
		responseError <- err
		close(responseChannel)
		return
	}
	_, err = buffer.WriteTo(*connection)
	if err != nil {
		err = fmt.Errorf("unable to write request with id [%s]. %w", id, err)
		logrus.Debug(err)
		responseError <- err
		close(responseChannel)
		return
	}

	// starting to read response
	buffer.Reset()

	_, err = buffer.ReadFrom(*connection)
	if err != nil {
		err = fmt.Errorf("unable to read response with id [%s]. %w", id, err)
		logrus.Debug(err)
		responseError <- err
		close(responseChannel)
		return
	}
	r.lengthFieldService.ReadLengthField(buffer)

	responseBytes := buffer.Bytes()
	var response dto.Response[R]
	// todo batches
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		err = fmt.Errorf("unable to unmarshal response with id [%s]. %w", id, err)
		logrus.Debug(err)
		responseError <- err
		close(responseChannel)
		return
	}
	logrus.Infof("received a response for request with id [%s]", id)
	responseChannel <- &response.Result
	close(responseError)
}

func (r *TcpRequest0[R]) Notification() {
}

func NoArg[R any](tcpClient *IonTcpClient, procedure api.JsonRemoteProcedure0[R]) client.Request0[R] {
	var request TcpRequest0[R]
	return &request
}
