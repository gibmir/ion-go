// testingNamespace default description
package testingNamespace
import api "github.com/gibmir/ion-go/api/pkg/describer"
var (

TestProcedureDescriber = api.Describer1[*TestType, string]{
  Describer: &api.Describer{
    Description: &api.ProcedureDescription{
      ProcedureName: "testProcedure",
      ArgNames: []string{
        "testComposedArgument",
      },
    },
  },
}

)
// TestType test type
type TestType struct{
  
  // TestTypeNumericProperty numeric property
  TestTypeNumericProperty float64
  
  // TestTypeErrorFlag returns error if true
  TestTypeErrorFlag bool
  
  // TestTypeJsonRpcErrorFlag returns json-rpc error if true
  TestTypeJsonRpcErrorFlag bool
  
}
