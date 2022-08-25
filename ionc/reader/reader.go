package reader

import (
	"fmt"
	"unicode"

	schema "github.com/gibmir/ion-go/schema/core"
	"github.com/sirupsen/logrus"
)

const (
	typesKey           string = "types"
	descriptionKey     string = "description"
	defaultDescription string = "default description"
	idKey              string = "id"
	defaultId          string = "default id"
	proceduresKey      string = "procedures"
	returnTypeKey      string = "return"
	argumentsKey       string = "arguments"
	propertiesKey      string = "properties"
	typeKey            string = "type"
	parametrizationKey string = "parametrization"
	returnTypeName     string = "return"
)

var (
	commaRune              rune = ','
	schemaStartGenericRune rune = '<'
	startGenericRune       rune = '['
	schemaEndGenericRune   rune = '>'
	endGenericRune         rune = ']'
)

func ReadSchema(jsonPath string, apiJson interface{}) (*schema.Schema, error) {
	if apiJson == nil {
		return nil, fmt.Errorf("provided schema [%s] is nil", jsonPath)
	}
	if apiJsonMap, ok := apiJson.(map[string]interface{}); ok {
		namespaces, err := readNamespaces(apiJsonMap)
		if err != nil {
			return nil, fmt.Errorf("unable to parse schema [%s]. %w", jsonPath, err)
		}
		logrus.Infof("schema was successfully read from [%s]", jsonPath)
		return &schema.Schema{Namespaces: namespaces}, nil
	} else {
		return nil, fmt.Errorf("provided schema [%s] has incorrect format", jsonPath)
	}
}

func readNamespaces(apiJsonMap map[string]interface{}) ([]schema.Namespace, error) {
	namespaces := make([]schema.Namespace, 0, len(apiJsonMap))
	for namespaceName, definition := range apiJsonMap {
		if namespaceDefinitionMap, ok := definition.(map[string]interface{}); ok {
			namespace, err := readNamespace(namespaceName, namespaceDefinitionMap)
			if err != nil {
				return nil, fmt.Errorf("unable to read namespace [%s]. %w", namespaceName, err)
			}
			namespaces = append(namespaces, *namespace)
		} else {
			return nil, fmt.Errorf("namespace [%s] definition has incorrect format", namespaceName)
		}
	}
	return namespaces, nil
}

