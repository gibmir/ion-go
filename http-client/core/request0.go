package core

import (
	"fmt"

	"github.com/gibmir/ion-go/api/dto"
	"github.com/gibmir/ion-go/api/errors"

	"encoding/json"
)

// zero arg request
type HttpRequest0[R any] struct {
	*HttpRequest
}

func (r *HttpRequest0[R]) Call(id string, responseChannel chan<- R, errorChannel chan<- *errors.JsonRpcError) {
	r.proc.Process(func() {
		//prepare request
		request := dto.Positional{
			Request: &dto.Request{
				Id:       id,
				Method:   r.methodName,
				Protocol: dto.DefaultJsonRpcProtocolVersion,
			},
		}

		requestBytes, err := json.Marshal(request)
		if err != nil {
			errorChannel <- errors.NewInternalError(fmt.Sprintf("unable to marshal request with id [%s]. %v", id, err))
			return
		}
		r.log.Infof("sending positional request with id [%s]", id)
		responseBytes, err := r.httpSender.sendRequest(requestBytes, id, r.methodName)
		if err != nil {
			errorChannel <- errors.NewInternalError(fmt.Sprintf("unable to send request with id [%s]. %v", id, err))
			return
		}
		//unmarshall response
		var response dto.Response[R]
		err = json.Unmarshal(responseBytes, &response)
		if err != nil {
			errorChannel <- errors.NewInternalError(fmt.Sprintf("unable to unmarshal response body for request with id [%s]. %v", id, err))
			return
		}
		r.log.Infof("response for request with id [%s] was received", id)
		if responseError := response.Error; responseError != nil {
			// user error api
			errorChannel <- responseError

		}
		responseChannel <- response.Result
	})
}

func (r *HttpRequest0[R]) Notify() {
	r.proc.Process(func() {
		//prepare notification
		request := dto.Positional{
			Request: &dto.Request{
				Method:   r.methodName,
				Protocol: dto.DefaultJsonRpcProtocolVersion,
			},
		}

		notificationBytes, err := json.Marshal(request)
		if err != nil {
			r.log.Errorf("unable to marshal notification for method [%s]. %v",
				r.methodName, err)
			return
		}

		r.log.Info("sending notification")
		r.httpSender.sendNotification(notificationBytes, r.methodName)
	})
}
