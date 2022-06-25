package generator

import (
	//"bytes"
	"bytes"
	"fmt"
	"testing"
	"text/template"

	//"text/template"

	schema "github.com/gibmir/ion-go/schema/core"
	"github.com/stretchr/testify/assert"
)

const tpl = `package {{ .PackageName }}

func (o {{ .StructName }}) ShallowCopy() {{ .StructName }} {
	return {{ .StructName }}{
		{{- range $field := .Fields }}
		{{ $field }}: o.{{ $field }},
		{{- end }}
	}
}
`

const testTemplate = `
//{{ .Namespace.Description }}
package {{ .Namespace.Name }}

{{- range $procedure := .Namespace.Procedures}}
type $procedure.Name struct {
}
{{- end }}
`

func TestGenerateTypesTemplate_Success(t *testing.T) {
	a := assert.New(t)
	b := bytes.NewBufferString("")
	template.Must(template.New("").Parse(testTemplate)).Execute(b, schema.Namespace{
		SchemaElement: &schema.SchemaElement{
			Name:        "testNamespaceName",
			Description: "testNamespaceDescription",
		},
		Procedures: []schema.Procedure{
			{
				SchemaElement: &schema.SchemaElement{
					Name:        "testProcedureName",
					Description: "testProcedureDescription",
				},
			},
		},
	})
	fmt.Print(b.String())
	a.Equal("", b.String())
	//	var schema *schema.Schema = nil
	//	for i := 0; i < len(schema.Namespaces); i++ {
	//		namespace := schema.Namespaces[i]
	//		temp := template.Must(template.New("").Parse(tpl))
	//		b := bytes.NewBufferString("")
	//		temp.Execute(b, namespace)
	//		a.Equal("", b.String())
	//	}
}
