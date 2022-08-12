package core

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AppProcedure interface {
	JsonRemoteProcedure0[CoolString]
}

type AppProcedureImpl struct {
}

func (api *AppProcedureImpl) Call() CoolString {
	return "arolf"
}

func (api *AppProcedureImpl) Notify() {
}

type CoolString string

func TestResponse_Success(t *testing.T) {
	a := assert.New(t)

	var app AppProcedure = &AppProcedureImpl{}
	app.Call()
	procedureType := reflect.TypeOf(new(AppProcedure)).Elem()
	method, found := procedureType.MethodByName("Call")
	if !found {
		panic("pizdec")
	}

	returnArgType := method.Type.Out(0)

	fmt.Printf("returnArgType.Name(): %v\n", returnArgType.Name())
	a.NotNil(procedureType)
	a.Nil(nil)
}
