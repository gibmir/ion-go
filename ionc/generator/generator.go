package generator

import (
	"strings"
	"text/template"

	schema "github.com/gibmir/ion-go/ion-schema/core"
)

const (
	genericMarkSymbol = "<"
	apiTemplateString = "package "
)

func GenerateTemplate(apiSchema *schema.Schema) *template.Template {
	apiTemplate := template.Must(template.New("api").Parse(apiTemplateString))
	return apiTemplate
}

func generateProceduresTemplate(apiProcedures []schema.Procedure) *template.Template {
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
