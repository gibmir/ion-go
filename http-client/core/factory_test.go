package core

import (
	"fmt"
	"testing"

	"github.com/gibmir/ion-go/api/core"
	mocks "github.com/gibmir/ion-go/http-client/mocks"
)

func TestGenerate(t *testing.T) {
	hrf := HttpRequestFactoryEnvironment{}
	var d = core.Describer0[string]{
		ReturnType: &core.Type[string]{},

		Describer: &core.Describer{
			Description: &core.ProcedureDescription{
				ProcedureName: "someName",
			},
		},
	}
	_, _ = NewRequest0(&hrf, &d)

}

func Example() {
	// you should create environment from configuration
	requestFactoryEnvironment := mocks.RequestFactoryEnvironment{}

	// apiProcedureDescriber will be generated by ionc
	apiProcedureDescriber := core.Describer1[string, int]{
		FirstArgument: &core.Type[string]{},
		Describer0: &core.Describer0[int]{
			ReturnType: &core.Type[int]{},
			Describer: &core.Describer{
				Description: &core.ProcedureDescription{
					ProcedureName: "your API procedure",
					ArgNames: []string{
						"your procedure first argument name",
					},
				},
			},
		},
	}

	// You can make API calls with request that provided by factory function
	request1, _ := NewRequest1(&requestFactoryEnvironment, &apiProcedureDescriber)
	// You can use responses channel for data pipelining
	responsesChannel := make(chan *int)
	// You can use error channel to aggregate errors processing
	errorsChannel := make(chan error)
	// You can make a json-rpc call
	request1.PositionalCall("example-request-id", "example-request-argument",
		responsesChannel, errorsChannel)
	// You can build a result processing pipelines
	select {
	case response := <-responsesChannel:
		// received response
		fmt.Println(response)
	case err := <-errorsChannel:
		// received err
		fmt.Println(err)
	}
}
