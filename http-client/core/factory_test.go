package core

import (
	"testing"

	"github.com/gibmir/ion-go/api/core"
)

func TestGenerate(t *testing.T) {
	hrf := HttpRequestFactory{}
	var d = core.Describer0[string]{
		ReturnType: &core.Type[string]{},

		Describer: &core.Describer{
			Description: &core.ProcedureDescription{
				ProcedureName: "someName",
			},
		},
	}
	_, _ = NewRequest0(&hrf, &d)

}
