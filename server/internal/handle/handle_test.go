package handle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {

	type testCase struct {
		name               string
		handle             MethodHandle
		args               []interface{}
		expectedResult     interface{}
		expectedErrContent string
	}

	cases := []testCase{

		// zero arg
		{
			name: "check 0 arg handle",
			handle: MethodHandle0[string]{
				CallFn: func() (string, error) {
					return "testString", nil
				},
			},
			expectedResult: "testString",
		},

		{
			name: "check 0 arg handle error",
			handle: MethodHandle0[string]{
				CallFn: func() (string, error) {
					return "", fmt.Errorf("testError")
				},
			},
			expectedErrContent: "testError",
		},

		//single arg
		{
			name: "check 1 arg handle",
			handle: MethodHandle1[string, string]{
				CallFn: func(arg string) (string, error) {
					return arg + "testString", nil
				},
			},
			args:           []interface{}{"testArg"},
			expectedResult: "testArgtestString",
		},

		{
			name: "check 1 arg handle. Incorrect args count",
			handle: MethodHandle1[string, string]{
				CallFn: func(arg string) (string, error) {
					return arg + "testString", nil
				},
			},
			expectedErrContent: "incorrect length",
		},

		{
			name: "check 1 arg handle error",
			handle: MethodHandle1[string, string]{
				CallFn: func(arg string) (string, error) {
					return "", fmt.Errorf("testError")
				},
			},
			args:               []interface{}{""},
			expectedErrContent: "testError",
		},

		{
			name: "check 1 arg handle type assertion error",
			handle: MethodHandle1[string, string]{
				CallFn: func(arg string) (string, error) {
					return "", nil
				},
			},
			args:               []interface{}{1},
			expectedErrContent: fmt.Sprintf("%T", 1),
		},

		// two arg
		{
			name: "check 2 arg handle",
			handle: MethodHandle2[string, string, string]{
				CallFn: func(arg1, arg2 string) (string, error) {
					return arg1 + arg2 + "testString", nil
				},
			},
			args:           []interface{}{"testArg1", "testArg2"},
			expectedResult: "testArg1testArg2testString",
		},

		{
			name: "check 2 arg handle. Incorrect args count",
			handle: MethodHandle2[string, string, string]{
				CallFn: func(arg1, arg2 string) (string, error) {
					return "testString", nil
				},
			},
			expectedErrContent: "incorrect length",
		},

		{
			name: "check 2 arg handle error",
			handle: MethodHandle2[string, string, string]{
				CallFn: func(arg1, arg2 string) (string, error) {
					return "", fmt.Errorf("testError")
				},
			},
			args:               []interface{}{"", ""},
			expectedErrContent: "testError",
		},

		{
			name: "check 2 first arg handle type assertion error",
			handle: MethodHandle2[string, int, string]{
				CallFn: func(arg1 string, arg2 int) (string, error) {
					return "", nil
				},
			},
			args:               []interface{}{1, 1},
			expectedErrContent: fmt.Sprintf("%T", 1),
		},

		{
			name: "check 2 second arg handle type assertion error",
			handle: MethodHandle2[string, int, string]{
				CallFn: func(arg1 string, arg2 int) (string, error) {
					return "", nil
				},
			},
			args:               []interface{}{"", ""},
			expectedErrContent: fmt.Sprintf("%T", ""),
		},

		// three arg
		{
			name: "check 3 arg handle",
			handle: MethodHandle3[string, string, string, string]{
				CallFn: func(arg1, arg2, arg3 string) (string, error) {
					return arg1 + arg2 + arg3 + "testString", nil
				},
			},
			args:           []interface{}{"testArg1", "testArg2", "testArg3"},
			expectedResult: "testArg1testArg2testArg3testString",
		},

		{
			name: "check 3 arg handle. Incorrect args count",
			handle: MethodHandle3[string, string, string, string]{
				CallFn: func(arg1, arg2, arg3 string) (string, error) {
					return "testString", nil
				},
			},
			expectedErrContent: "incorrect length",
		},

		{
			name: "check 3 arg handle error",
			handle: MethodHandle3[string, string, string, string]{
				CallFn: func(arg1, arg2, arg3 string) (string, error) {
					return "", fmt.Errorf("testError")
				},
			},
			args:               []interface{}{"", "", ""},
			expectedErrContent: "testError",
		},

		{
			name: "check 3 first arg handle type assertion error",
			handle: MethodHandle3[string, int, bool, string]{
				CallFn: func(arg1 string, arg2 int, arg3 bool) (string, error) {
					return "", nil
				},
			},
			args:               []interface{}{1, 1, 1},
			expectedErrContent: fmt.Sprintf("%T", 1),
		},

		{
			name: "check 3 second arg handle type assertion error",
			handle: MethodHandle3[string, int, bool, string]{
				CallFn: func(arg1 string, arg2 int, arg3 bool) (string, error) {
					return "", nil
				},
			},
			args:               []interface{}{"", "", 1},
			expectedErrContent: fmt.Sprintf("%T", ""),
		},

		{
			name: "check 3 third arg handle type assertion error",
			handle: MethodHandle3[string, int, bool, string]{
				CallFn: func(arg1 string, arg2 int, arg3 bool) (string, error) {
					return "", nil
				},
			},
			args:               []interface{}{"", 1, 1},
			expectedErrContent: fmt.Sprintf("%T", 1),
		},

	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			a := assert.New(t)

			result, err := c.handle.Call(c.args)

			if c.expectedErrContent != "" {
				a.Contains(err.Error(), c.expectedErrContent)
			} else {
				a.Equal(c.expectedResult, result)
			}
		})
	}
}
