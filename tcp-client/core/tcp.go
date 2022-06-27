package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

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

type LengthFieldPrepender struct {
	// lengthFieldLength must be either 1, 2, 3, 4, or 8
	LengthFieldLength int
}

func (p *LengthFieldPrepender) prependeLength(length int, buffer *bytes.Buffer) error {
	if buffer == nil {
		return fmt.Errorf("provided buffer is nil")
	}
	if length < 0 {
		return fmt.Errorf("buffer length is less than zero")
	}
	switch p.LengthFieldLength {
	case 1:
		if length >= 256 {
			return fmt.Errorf("length does not fit into a byte: [%d] ", length)
		}
		return buffer.WriteByte(uint8(length))
	case 2:
		if length >= 65536 {
			return fmt.Errorf("length does not fit into a short: [%d] ", length)
		}
		n, err := buffer.Write([]byte(strconv.Itoa(length)))
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("writes more than 2 bytes for length")
		}
		return nil
	case 3:
		if length >= 16777216 {
			return fmt.Errorf("length does not fit into a medium: [%d] ", length)
		}

		n, err := buffer.Write([]byte(strconv.Itoa(length)))
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("writes more than 2 bytes for length")
		}
		return nil
	case 4:
		n, err := buffer.Write([]byte(strconv.Itoa(length)))
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("writes more than 2 bytes for length")
		}
		return nil
	case 8:
		n, err := buffer.Write([]byte(strconv.Itoa(length)))
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("writes more than 2 bytes for length")
		}
		return nil
	default:
		return fmt.Errorf("lengthFieldLength has incorrect value: [%d]", p.LengthFieldLength)
	}
}

type TcpRequest0[R any] struct {
	procedureName        string
	address              string
	LengthFieldPrepender *LengthFieldPrepender
	bufferPool           *pool.BufferPool
	connectionPool       *pool.ConnectionPool
	callbacks            *cache.CallbacksCache
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

	r.writeRequest(requestBytes,connection)

	buffer := r.bufferPool.Get()
	defer r.bufferPool.Return(buffer)
	buffer.Reset()

	n, err := buffer.ReadFrom(*connection)

	if err != nil {
		responseError <- err
	}

	logrus.Infof("request with id [%s] was send", id)
}

func (r *TcpRequest0[R]) writeRequest(requestBytes []byte, connection *net.Conn) error {
	buffer := r.bufferPool.Get()
	defer r.bufferPool.Return(buffer)
	length := len(requestBytes)
	err := r.LengthFieldPrepender.prependeLength(length, buffer)
	if err != nil {
		return err
	}
	n, err := buffer.Write(requestBytes)
	if err != nil {
		return err
	}
	n64, err := buffer.WriteTo(*connection)
	if err != nil {
		return err
	}

	if n64 != int64(n+r.LengthFieldPrepender.LengthFieldLength) {
		return fmt.Errorf("Incorrect bytes count request bytes [%d], request with length prefix [%d], length field [%d]",
			n, n64, r.LengthFieldPrepender.LengthFieldLength)
	}
	return nil
}

func (r *TcpRequest0[R]) Notification() {
}

func NoArg[R any](tcpClient *IonTcpClient, procedure api.JsonRemoteProcedure0[R]) *client.Request0[R] {
	var request client.Request0[R]
	return &request
}
