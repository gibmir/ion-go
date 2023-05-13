package core

import (
	"testing"
)

type CustomType struct {
	CustomString string
	CustomInt    int
}

func TestTypeName(t *testing.T) {
	type testCase struct {
		name   string
		json   string
		result *CustomType
	}

	cases := []testCase{
		{
			name: "smoke",
			json: `{"a":"p","b":"o","c":"g"}`,
			result: &CustomType{
				CustomString: "pogU",
				CustomInt:    123,
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {

			describer3 := Describer3[string, string, string, string]{
				Describer: &Describer{
					Description: &ProcedureDescription{
						ArgNames: []string{"a", "b", "c"},
					},
				},
			}
			r1, r2, r3, err := describer3.unmarshal([]byte(c.json))

			if err != nil {
				t.Fail()
			}
			if r1 == nil || r2 == nil || r3 == nil {
				t.Fatal(r1, r2, r3, err)
			}

		})
	}
}
