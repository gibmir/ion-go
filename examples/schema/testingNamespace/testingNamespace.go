// testingNamespace default description
package testingNamespace

import api "github.com/gibmir/ion-go/api/core"

var (
	testProcedureDescriber = api.Describer1[TestType, string]{
		FirstArgument: &api.Type[TestType]{},
		Describer0: &api.Describer0[string]{
			ReturnType: &api.Type[string]{},
			Describer: &api.Describer{
				Description: &api.ProcedureDescription{
					ProcedureName: "testProcedure",
					ArgNames: []string{
						"testComposedArgument",
					},
				},
			},
		},
	}
)

//TestType test type
type TestType struct {

	//testTypeNumericProperty numeric property
	testTypeNumericProperty float64
}
