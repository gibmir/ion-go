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

	simpleTypeTemplateText = `
//{{.Name}} {{.Description}}
type {{.Name}} struct {
  {{range $property:=.PropertyTypes}}
    //{{$property.Name}} {{$property.Description}}
    {{$property.Name}} {{$property.TypeName}},
  {{end}}
}
`
)

var (
	zeroArgumentProcedureTemplate *template.Template = template.Must(template.
					New("zero-arg").
					Parse(zeroArgumentProcedureTemplateText))
	multipleArgumentProcedureTemplate *template.Template = template.Must(template.
						New("multiple-arg").
						Parse(multipleArgumentsTemplateText))
	simpleTypeTemplate *template.Template = template.Must(template.
				New("simple-type").
				Parse(simpleTypeTemplateText))
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
	for _, procedure := range namespace.Procedures {
		err = generateProcedure(&stringBuilder, &procedure)
		if err != nil {
			return "", err
		}
	}
	for _, typeDeclaration := range namespace.Types {
		err = generateType(&stringBuilder, &typeDeclaration)
		if err != nil {
			return "", err
		}
	}
	return stringBuilder.String(), nil
}

func generateType(stringBuilder *strings.Builder, typeDeclaration *schema.TypeDeclaration) error {
	if stringBuilder == nil {
		return fmt.Errorf("stringbuilder is nil")
	}
	if typeDeclaration == nil {
		return fmt.Errorf("type declaration is nil")
	}
	if len(typeDeclaration.TypeParameters) == 0 {
		err := simpleTypeTemplate.Execute(stringBuilder, typeDeclaration)
		if err != nil {
			return err
		}
	}
	return nil
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
