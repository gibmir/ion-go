package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistry_Add(t *testing.T) {

	type testCase struct {
		name          string
		registry      *Registry
		procedureName string
		descriptor    *RpcDescriptor
	}

	cases := []testCase{
		{
			name:          "smoke",
			registry:      &Registry{Registry: map[string]*RpcDescriptor{}},
			procedureName: "testProcedureName",
			descriptor:    &RpcDescriptor{},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			a := assert.New(t)

			a.Len(c.registry.Registry, 0)
			c.registry.Add(c.procedureName, c.descriptor)

			a.Len(c.registry.Registry, 1)
		})
	}
}

func TestRegistry_Descriptor(t *testing.T) {

	type testCase struct {
		name string

		registry      *Registry
		procedureName string
		found         bool
	}

	cases := []testCase{

		{
			name: "smoke. found",
			registry: &Registry{Registry: map[string]*RpcDescriptor{
				"testProcedureName": {},
			},
			},
			procedureName: "testProcedureName",
			found:         true,
		},

		{
			name: "smoke. not found",
			registry: &Registry{Registry: map[string]*RpcDescriptor{
				"testProcedureName": {},
			},
			},
			procedureName: "testAnotherProcedureName",
			found:         false,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			a := assert.New(t)

			_, ok := c.registry.Descriptor(c.procedureName)

			a.Equal(c.found, ok)
		})
	}
}
