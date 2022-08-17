package generator

import (
	"strings"
	"testing"

	schema "github.com/gibmir/ion-go/schema/core"
	"github.com/stretchr/testify/assert"
)

const (
	testProcedureName                  = "testProcedureName"
	testProcedureDescription           = "test procedure description"
	testProcedureArgumentName          = "testArgumentName"
	testProcedureArgumentDescription   = "test argument description"
	testProcedureArgumentTypeName      = "testArgumentType"
	testProcedureReturnTypeDescription = "test return type description"
	testProcedureReturnTypeTypeName    = "testReturnTypeTypeName"
)

func TestGenerateProcedure_OneArg_Success(t *testing.T) {
	a := assert.New(t)
	testStringBuilder := strings.Builder{}

	p := schema.Procedure{
		SchemaElement: &schema.SchemaElement{
			Name:        testProcedureName,
			Description: testProcedureDescription,
		},
		ArgumentTypes: []schema.PropertyType{
			{
				SchemaElement: &schema.SchemaElement{
					Name:        testProcedureArgumentName,
					Description: testProcedureArgumentDescription,
				},
				TypeName: testProcedureArgumentTypeName,
			},
		},
		ReturnType: &schema.PropertyType{
			SchemaElement: &schema.SchemaElement{
				Description: testProcedureReturnTypeDescription,
			},
			TypeName: testProcedureReturnTypeTypeName,
		},
	}
	generateProcedure(&testStringBuilder, &p)
	result := testStringBuilder.String()

	a.Contains(result, testProcedureName)
	a.Contains(result, testProcedureDescription)
	//argument
	a.Contains(result, testProcedureArgumentName)
	a.Contains(result, testProcedureArgumentDescription)
	a.Contains(result, testProcedureArgumentTypeName)
	//return type
	a.Contains(result, testProcedureReturnTypeDescription)
	a.Contains(result, testProcedureReturnTypeTypeName)
}
