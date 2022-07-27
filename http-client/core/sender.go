package core

import (
	"fmt"
	"github.com/gibmir/ion-go/pool"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HttpSender struct {
	url        string
	httpClient *http.Client
	bufferPool *pool.BufferPool
}

func (httpSender *HttpSender) sendRequest(requestBytes []byte, id, methodName string) ([]byte, error) {
	buffer := httpSender.bufferPool.Get()
	defer httpSender.bufferPool.Put(buffer)
	_, err := buffer.Write(requestBytes)
	if err != nil {
		return nil, fmt.Errorf("unable to write request with id [%s]. %w", id, err)

	}

	// send request
	logrus.Infof("send request with id [%s] ", id)
	httpResponse, err := httpSender.httpClient.Post(httpSender.url, applicationJsonContentType, buffer)
	if err != nil {
		return nil, fmt.Errorf("http client return error while trying to send request with id [%s]. %w",
			id, err)
	}
	//read response
	buffer.Reset()
	_, err = buffer.ReadFrom(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body for request with id [%s]. %w ",
			id, err)
	}
	return buffer.Bytes(), nil
}

func (httpSender *HttpSender) sendNotification(notificationBytes []byte, methodName string) {
	buffer := httpSender.bufferPool.Get()
	defer httpSender.bufferPool.Put(buffer)

	_, err := buffer.Write(notificationBytes)
	if err != nil {
		logrus.Errorf("unable to write notification [%s]. %v", methodName, err)
	}

	// send notification
	logrus.Infof("send notification for method [%s] ", methodName)

	_, err = httpSender.httpClient.Post(httpSender.url, applicationJsonContentType, buffer)
	if err != nil {
		logrus.Errorf("http client return error while trying to send notification for method [%s]. %v", methodName, err)
	}
}
