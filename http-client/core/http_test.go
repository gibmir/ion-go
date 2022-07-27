package core

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gibmir/ion-go/pool"
)

func TestHttpRequest0_Success(t *testing.T) {
	bufferPool := pool.NewBufferPool(2, 100)
	request := HttpRequest1[string,string]{
		HttpRequest: &HttpRequest{
			methodName: "namedProcedure",
			httpSender: &HttpSender{
				httpClient: http.DefaultClient,
				bufferPool: bufferPool,
				url:"http://localhost:52222",
			},
		},
	}
	responseChannel, errorChannel := request.PositionalCall("test-id", "pepega")
	select {
	case response := <-responseChannel:
		fmt.Println(*response)
	case err := <-errorChannel:
		fmt.Println(err)
	}
}
