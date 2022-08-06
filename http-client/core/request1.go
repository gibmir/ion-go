package core

import (
	"encoding/json"
	"fmt"
	"github.com/gibmir/ion-go/api/dto"
	"github.com/sirupsen/logrus"
)

//one arg requets
type HttpRequest1[T, R any] struct {
	*HttpRequest
}

func (r *HttpRequest1[T, R]) PositionalCall(id string, argument T, responseChannel chan *R, errorChannel chan error) {
	go func(id string, argument T, responseChannel chan *R, errorChannel chan error) {
		defer close(responseChannel)
		defer close(errorChannel)

		request := dto.PositionalRequest{
			Parameters: []interface{}{argument},
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
	}(id, argument, responseChannel, errorChannel)
}

func (r *HttpRequest1[T, R]) PositionalNotification(argument T) {
	go func(argument T) {
		//prepare notification
		request := dto.PositionalRequest{
			Parameters: []interface{}{argument},
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
	}(argument)
}
