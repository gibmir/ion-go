package core

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gibmir/ion-go/pool"
	"github.com/gibmir/ion-go/processor"
	"github.com/sirupsen/logrus"
)

func TestRequest(t *testing.T) {
	bufferPool := pool.NewBufferPool(2, 100)
	proc := processor.NewAsyncProcessor(logrus.New(), 1, 1)
	proc.Start()
	request := HttpRequest1[string, string]{
		HttpRequest: &HttpRequest{
			methodName: "namedProcedure",
			proc:       proc,
			httpSender: &HttpSender{
				httpClient: http.DefaultClient,
				bufferPool: bufferPool,
				url:        "http://localhost:52222",
			},
		},
	}

	responseChannel := make(chan *string)
	errorChannel := make(chan error)
	defer close(errorChannel)
	defer close(responseChannel)

	request.PositionalCall("test-id", "pepega", responseChannel, errorChannel)

	select {
	case response := <-responseChannel:
		fmt.Println(*response)
	case err := <-errorChannel:
		fmt.Println(err)
	}
}
