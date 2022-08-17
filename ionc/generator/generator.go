package generator

import (
	"fmt"
	"strings"
	"text/template"

	schema "github.com/gibmir/ion-go/schema/core"
)

const (
	genericMarkSymbol             = "<"
	apiTemplateString             = "package "
	multipleArgumentsTemplateText = `
//{{.Name}} {{.Description}}
type {{.Name}} interface {
  // Call {{.Description}}
  // {{.ReturnType.Name}} {{.ReturnType.Description}}
  Call({{range $arg := .ArgumentTypes}}
    // {{$arg.Name}} {{$arg.Description}}
    {{$arg.Name}}  {{$arg.TypeName}},
  {{end}}){{.ReturnType.TypeName}}
  // Notify {{.Description}}
  Notify({{range $arg := .ArgumentTypes}}
    // {{$arg.Name}} {{$arg.Description}}
    {{$arg.Name}}  {{$arg.TypeName}},
  {{end}})
}
`

	zeroArgumentProcedureTemplateText = `
//{{.Name}} {{.Description}}
type {{.Name}} interface {
  // Call {{.Description}}
  Call(){{.ReturnType.TypeName}}
  // Notify {{.Description}}
  Notify()
}
`
)

var (
	zeroArgumentProcedureTemplate     *template.Template = template.Must(template.New("zero-arg").Parse(zeroArgumentProcedureTemplateText))
	multipleArgumentProcedureTemplate *template.Template = template.Must(template.New("multiple-arg").Parse(multipleArgumentsTemplateText))
)

func GenerateTemplate(apiSchema *schema.Schema) ([]string, error) {
	if apiSchema == nil {
		return nil, fmt.Errorf("schema is nil")
	}

	namespaceCount := len(apiSchema.Namespaces)
	if namespaceCount == 0 {
		return nil, fmt.Errorf("there is no namespaces in schema")
	}

	apiFilesContent := make([]string, namespaceCount)
	var err error
	for i := 0; i < namespaceCount; i++ {
		apiFilesContent[i], err = generateNamespace(&apiSchema.Namespaces[i])
		if err != nil {
			return nil, err
		}
	}
	return apiFilesContent, nil
}

func generateNamespace(namespace *schema.Namespace) (string, error) {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(fmt.Sprintf(`// %s`, namespace.Description))
	stringBuilder.WriteString(fmt.Sprintf("package %s\n", namespace.Name))
	var err error
	for i := 0; i < len(namespace.Procedures); i++ {
		err = generateProcedure(&stringBuilder, &namespace.Procedures[i])
		if err != nil {
			return "", err
		}
	}
	return stringBuilder.String(), nil
}

func generateProcedure(stringBuilder *strings.Builder, procedure *schema.Procedure) error {
	if stringBuilder == nil {
		return fmt.Errorf("stringbuilder is nil")
	}
	if procedure == nil {
		return fmt.Errorf("procedure is nil")
	}
	if len(procedure.ArgumentTypes) > 0 {
		err := multipleArgumentProcedureTemplate.Execute(stringBuilder, procedure)
		if err != nil {
			return err
		}

	} else {
		err := zeroArgumentProcedureTemplate.Execute(stringBuilder, procedure)
		if err != nil {
			return err
		}

	}
	return nil
}

func isGenericProcedure(procedure schema.Procedure) bool {
	if isGenericType(procedure.ReturnType.TypeName) {
		return true
	}
	argumentTypes := procedure.ArgumentTypes
	for _, argumentType := range argumentTypes {
		if isGenericType(argumentType.TypeName) {
			return true
		}
	}
	return false
}

func isGenericType(typeName string) bool {
	return strings.Contains(typeName, genericMarkSymbol)

}

func generateTypesTemplate(apiTypes map[string]schema.TypeDeclaration) *template.Template {
	return nil
}
