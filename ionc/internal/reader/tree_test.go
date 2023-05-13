package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeName(t *testing.T) {
	type testCase struct {
		name             string
		typeName         string
		expectedTypeName string
	}

	cases := []testCase{
		{
			name:             "check simple string",
			typeName:         "string",
			expectedTypeName: "string",
		},

		{
			name:             "check parametrization with simple string",
			typeName:         "TestType<string>",
			expectedTypeName: "TestType[string]",
		},

		{
			name:             "check parametrization with few simple strings",
			typeName:         "TestType<string,string>",
			expectedTypeName: "TestType[string,string]",
		},

		{
			name:             "check parametrization with parametrized type",
			typeName:         "TestType<ParametrizedType<string>>",
			expectedTypeName: "TestType[ParametrizedType[string]]",
		},

		{
			name:             "hardmode",
			typeName:         "a<b,c<d,e>,f<g,k<l>>>",
			expectedTypeName: "A[B,C[D,E],F[G,K[L]]]",
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			a := assert.New(t)

			tree := NewTypeTree(c.typeName)

			typeName := ToTypeName(tree)

			a.Equal(c.expectedTypeName, typeName)
		})
	}
}
