package main

import (
	"fmt"

	"github.com/gibmir/ion-go/api/pkg/errors"
	"github.com/gibmir/ion-go/client/pkg/configuration"
	"github.com/gibmir/ion-go/client/pkg/http"
	"github.com/gibmir/ion-go/demo/api/testingNamespace"
	"github.com/gibmir/ion-go/processor"
	"github.com/sirupsen/logrus"
)

func main() {

	proc := processor.NewAsyncProcessor(logrus.WithField("id", "client-task-processor").Logger, 10, 1)
	proc.Start()
	c:=configuration.Configuration{
		Url:"http://localhost:55555/",
	}
	factory := http.NewHttpRequestFactory(proc, &c)
	request, err := http.NewRequest1(factory, &testingNamespace.TestProcedureDescriber)
	if err != nil {
		logrus.Fatal(err)
	}
	responses := make(chan string)
	errors := make(chan *errors.JsonRpcError)
	request.PositionalCall("test-id", &testingNamespace.TestType{TestTypeNumericProperty:10, TestListType: []float64{1,2,3}}, responses, errors)
	select {
	case response := <-responses:
		fmt.Println(response)
	case err := <-errors:
		fmt.Println(err)
	}
}
