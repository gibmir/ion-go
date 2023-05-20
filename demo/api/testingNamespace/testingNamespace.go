// testingNamespace default description
package testingNamespace
import . "github.com/gibmir/ion-go/api/pkg/describer"
var (

TestProcedureDescriber = Describer1[*TestType, string]{
  Describer: &Describer{
    Description: &ProcedureDescription{
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
  
  // TestTypeErrorFlag returns error if true
  TestTypeErrorFlag bool
  
  // TestTypeJsonRpcErrorFlag returns json-rpc error if true
  TestTypeJsonRpcErrorFlag bool
  
  // TestListType test array property
  TestListType List[float64]
  
  // TestTypeNumericProperty numeric property
  TestTypeNumericProperty float64
  
}
