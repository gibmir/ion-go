package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	schema "github.com/gibmir/ion-go/schema/core"
)

const (
	defaultImport       = "github.com/gibmir/ion-go/api/core"
	apiTemplateString   = "package "
	zeroArgTemplateText = `
//{{.Name}} {{.Description}}
{{.Name}}Describer = api.Describer0[
  // {{.ReturnType.Description}}
  {{.ReturnType.TypeName}}
]{
  ReturnType: &api.Type[{{.ReturnType.TypeName}}]{},
  Describer: &api.ProcedureDescription{
    ProcedureName: "{{.Name}}",
  },
}
`

	simpleTypeTemplateText = `
//{{.Name}} {{.Description}}
type {{.Name}} struct {
  {{range $property:=.PropertyTypes}}
    //{{$property.Name}} {{$property.Description}}
    {{$property.Name}} {{$property.TypeName}}
  {{end}}
}
`
)

var (
	zeroArgProcedureTemplate *template.Template = template.Must(template.
					New("zero-arg").
					Parse(zeroArgTemplateText))
	simpleTypeTemplate *template.Template = template.Must(template.
				New("simple-type").
				Parse(simpleTypeTemplateText))
)

func GenerateTemplate(apiSchema *schema.Schema, outDir string) error {
	if apiSchema == nil {
		return fmt.Errorf("schema is nil")
	}

	namespaceCount := len(apiSchema.Namespaces)
	if namespaceCount == 0 {
		return fmt.Errorf("there is no namespaces in schema")
	}

	for i := 0; i < namespaceCount; i++ {
		apiFilesContent, err := generateNamespace(&apiSchema.Namespaces[i])
		if err != nil {
			return fmt.Errorf("unable to generate file for namespace [%s]. %w",
				apiSchema.Namespaces[i].Name, err)
		}
		outDirectory := filepath.Join(outDir, apiSchema.Namespaces[i].Name)
		err = os.MkdirAll(outDirectory, os.ModePerm)
		if err != nil {
			return fmt.Errorf("unable to generate out directory [%s] for namespace [%s]. %w",
				outDirectory, apiSchema.Namespaces[i].Name, err)
		}
		filePath := filepath.Join(outDirectory, apiSchema.Namespaces[i].Name) + ".go"
		f, err := os.Create(filePath)
		defer f.Close()
		if err != nil {
			return fmt.Errorf("unable to create file for namespace [%s] in [%s]. %w",
				apiSchema.Namespaces[i].Name, filePath, err)
		}
		_, err = f.WriteString(apiFilesContent)
		if err != nil {
			return fmt.Errorf("unable to write content to file for namespace [%s] in [%s]. %w",
				apiSchema.Namespaces[i].Name, filePath, err)
		}
	}
	return nil
}

func generateNamespace(namespace *schema.Namespace) (string, error) {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(fmt.Sprintf("// %s %s\n", namespace.Name, namespace.Description))
	stringBuilder.WriteString(fmt.Sprintf("package %s\n", namespace.Name))
	stringBuilder.WriteString(fmt.Sprintf("import api \"%s\"\n", defaultImport))
	var err error
	stringBuilder.WriteString("var (\n")
	for _, procedure := range namespace.Procedures {
		err = generateProcedure(&stringBuilder, &procedure)
		if err != nil {
			return "", err
		}
	}
	stringBuilder.WriteString("\n)")
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
	argumentsCount := len(procedure.ArgumentTypes)
	switch argumentsCount {
	case 0:
		generateProcedure0(stringBuilder, procedure)
	case 1:
		generateProcedure1(stringBuilder, procedure)
	case 2:
		generateProcedure2(stringBuilder, procedure)
	case 3:
		generateProcedure3(stringBuilder, procedure)
	default:
		return fmt.Errorf("incorrect procedure [%s] arguments count [%d]", procedure.Name, argumentsCount)

	}
	return nil
}

func generateProcedure0(stringBuilder *strings.Builder, procedure *schema.Procedure) {
	zeroArgProcedureTemplate.Execute(stringBuilder, procedure)
}

func generateProcedure1(stringBuilder *strings.Builder, procedure *schema.Procedure) {
	procedureDescriber := fmt.Sprintf(`
%s := api.Describer1[%s, %s]{
  FirstArgument: &api.Type[%s]{},
  Describer0: &api.Describer0[%s]{
    ReturnType: &api.Type[%s]{},
    Describer: &api.Describer{
      Description: &api.ProcedureDescription{
	ProcedureName: "%s",
        ArgNames: []string{
	  "%s",
	},
      },
    },
  },
}
`, procedure.Name+"Describer", procedure.ArgumentTypes[0].TypeName, procedure.ReturnType.TypeName,
		procedure.ArgumentTypes[0].TypeName, procedure.ReturnType.TypeName,
		procedure.ReturnType.TypeName, procedure.Name, procedure.ArgumentTypes[0].Name)

	stringBuilder.WriteString(procedureDescriber)
}

func generateProcedure2(stringBuilder *strings.Builder, procedure *schema.Procedure) {
	procedureDescriber := fmt.Sprintf(`
%s := api.Describer2[%s, %s, %s]{
  FirstArgument: &api.Type[%s]{},
  SecondArgument: &api.Type[%s]{},
  Describer0: &api.Describer0[%s]{
    ReturnType: &api.Type[%s]{},
    Describer: &api.Describer{
      Description: &api.ProcedureDescription{
	ProcedureName: "%s",
        ArgNames: []string{
	  "%s",
	  "%s",
	},
      },
    },
  },
}
`, procedure.Name+"Describer", procedure.ArgumentTypes[0].TypeName, procedure.ArgumentTypes[1].TypeName,
		procedure.ReturnType.TypeName, procedure.ArgumentTypes[0].TypeName,
		procedure.ArgumentTypes[1].TypeName, procedure.ReturnType.TypeName,
		procedure.ReturnType.TypeName, procedure.Name, procedure.ArgumentTypes[0].Name,
		procedure.ArgumentTypes[1].TypeName)

	stringBuilder.WriteString(procedureDescriber)
}

func generateProcedure3(stringBuilder *strings.Builder, procedure *schema.Procedure) {
	procedureDescriber := fmt.Sprintf(`
%s := api.Describer3[%s, %s, %s,%s]{
  FirstArgument: &api.Type[%s]{},
  SecondArgument: &api.Type[%s]{},
  ThirdArgument: &api.Type[%s]{},
  Describer0: &api.Describer0[%s]{
    ReturnType: &api.Type[%s]{},
    Describer: &api.Describer{
      Description: &api.ProcedureDescription{
	ProcedureName: "%s",
        ArgNames: []string{
	  "%s",
	  "%s",
	  "%s",
	},
      },
    },
  },
}
`, procedure.Name+"Describer", procedure.ArgumentTypes[0].TypeName, procedure.ArgumentTypes[1].TypeName,
		procedure.ArgumentTypes[2].TypeName, procedure.ReturnType.TypeName,
		procedure.ArgumentTypes[0].TypeName, procedure.ArgumentTypes[1].TypeName,
		procedure.ArgumentTypes[2].TypeName, procedure.ReturnType.TypeName,
		procedure.ReturnType.TypeName, procedure.Name, procedure.ArgumentTypes[0].Name,
		procedure.ArgumentTypes[1].TypeName, procedure.ArgumentTypes[2].TypeName)

	stringBuilder.WriteString(procedureDescriber)
}
