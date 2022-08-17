package reader

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// NAMESPACE

	testNamespaceName        = "testNamespaceName"
	testNamespaceId          = "testNamespaceId"
	testNamespaceDescription = "testNamespaceDescription"

	// TYPES

	testTypeName                = "testType"
	testTypeId                  = "testTypeId"
	testTypeDescription         = "test type description"
	testTypePropertyName        = "testTypeProperty"
	testTypePropertyType        = "string"
	testTypePropertyId          = "testTypePropertyId"
	testTypePropertyDescription = "test type property description"

	// PARAMETERS

	testParameterName        = "testParameter"
	testParameterId          = "testParameterId"
	testParameterDescription = "testParameterDescription"

	// PROCEDURES

	testProcedureName                = "testProcedure"
	testProcedureId                  = "testId"
	testProcedureDescription         = "test procedure description"
	testProcedureArgumentName        = "testArgument"
	testProcedureArgumentType        = "string"
	testProcedureArgumentDescription = "test argument description"
	testReturnArgumentType           = "string"
	testReturnArgumentDescription    = "test return argument description"
)

func TestUnmarshalNamespace_Smoke(t *testing.T) {
	a := assert.New(t)
	jsonString := "{\n" +
		"    \"" + testNamespaceName + "\": {\n" +
		"    \"" + idKey + "\": \"" + testNamespaceId + "\",\n" +
		"    \"" + descriptionKey + "\": \"" + testNamespaceDescription + "\",\n" +
		"    \"types\": {\n" +
		"      \"" + testTypeName + "\": {\n" +
		"        \"description\": \"" + testTypePropertyDescription + "\",\n" +
		"        \"parametrization\": {\n" +
		"          \"" + testParameterName + "\": {\n" +
		"            \"id\": \"" + testParameterId + "\",\n" +
		"            \"description\": \"" + testParameterDescription + "\"\n" +
		"          }\n" +
		"        },\n" +
		"        \"properties\": {\n" +
		"          \"" + testTypePropertyName + "\": {\n" +
		"            \"id\": \"" + testTypePropertyId + "\",\n" +
		"            \"type\": \"" + testTypePropertyType + "\",\n" +
		"            \"description\": \"" + testTypePropertyDescription + "\"\n" +
		"          }\n" +
		"        }\n" +
		"      }\n" +
		"    },\n" +
		"    \"procedures\": {\n" +
		"      \"" + testProcedureName + "\": {\n" +
		"        \"description\": \"" + testProcedureDescription + "\",\n" +
		"        \"arguments\": {\n" +
		"          \"" + testProcedureArgumentName + "\": {\n" +
		"            \"type\": \"" + testProcedureArgumentType + "\",\n" +
		"            \"description\": \"" + testProcedureArgumentDescription + "\"\n" +
		"          }\n" +
		"        },\n" +
		"        \"return\": {\n" +
		"          \"type\": \"" + testReturnArgumentType + "\",\n" +
		"          \"description\": \"" + testReturnArgumentDescription + "\"\n" +
		"        }\n" +
		"      }\n" +
		"    }\n" +
		"  }\n" +
		"}\n"
	b := []byte(jsonString)
	var f interface{}

	err := json.Unmarshal(b, &f)

	a.Nil(err)
	a.Equal(map[string]interface{}{
		testNamespaceName: map[string]interface{}{
			idKey:          testNamespaceId,
			descriptionKey: testNamespaceDescription,
			typesKey: map[string]interface{}{
				testTypeName: map[string]interface{}{
					descriptionKey: testTypePropertyDescription,
					parametrizationKey: map[string]interface{}{
						testParameterName: map[string]interface{}{
							idKey:          testParameterId,
							descriptionKey: testParameterDescription,
						},
					},
					propertiesKey: map[string]interface{}{
						testTypePropertyName: map[string]interface{}{
							typeKey:        testTypePropertyType,
							descriptionKey: testTypePropertyDescription,
							idKey:          testTypePropertyId,
						},
					},
				},
			},
			proceduresKey: map[string]interface{}{
				testProcedureName: map[string]interface{}{
					descriptionKey: testProcedureDescription,
					argumentsKey: map[string]interface{}{
						testProcedureArgumentName: map[string]interface{}{
							typeKey:        testProcedureArgumentType,
							descriptionKey: testProcedureArgumentDescription,
						},
					},
					returnTypeKey: map[string]interface{}{
						typeKey:        testReturnArgumentType,
						descriptionKey: testReturnArgumentDescription,
					},
				},
			},
		},
	}, f)
}