func readNamespace(namespaceName string, namespaceDefinitionMap map[string]interface{}) (*schema.Namespace, error) {

	types, err := readTypes(namespaceDefinitionMap[typesKey].(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing [%s] key. %w", typesKey, err)
	}

	procedures, err := readProcedures(namespaceDefinitionMap[proceduresKey].(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing [%s] key. %w", proceduresKey, err)
	}

	return &schema.Namespace{
		SchemaElement: &schema.SchemaElement{
			Id:          readId(namespaceDefinitionMap),
			Name:        namespaceName,
			Description: readDescription(namespaceDefinitionMap),
		},
		Types:      types,
		Procedures: procedures,
	}, nil
}

func readTypes(typesMap map[string]interface{}) ([]schema.TypeDeclaration, error) {
	types := []schema.TypeDeclaration{}
	for typeName, definition := range typesMap {
		definitionMap := definition.(map[string]interface{})
		typeDeclaration := schema.TypeDeclaration{
			SchemaElement: &schema.SchemaElement{
				Id:          readId(definitionMap),
				Name:        typeName,
				Description: readDescription(definitionMap),
			},
			PropertyTypes:  readProperties(definitionMap),
			TypeParameters: readParametrization(definitionMap),
		}
		types = append(types, typeDeclaration)
	}
	return types, nil
}

func readProperties(definitionMap map[string]interface{}) []schema.PropertyType {
	properties := definitionMap[propertiesKey]
	if properties == nil {
		return make([]schema.PropertyType, 0)
	}
	propertiesMap := properties.(map[string]interface{})
	result := make([]schema.PropertyType, 0, len(propertiesMap))
	for propertyName, propertyDefinition := range propertiesMap {
		propertyDefinitionMap := propertyDefinition.(map[string]interface{})
		result = append(result, readProperty(propertyName, propertyDefinitionMap))
	}
	return result
}

func readProperty(propertyName string, propertyDefinition map[string]interface{}) schema.PropertyType {
	return schema.PropertyType{
		TypeName: propertyDefinition[typeKey].(string),
		SchemaElement: &schema.SchemaElement{
			Id:          readId(propertyDefinition),
			Description: readDescription(propertyDefinition),
			Name:        propertyName,
		},
	}
}

func readParametrization(definitionMap map[string]interface{}) []schema.TypeParameter {
	parametrization := definitionMap[parametrizationKey]
	if parametrization == nil {
		return make([]schema.TypeParameter, 0)
	}
	parametrizationMap := parametrization.(map[string]interface{})
	result := make([]schema.TypeParameter, 0, len(parametrizationMap))
	for parameterName, parameterDefinition := range parametrizationMap {
		parameterDefinitionMap := parameterDefinition.(map[string]interface{})
		result = append(result, readParameter(parameterName, parameterDefinitionMap))
	}
	return result
}

func readParameter(parameterName string, parameterDefinition map[string]interface{}) schema.TypeParameter {
	return schema.TypeParameter{
		SchemaElement: &schema.SchemaElement{
			Id:          readId(parameterDefinition),
			Description: readDescription(parameterDefinition),
			Name:        parameterName,
		},
	}
}

func readProcedures(proceduresMap map[string]interface{}) ([]schema.Procedure, error) {
	procedures := make([]schema.Procedure, 0, len(proceduresMap))
	for name, definition := range proceduresMap {
		definitionMap := definition.(map[string]interface{})
		argumentsMap := definitionMap[argumentsKey].(map[string]interface{})
		arguments, err := readProcedureArguments(argumentsMap)
		if err != nil {
			return nil, fmt.Errorf("error occurred while reading procedure [%s] arguments. %w", name, err)
		}
		returnTypeMap := definitionMap[returnTypeKey]
		returnType, err := readReturnType(returnTypeMap.(map[string]interface{}))
		if err != nil {
			return nil, fmt.Errorf("error occurred while reading procedure [%s] return type. %w", name, err)
		}
		procedure := schema.Procedure{
			SchemaElement: &schema.SchemaElement{
				Name:        name,
				Description: readDescription(definitionMap),
				Id:          readId(definitionMap),
			},
			ArgumentTypes: arguments,
			ReturnType:    returnType,
		}
		procedures = append(procedures, procedure)
	}
	return procedures, nil
}

func readProcedureArguments(argumentsMap map[string]interface{}) ([]schema.PropertyType, error) {
	arguments := make([]schema.PropertyType, 0, len(argumentsMap))
	for name, definition := range argumentsMap {
		definitionMap := definition.(map[string]interface{})
		typeName := definitionMap[typeKey]
		if typeName == nil {
			return arguments, fmt.Errorf("type wasn't specified for procedure argument [%s]", name)
		}
		arguments = append(arguments, schema.PropertyType{
			TypeName: typeName.(string),
			SchemaElement: &schema.SchemaElement{
				Name:        name,
				Id:          readId(definitionMap),
				Description: readDescription(definitionMap),
			},
		})
	}
	return arguments, nil
}

func readReturnType(definition map[string]interface{}) (*schema.PropertyType, error) {
	typeName := definition[typeKey]
	if typeName == nil {
		return nil, fmt.Errorf("return type wasn't specified")
	}
	return &schema.PropertyType{
		TypeName: typeName.(string),
		SchemaElement: &schema.SchemaElement{
			Name:        returnTypeName,
			Id:          readId(definition),
			Description: readDescription(definition),
		},
	}, nil
}

func readId(definition map[string]interface{}) string {
	return readString(idKey, defaultId, definition)
}

func readDescription(definition map[string]interface{}) string {
	return readString(descriptionKey, defaultDescription, definition)
}

func readType(typeName string) (string, error) {
	typeNameLength := len(typeName)
	if typeNameLength == 0 {
		return "", fmt.Errorf("incorrect type. Type can't be empty")
	}
	typeNameRunes := []rune(typeName)
	typeNameRunes[0] = unicode.ToUpper(typeNameRunes[0])
	return string(typeNameRunes), nil
}

func readString(key, defaultValue string, definition map[string]interface{}) string {
	value := definition[key]
	if value == nil {
		return defaultValue
	}
	return value.(string)
}
