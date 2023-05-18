package http

import (
	"testing"

	api "github.com/gibmir/ion-go/api/pkg/describer"
	"github.com/gibmir/ion-go/server/internal/handle"
	"github.com/gibmir/ion-go/server/internal/registry"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandleRequest(t *testing.T) {
	type testCase struct {
		name               string
		server             *HttpServer
		body               string
		notification       bool
		expectedResponse   string
		expectedErrContent string
	}

	cases := []testCase{

		{
			name: "check request with no arg handling. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer0[string]{},
							MethodHandle: handle.MethodHandle0[string]{
								CallFn: func() (string, error) {
									return "testString", nil
								},
							},
						},
					},
				},
			},
			body:             `{"jsonrpc":"2.0","method":"testProcedure","id":"1"}`,
			expectedResponse: `{"id":"1","jsonrpc":"2.0","result":"testString"}`,
		},

		{
			name: "check incorrect json error. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
			},
			body:             `invalid json`,
			expectedErrContent: "invalid",
		},


		{
			name: "check notification with no arg handling. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer0[string]{},
							MethodHandle: handle.MethodHandle0[string]{
								CallFn: func() (string, error) {
									return "testString", nil
								},
							},
						},
					},
				},
			},
			body:         `{"jsonrpc":"2.0","method":"testProcedure"}`,
			notification: true,
		},

		{
			name: "check notification with one arg handling. positional args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer1[string, string]{},
							MethodHandle: handle.MethodHandle1[string, string]{
								CallFn: func(arg string) (string, error) {
									return arg + "testString", nil
								},
							},
						},
					},
				},
			},
			body:         `{"jsonrpc":"2.0","method":"testProcedure","params":["arg"]}`,
			notification: true,
		},

		{
			name: "check request with one arg handling. positional args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer1[string, string]{},
							MethodHandle: handle.MethodHandle1[string, string]{
								CallFn: func(arg string) (string, error) {
									return arg + "testString", nil
								},
							},
						},
					},
				},
			},
			body:             `{"jsonrpc":"2.0","method":"testProcedure","id":"1","params":["arg"]}`,
			expectedResponse: `{"id":"1","jsonrpc":"2.0","result":"argtestString"}`,
		},

		{
			name: "check request with one arg handling. named args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer1[string, string]{
								Describer: &api.Describer{
									Description: &api.ProcedureDescription{
										ArgNames: []string{
											"testArg",
										},
									},
								},
							},
							MethodHandle: handle.MethodHandle1[string, string]{
								CallFn: func(arg string) (string, error) {
									return arg + "testString", nil
								},
							},
						},
					},
				},
			},
			body:             `{"jsonrpc":"2.0","method":"testProcedure","id":"1","params":{"testArg":"arg"}}`,
			expectedResponse: `{"id":"1","jsonrpc":"2.0","result":"argtestString"}`,
		},

		{
			name: "check notification with two arg handling. positional args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer2[string, string, string]{},
							MethodHandle: handle.MethodHandle2[string, string, string]{
								CallFn: func(arg1, arg2 string) (string, error) {
									return arg1 + arg2 + "testString", nil
								},
							},
						},
					},
				},
			},
			body:         `{"jsonrpc":"2.0","method":"testProcedure","params":["arg1","arg2"]}`,
			notification: true,
		},

		{
			name: "check request with two arg handling. positional args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer2[string, string, string]{},
							MethodHandle: handle.MethodHandle2[string, string, string]{
								CallFn: func(arg1, arg2 string) (string, error) {
									return arg1 + arg2 + "testString", nil
								},
							},
						},
					},
				},
			},
			body:             `{"jsonrpc":"2.0","method":"testProcedure","id":"1","params":["arg1","arg2"]}`,
			expectedResponse: `{"id":"1","jsonrpc":"2.0","result":"arg1arg2testString"}`,
		},

		{
			name: "check request with two arg handling. named args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer2[string, string, string]{
								Describer: &api.Describer{
									Description: &api.ProcedureDescription{
										ArgNames: []string{
											"testArg1",
											"testArg2",
										},
									},
								},
							},
							MethodHandle: handle.MethodHandle2[string, string, string]{
								CallFn: func(arg1, arg2 string) (string, error) {
									return arg1 + arg2 + "testString", nil
								},
							},
						},
					},
				},
			},
			body:             `{"jsonrpc":"2.0","method":"testProcedure","id":"1","params":{"testArg1":"arg1","testArg2":"arg2"}}`,
			expectedResponse: `{"id":"1","jsonrpc":"2.0","result":"arg1arg2testString"}`,
		},

		{
			name: "check notification with three arg handling. positional args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer3[string, string, string, string]{},
							MethodHandle: handle.MethodHandle3[string, string, string, string]{
								CallFn: func(arg1, arg2, arg3 string) (string, error) {
									return arg1 + arg2 + arg3 + "testString", nil
								},
							},
						},
					},
				},
			},
			body:         `{"jsonrpc":"2.0","method":"testProcedure","params":["arg1","arg2","arg3"]}`,
			notification: true,
		},

		{
			name: "check request with three arg handling. positional args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer3[string, string, string, string]{},
							MethodHandle: handle.MethodHandle3[string, string, string, string]{
								CallFn: func(arg1, arg2, arg3 string) (string, error) {
									return arg1 + arg2 + arg3 + "testString", nil
								},
							},
						},
					},
				},
			},
			body:             `{"jsonrpc":"2.0","method":"testProcedure","id":"1","params":["arg1","arg2","arg3"]}`,
			expectedResponse: `{"id":"1","jsonrpc":"2.0","result":"arg1arg2arg3testString"}`,
		},

		{
			name: "check request with three arg handling. named args. smoke",
			server: &HttpServer{
				log: logrus.NewEntry(logrus.StandardLogger()),
				registry: &registry.Registry{
					Registry: map[string]*registry.RpcDescriptor{
						"testProcedure": {
							Marshaller: &api.Describer3[string, string, string, string]{
								Describer: &api.Describer{
									Description: &api.ProcedureDescription{
										ArgNames: []string{
											"testArg1",
											"testArg2",
											"testArg3",
										},
									},
								},
							},
							MethodHandle: handle.MethodHandle3[string, string, string, string]{
								CallFn: func(arg1, arg2, arg3 string) (string, error) {
									return arg1 + arg2 + arg3 + "testString", nil
								},
							},
						},
					},
				},
			},
			body:             `{"jsonrpc":"2.0","method":"testProcedure","id":"1","params":{"testArg1":"arg1","testArg2":"arg2","testArg3":"arg3"}}`,
			expectedResponse: `{"id":"1","jsonrpc":"2.0","result":"arg1arg2arg3testString"}`,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			a := assert.New(t)

			response, err := c.server.handleRequest([]byte(c.body))

			if c.expectedErrContent != "" {
				a.Contains(err.Error(), c.expectedErrContent)
			} else {
				if len(response) == 0 {
					a.True(c.notification)
				} else {
					a.Equal(c.expectedResponse, string(response))
				}
			}
		})
	}
}

