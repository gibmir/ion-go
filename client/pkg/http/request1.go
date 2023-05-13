package http

import (
	"encoding/json"
	"fmt"

	"github.com/gibmir/ion-go/api/pkg/dto"
	"github.com/gibmir/ion-go/api/pkg/errors"
)

//one arg requets
type HttpRequest1[T, R any] struct {
	ArgumentName string
	*HttpRequest
}

func (r *HttpRequest1[T, R]) PositionalCall(id string, argument T, responseChannel chan<- R, errorChannel chan<- *errors.JsonRpcError) {
	r.proc.Process(func() {
		defer close(responseChannel)
		defer close(errorChannel)

		request := dto.Positional{
			Parameters: []interface{}{argument},
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
			errorChannel <- errors.NewInternalError( fmt.Sprintf("unable to send request with id [%s]. %v", id, err))
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

func (r *HttpRequest1[T, R]) PositionalNotification(argument T) {
	r.proc.Process(func() {
		//prepare notification
		request := dto.Positional{
			Parameters: []interface{}{argument},
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
		r.log.Info("sending positional notification")
		r.httpSender.sendNotification(notificationBytes, r.methodName)
	})
}
