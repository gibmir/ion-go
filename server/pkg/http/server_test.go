package http

import (
	api "github.com/gibmir/ion-go/api/pkg/describer"
)

func Example() {
	// init server
	server := &HttpServer{}

	// create describer(or use ionc and json schema)
	describer := &api.Describer1[string, string]{}

	// add processor for procedure to server
	NewProcessor1(server, describer, func(testString string) (string, error) {
		return testString, nil
	})

}