func TestUnmarshalTypes_Smoke(t *testing.T) {
	a := assert.New(t)
	jsonString := "{\n" +
		"    \"types\": {\n" +
		"    \"" + testTypeName + "\": {\n" +
		"      \"description\": \"" + testTypePropertyDescription + "\",\n" +
		"      \"parametrization\": {\n" +
		"        \"" + testParameterName + "\": {\n" +
		"          \"id\": \"" + testParameterId + "\",\n" +
		"          \"description\": \"" + testParameterDescription + "\"\n" +
		"        }\n" +
		"       },\n" +
		"      \"properties\": {\n" +
		"        \"" + testTypePropertyName + "\": {\n" +
		"          \"id\": \"" + testTypePropertyId + "\",\n" +
		"          \"type\": \"" + testTypePropertyType + "\",\n" +
		"          \"description\": \"" + testTypePropertyDescription + "\"\n" +
		"        }\n" +
		"      }\n" +
		"    }\n" +
		"  }\n" +
		"}"
	b := []byte(jsonString)
	var f interface{}

	err := json.Unmarshal(b, &f)

	a.Nil(err)
	a.Equal(map[string]interface{}{
		typesKey: map[string]interface{}{
			testTypeName: map[string]interface{}{
				descriptionKey: testTypePropertyDescription,
				parametrizationKey: map[string]interface{}{
					testParameterName: map[string]interface{}{
						idKey:          testParameterId,
						descriptionKey: testParameterDescription,
					},
				},
				propertiesKey: map[string]interface{}{
					testTypePropertyName: map[string]interface{}{
						typeKey:        testTypePropertyType,
						descriptionKey: testTypePropertyDescription,
						idKey:          testTypePropertyId,
					},
				},
			},
		},
	}, f)
}

func TestUnmarshalProcedures_Smoke(t *testing.T) {
	a := assert.New(t)

	jsonString := "{\n" +
		"    \"procedures\": {\n" +
		"    \"" + testProcedureName + "\": {\n" +
		"      \"description\": \"" + testProcedureDescription + "\",\n" +
		"      \"arguments\": {\n" +
		"        \"" + testProcedureArgumentName + "\": {\n" +
		"          \"type\": \"" + testProcedureArgumentType + "\",\n" +
		"          \"description\": \"" + testProcedureArgumentDescription + "\"\n" +
		"        }\n" +
		"      },\n" +
		"      \"return\": {\n" +
		"        \"type\": \"" + testReturnArgumentType + "\",\n" +
		"        \"description\": \"" + testReturnArgumentDescription + "\"\n" +
		"      }\n" +
		"    }\n" +
		"  }\n" +
		"}"
	b := []byte(jsonString)
	var f interface{}
	err := json.Unmarshal(b, &f)
	a.Nil(err)
	a.Equal(map[string]interface{}{
		"procedures": map[string]interface{}{
			testProcedureName: map[string]interface{}{
				descriptionKey: testProcedureDescription,
				argumentsKey: map[string]interface{}{
					testProcedureArgumentName: map[string]interface{}{
						typeKey:        testProcedureArgumentType,
						descriptionKey: testProcedureArgumentDescription,
					},
				},
				returnTypeKey: map[string]interface{}{
					typeKey:        testReturnArgumentType,
					descriptionKey: testReturnArgumentDescription,
				},
			},
		},
	}, f)
	a.Nil(nil)
}

func TestReadNamespaces_Smoke(t *testing.T) {
	a := assert.New(t)
	namespacesMap := map[string]interface{}{
		testNamespaceName: map[string]interface{}{
			idKey:          testNamespaceId,
			descriptionKey: testNamespaceDescription,
			typesKey: map[string]interface{}{
				testTypeName: map[string]interface{}{
					descriptionKey: testTypePropertyDescription,
					parametrizationKey: map[string]interface{}{
						testParameterName: map[string]interface{}{
							idKey:          testParameterId,
							descriptionKey: testParameterDescription,
						},
					},
					propertiesKey: map[string]interface{}{
						testTypePropertyName: map[string]interface{}{
							typeKey:        testTypePropertyType,
							descriptionKey: testTypePropertyDescription,
							idKey:          testTypePropertyId,
						},
					},
				},
			},
			proceduresKey: map[string]interface{}{
				testProcedureName: map[string]interface{}{
					descriptionKey: testProcedureDescription,
					idKey:          testProcedureId,
					argumentsKey: map[string]interface{}{
						testProcedureArgumentName: map[string]interface{}{
							typeKey:        testProcedureArgumentType,
							descriptionKey: testProcedureArgumentDescription,
						},
					},
					returnTypeKey: map[string]interface{}{
						typeKey:        testReturnArgumentType,
						descriptionKey: testReturnArgumentDescription,
					},
				},
			},
		},
	}

	namespaces, err := readNamespaces(namespacesMap)

	a.Nil(err)
	a.NotNil(namespaces)
	a.Equal(1, len(namespaces))
	//types
	types := namespaces[0].Types
	a.NotNil(types[testTypeName])
	testType := types[testTypeName]
	a.Equal(defaultId, testType.SchemaElement.Id)
	a.Equal(testTypePropertyDescription, testType.SchemaElement.Description)
	//properties
	a.Equal(1, len(testType.PropertyTypes))
	a.NotNil(testType.PropertyTypes[0])
	a.Equal(testTypePropertyType, testType.PropertyTypes[0].TypeName)
	a.Equal(testTypePropertyId, testType.PropertyTypes[0].Id)
	a.Equal(testTypePropertyName, testType.PropertyTypes[0].Name)
	a.Equal(testTypePropertyDescription, testType.PropertyTypes[0].Description)
	//parameters
	a.Equal(1, len(testType.TypeParameters))
	a.NotNil(testType.TypeParameters[0])
	a.Equal(testParameterId, testType.TypeParameters[0].SchemaElement.Id)
	a.Equal(testParameterName, testType.TypeParameters[0].SchemaElement.Name)
	a.Equal(testParameterDescription, testType.TypeParameters[0].SchemaElement.Description)
	//procedures

	procedures := namespaces[0].Procedures
	a.Equal(1, len(procedures))
	//procedure
	a.Equal(testProcedureName, procedures[0].Name)
	a.Equal(testProcedureId, procedures[0].Id)
	a.Equal(testProcedureDescription, procedures[0].Description)
	//arguments
	a.Equal(1, len(procedures[0].ArgumentTypes))
	a.Equal(testProcedureArgumentType, procedures[0].ArgumentTypes[0].TypeName)
	a.Equal(defaultId, procedures[0].ArgumentTypes[0].Id)
	a.Equal(testProcedureArgumentDescription, procedures[0].ArgumentTypes[0].Description)
	a.Equal(testProcedureArgumentName, procedures[0].ArgumentTypes[0].Name)
	// return type
	a.Equal(testReturnArgumentType, procedures[0].ReturnType.TypeName)
	a.Equal(testReturnArgumentDescription, procedures[0].ReturnType.Description)
	a.Equal(returnTypeName, procedures[0].ReturnType.Name)
	a.Equal(defaultId, procedures[0].ReturnType.Id)
}

