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

func TestGenerateProcedureDescriber(t *testing.T) {
	type testCase struct {
		name               string
		procedure          *schema.Procedure
		expectedType       string
		expectedErrContent string
	}

	cases := []testCase{
		{
			name: "check procedure generation",
			procedure: &schema.Procedure{
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
			},
			expectedType: `
TestProcedureNameDescriber = api.Describer1[*TestArgumentType, *TestReturnTypeTypeName]{
  Describer: &api.Describer{
    Description: &api.ProcedureDescription{
      ProcedureName: "testProcedureName",
      ArgNames: []string{
        "testArgumentName",
      },
    },
  },
}
`,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			a := assert.New(t)
			var builder strings.Builder
			err := generateProcedure(&builder, c.procedure)
			if c.expectedErrContent != "" {
				a.Contains(err.Error(), c.expectedErrContent)
			} else {
				a.Equal(c.expectedType, builder.String())
			}
		})
	}
}

func TestGenerateType(t *testing.T) {
	type testCase struct {
		name               string
		typeDeclaration    *schema.TypeDeclaration
		expectedType       string
		expectedErrContent string
	}

	cases := []testCase{

		{
			name: "check simple type",
			typeDeclaration: &schema.TypeDeclaration{
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
			},
			expectedType: `
// TestTypeName test type description
type TestTypeName struct{
  
  // TestPropertyTypeName test property type description
  TestPropertyTypeName *TestPropertyTypeTypeName
  
}
`,
		},

		{
			name: "check generic type",
			typeDeclaration: &schema.TypeDeclaration{
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
				TypeParameters: []schema.TypeParameter{
					{
						SchemaElement: &schema.SchemaElement{
							Name: "T",
						},
					},
				},
			},
			expectedType: `
// TestTypeName test type description
type TestTypeName[T any] struct{
  
  // TestPropertyTypeName test property type description
  TestPropertyTypeName *TestPropertyTypeTypeName
  
}
`,
		},

		{
			name: "check generic type(multiple generic)",
			typeDeclaration: &schema.TypeDeclaration{
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
				TypeParameters: []schema.TypeParameter{
					{
						SchemaElement: &schema.SchemaElement{
							Name: "T1",
						},
					},
					{
						SchemaElement: &schema.SchemaElement{
							Name: "T2",
						},
					},
				},
			},
			expectedType: `
// TestTypeName test type description
type TestTypeName[T1 any,T2 any] struct{
  
  // TestPropertyTypeName test property type description
  TestPropertyTypeName *TestPropertyTypeTypeName
  
}
`,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			a := assert.New(t)
			var builder strings.Builder
			err := generateType(&builder, c.typeDeclaration)
			if c.expectedErrContent != "" {
				a.Contains(err.Error(), c.expectedErrContent)
			} else {
				a.Equal(c.expectedType, builder.String())
			}
		})
	}
}
