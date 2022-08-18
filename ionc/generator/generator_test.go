package generator

import (
	"strings"
	"testing"

	schema "github.com/gibmir/ion-go/schema/core"
	"github.com/stretchr/testify/assert"
)

const (
	// procedures
	testProcedureName                  = "testProcedureName"
	testProcedureDescription           = "test procedure description"
	testProcedureArgumentName          = "testArgumentName"
	testProcedureArgumentDescription   = "test argument description"
	testProcedureArgumentTypeName      = "testArgumentType"
	testProcedureReturnTypeDescription = "test return type description"
	testProcedureReturnTypeTypeName    = "testReturnTypeTypeName"
	// types
	testTypeName                = "testTypeName"
	testTypeDescription         = "test type description"
	testPropertyTypeName        = "testPropertyTypeName"
	testPropertyTypeDescription = "test property type description"
	testPropertyTypeTypeName    = "testPropertyTypeTypeName"
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
	err := generateProcedure(&testStringBuilder, &p)

	a.Nil(err)

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

func TestGenerateType_Simple_OneProperty_Success(t *testing.T) {
	a := assert.New(t)
	testStringBuilder := strings.Builder{}

	td := schema.TypeDeclaration{
		SchemaElement: &schema.SchemaElement{
			Name:        testTypeName,
			Description: testTypeDescription,
		},
		PropertyTypes: []schema.PropertyType{
			{
				SchemaElement: &schema.SchemaElement{
					Name:        testPropertyTypeName,
					Description: testPropertyTypeDescription,
				},
				TypeName: testPropertyTypeTypeName,
			},
		},
	}

	err := generateType(&testStringBuilder, &td)
	a.Nil(err)

	result := testStringBuilder.String()

	a.Contains(result, testTypeName)
	a.Contains(result, testTypeDescription)
	//properties
	a.Contains(result, testPropertyTypeName)
	a.Contains(result, testPropertyTypeDescription)
	a.Contains(result, testPropertyTypeTypeName)
}
