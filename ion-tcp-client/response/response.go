package response

import (
	"encoding/json"
	"io/ioutil"
	"github.com/gibmir/ion-go/ion-api/dto"
	"github.com/gibmir/ion-go/ion-client/cache"
	"github.com/gibmir/ion-go/ion-tcp-client/pool"
	"github.com/sirupsen/logrus"

)

const (
	idKey     = "id"
	resultKey = "result"
	errorKey  = "error"
)

type ResponseReader struct {
	channels   *cache.CallbacksCache
	connectionPool *pool.ConnectionPool
}

func (r *ResponseReader) Run() {

	responseBytes, err := ioutil.ReadAll(r.connectionPool)
	if err != nil {
		logrus.Warnf("unable to read from connection [%v]", r.connection.LocalAddr())
		return
	}

	var responseMap map[string]interface{}
	//transfer this code to ion-client
	err = json.Unmarshal(responseBytes, &responseMap)
	if err != nil {
		logrus.Error(err)
	}

	responseId := responseMap[idKey].(string)
	callback := r.channels.Poll(responseId)
	if callback == nil {
		logrus.Warn("there is no callback in cache for id [%s]", responseId)
		return
	}

	responseError := responseMap[errorKey]
	if responseError != nil {
		//process error
		var responseErrorDto dto.ErrorResponse
		err := json.Unmarshal(nil, &responseErrorDto)
		if err != nil {
		}
		callback.Err <- &responseErrorDto
	}
	responseResult := responseMap[resultKey]
}
