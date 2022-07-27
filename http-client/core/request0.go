package core

import (
	"fmt"
	"github.com/gibmir/ion-go/api/dto"

	"encoding/json"
	"github.com/sirupsen/logrus"
)

// zero arg request
type HttpRequest0[R any] struct {
	*HttpRequest
}

func (r *HttpRequest0[R]) Call(id string) (chan *R, chan error) {
	responseChannel := make(chan *R)
	errorChannel := make(chan error)
	go func() {
		defer close(responseChannel)
		defer close(errorChannel)

		//prepare request
		request := dto.PositionalRequest{
			Request: &dto.Request{
				Id:       id,
				Method:   r.methodName,
				Protocol: dto.DefaultJsonRpcProtocolVersion,
			},
		}

		requestBytes, err := json.Marshal(request)
		if err != nil {
			errorChannel <- fmt.Errorf("unable to marshal request with id [%s]. %w", id, err)
			return
		}
		responseBytes, err := r.httpSender.sendRequest(requestBytes, id, r.methodName)
		if err != nil {
			errorChannel <- fmt.Errorf("unable to send request with id [%s]. %w", id, err)
			return
		}
		//unmarshall response
		var response dto.Response[R]
		err = json.Unmarshal(responseBytes, &response)
		if err != nil {
			errorChannel <- fmt.Errorf("unable to unmarshal response body for request with id [%s]. %w", id, err)
			return
		}
		logrus.Infof("response for request with id [%s] was received", id)
		if responseError := response.Error; responseError != nil {
			// user error api
			errorChannel <- fmt.Errorf("received api error as response for request with id [%s].", id)

		}
		responseChannel <- &response.Result
	}()
	return responseChannel, errorChannel
}

func (r *HttpRequest0[R]) Notify() {
	go func() {
		//prepare notification
		request := dto.PositionalRequest{
			Request: &dto.Request{
				Method:   r.methodName,
				Protocol: dto.DefaultJsonRpcProtocolVersion,
			},
		}

		notificationBytes, err := json.Marshal(request)
		if err != nil {
			logrus.Errorf("unable to marshal notification for method [%s]. %v",
				r.methodName, err)
				return
		}
		r.httpSender.sendNotification(notificationBytes, r.methodName)
	}()
}