func TestNewProcessor0(t *testing.T) {
	a := assert.New(t)
	server := NewServer(logrus.NewEntry(logrus.StandardLogger()))

	// create describer(or use ionc and json schema)
	describer := &api.Describer0[string]{
		Describer: &api.Describer{
			Description: &api.ProcedureDescription{
				ProcedureName: "testProcedureName",
			},
		},
	}

	// add processor for procedure to server
	NewProcessor0(server, describer, nil)

	descriptor, ok := server.registry.Descriptor("testProcedureName")

	a.True(ok)
	a.Equal(describer, descriptor.Marshaller)
}

func TestNewProcessor1(t *testing.T) {
	a := assert.New(t)
	server := NewServer(logrus.NewEntry(logrus.StandardLogger()))

	// create describer(or use ionc and json schema)
	describer := &api.Describer1[string, string]{
		Describer: &api.Describer{
			Description: &api.ProcedureDescription{
				ProcedureName: "testProcedureName",
			},
		},
	}

	// add processor for procedure to server
	NewProcessor1(server, describer, nil)

	descriptor, ok := server.registry.Descriptor("testProcedureName")

	a.True(ok)
	a.Equal(describer, descriptor.Marshaller)
}

func TestNewProcessor2(t *testing.T) {
	a := assert.New(t)
	server := NewServer(logrus.NewEntry(logrus.StandardLogger()))

	// create describer(or use ionc and json schema)
	describer := &api.Describer2[string, string, string]{
		Describer: &api.Describer{
			Description: &api.ProcedureDescription{
				ProcedureName: "testProcedureName",
			},
		},
	}

	// add processor for procedure to server
	NewProcessor2(server, describer, nil)

	descriptor, ok := server.registry.Descriptor("testProcedureName")

	a.True(ok)
	a.Equal(describer, descriptor.Marshaller)
}

func TestNewProcessor3(t *testing.T) {
	a := assert.New(t)
	server := NewServer(logrus.NewEntry(logrus.StandardLogger()))

	// create describer(or use ionc and json schema)
	describer := &api.Describer3[string, string, string, string]{
		Describer: &api.Describer{
			Description: &api.ProcedureDescription{
				ProcedureName: "testProcedureName",
			},
		},
	}

	// add processor for procedure to server
	NewProcessor3(server, describer, nil)

	descriptor, ok := server.registry.Descriptor("testProcedureName")

	a.True(ok)
	a.Equal(describer, descriptor.Marshaller)
}
