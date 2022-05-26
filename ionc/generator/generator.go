package generator

import (
	"text/template"

	"github.com/gibmir/ion-go/ion-schema/schema"
)

const (
	apiTemplateString = "PepegaStarege"
)

func GenerateTemplate(apiSchema *schema.Schema) *template.Template {
	apiTemplate := template.Must(template.New("api").Parse(apiTemplateString))
	return apiTemplate
}

func generateProceduresTemplate(apiProcedures []schema.Procedure) *template.Template {
	return nil
}

func generateTypesTemplate(apiTypes map[string]schema.TypeDeclaration) *template.Template {
	return nil
}
