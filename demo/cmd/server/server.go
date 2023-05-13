package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gibmir/ion-go/demo/api/testingNamespace"
	server "github.com/gibmir/ion-go/server/pkg/http"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	httpServer := server.NewServer(logrus.WithField("id", "ion-http-server").Logger)
	server.NewProcessor1(httpServer, &testingNamespace.TestProcedureDescriber, func(t *testingNamespace.TestType) (string, error) {
		return fmt.Sprint(t.TestTypeNumericProperty), nil
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", httpServer.Handle)

	err := http.ListenAndServe(":55555", mux)
	logrus.Fatal(err)
}
