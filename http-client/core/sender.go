package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type HttpSender struct {
	url        string
	httpClient *http.Client
}

func (httpSender *HttpSender) sendRequest(requestBytes []byte, id, methodName string) ([]byte, error) {
	body := bytes.NewReader(requestBytes)

	// send request
	logrus.Infof("send request with id [%s] ", id)
	req, err := http.NewRequest(http.MethodPost, httpSender.url, body)

	httpResponse, err := httpSender.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client return error while trying to send request with id [%s]. %w",
			id, err)
	}
	//read response
	response, err := ioutil.ReadAll(httpResponse.Body)
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to read response body for request with id [%s]. %w ",
			id, err)
	}
	return response, nil
}

func (httpSender *HttpSender) sendNotification(notificationBytes []byte, methodName string) {
	body := bytes.NewReader(notificationBytes)

	// send notification
	logrus.Infof("send notification for method [%s] ", methodName)

	_, err := httpSender.httpClient.Post(httpSender.url, applicationJsonContentType, body)
	if err != nil {
		logrus.Errorf("http client return error while trying to send notification for method [%s]. %v", methodName, err)
	}
}