func TestReadTypes_Smoke(t *testing.T) {
	a := assert.New(t)
	typesMap := map[string]interface{}{
		testTypeName: map[string]interface{}{
			descriptionKey: testTypePropertyDescription,
			parametrizationKey: map[string]interface{}{
				testParameterName: map[string]interface{}{
					idKey:          testParameterId,
					descriptionKey: testParameterDescription,
				},
			},
			propertiesKey: map[string]interface{}{
				testTypePropertyName: map[string]interface{}{
					typeKey:        testTypePropertyType,
					descriptionKey: testTypePropertyDescription,
					idKey:          testTypePropertyId,
				},
			},
		},
	}

	types, err := readTypes(typesMap)

	a.Nil(err)
	a.NotNil(types[testTypeName])
	testType := types[testTypeName]
	a.Equal(defaultId, testType.SchemaElement.Id)
	a.Equal(testTypePropertyDescription, testType.SchemaElement.Description)
	//properties
	a.Equal(1, len(testType.PropertyTypes))
	a.NotNil(testType.PropertyTypes[0])
	a.Equal(testTypePropertyType, testType.PropertyTypes[0].TypeName)
	a.Equal(testTypePropertyId, testType.PropertyTypes[0].Id)
	a.Equal(testTypePropertyName, testType.PropertyTypes[0].Name)
	a.Equal(testTypePropertyDescription, testType.PropertyTypes[0].Description)
	//parameters
	a.Equal(1, len(testType.TypeParameters))
	a.NotNil(testType.TypeParameters[0])
	a.Equal(testParameterId, testType.TypeParameters[0].SchemaElement.Id)
	a.Equal(testParameterName, testType.TypeParameters[0].SchemaElement.Name)
	a.Equal(testParameterDescription, testType.TypeParameters[0].SchemaElement.Description)
}

func TestReadProcedures_Success(t *testing.T) {
	a := assert.New(t)
	proceduresMap := map[string]interface{}{
		testProcedureName: map[string]interface{}{
			descriptionKey: testProcedureDescription,
			idKey:          testProcedureId,
			argumentsKey: map[string]interface{}{
				testProcedureArgumentName: map[string]interface{}{
					typeKey:        testProcedureArgumentType,
					descriptionKey: testProcedureArgumentDescription,
				},
			},
			returnTypeKey: map[string]interface{}{
				typeKey:        testReturnArgumentType,
				descriptionKey: testReturnArgumentDescription,
			},
		},
	}

	procedures, err := readProcedures(proceduresMap)

	a.Nil(err)
	a.NotNil(procedures)
	a.Equal(1, len(procedures))
	//procedure
	a.Equal(testProcedureName, procedures[0].Name)
	a.Equal(testProcedureId, procedures[0].Id)
	a.Equal(testProcedureDescription, procedures[0].Description)
	//arguments
	a.Equal(1, len(procedures[0].ArgumentTypes))
	a.Equal(testProcedureArgumentType, procedures[0].ArgumentTypes[0].TypeName)
	a.Equal(defaultId, procedures[0].ArgumentTypes[0].Id)
	a.Equal(testProcedureArgumentDescription, procedures[0].ArgumentTypes[0].Description)
	a.Equal(testProcedureArgumentName, procedures[0].ArgumentTypes[0].Name)
	// return type
	a.Equal(testReturnArgumentType, procedures[0].ReturnType.TypeName)
	a.Equal(testReturnArgumentDescription, procedures[0].ReturnType.Description)
	a.Equal(returnTypeName, procedures[0].ReturnType.Name)
	a.Equal(defaultId, procedures[0].ReturnType.Id)
}
