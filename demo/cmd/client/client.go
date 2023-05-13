package main

import (
	"fmt"

	"github.com/gibmir/ion-go/api/errors"
	"github.com/gibmir/ion-go/demo/api/testingNamespace"
	"github.com/gibmir/ion-go/http-client/configuration"
	client "github.com/gibmir/ion-go/http-client/core"
	"github.com/gibmir/ion-go/processor"
	"github.com/sirupsen/logrus"
)

func main() {

	proc := processor.NewAsyncProcessor(logrus.WithField("id", "client-task-processor").Logger, 10, 1)
	proc.Start()
	factory := client.NewHttpRequestFactory(proc, &configuration.Configuration{
		Url: "http://localhost:55555/",
	})
	request, err := client.NewRequest1(factory, &testingNamespace.TestProcedureDescriber)
	if err != nil {
		logrus.Fatal(err)
	}
	responses := make(chan string)
	errors := make(chan *errors.JsonRpcError)
	request.PositionalCall("test-id", &testingNamespace.TestType{TestTypeNumericProperty:10}, responses, errors)
	select {
	case response := <-responses:
		fmt.Println(response)
	case err := <-errors:
		fmt.Println(err)
	}
}
