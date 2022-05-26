package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTypesTemplate_Success(t *testing.T) {
	a := assert.New(t)
	temp := generateProceduresTemplate(nil)
	a.Nil(temp)
}
