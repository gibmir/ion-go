package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	api "github.com/gibmir/ion-go/api/pkg/describer"
	"github.com/sirupsen/logrus"

	"github.com/gibmir/ion-go/api/pkg/errors"
	"github.com/gibmir/ion-go/server/internal/handle"
	"github.com/gibmir/ion-go/server/internal/registry"
)

type HttpServer struct {
	logger   *logrus.Logger
	registry *registry.Registry
}

func NewServer(logger *logrus.Logger) *HttpServer {
	return &HttpServer{
		logger:   logger,
		registry: &registry.Registry{Registry: map[string]*registry.RpcDescriptor{}},
	}
}

func NewProcessor0[R any](
	server *HttpServer,
	describer *api.Describer0[R],
	procedure func() (R, error),
) {
	server.registry.Add(describer.Description.ProcedureName, &registry.RpcDescriptor{
		MethodHandle: handle.MethodHandle0[R]{CallFn: procedure},
		Marshaller:   describer,
	})
}

func NewProcessor1[T, R any](
	server *HttpServer,
	describer *api.Describer1[T, R],
	procedure func(t T) (R, error),
) {
	server.registry.Add(describer.Description.ProcedureName, &registry.RpcDescriptor{
		MethodHandle: handle.MethodHandle1[T, R]{CallFn: procedure},
		Marshaller:   describer,
	})
}

func NewProcessor2[T1, T2, R any](
	server *HttpServer,
	describer *api.Describer2[T1, T2, R],
	procedure func(t1 T1, t2 T2) (R, error),
) {
	server.registry.Add(describer.Description.ProcedureName, &registry.RpcDescriptor{
		MethodHandle: handle.MethodHandle2[T1, T2, R]{CallFn: procedure},
		Marshaller:   describer,
	})
}

func NewProcessor3[T1, T2, T3, R any](
	server *HttpServer,
	describer *api.Describer3[T1, T2, T3, R],
	procedure func(t1 T1, t2 T2, t3 T3) (R, error),
) {
	server.registry.Add(describer.Description.ProcedureName, &registry.RpcDescriptor{
		MethodHandle: handle.MethodHandle3[T1, T2, T3, R]{CallFn: procedure},
		Marshaller:   describer,
	})
}

func (s *HttpServer) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.logger.Error(errors.NewInternalError(err.Error()))
	}
	responseBytes, err := s.handleRequest(body)
	if err != nil {
		responseBytes = api.MarshalError(err)
	}
	if len(responseBytes) > 0 {
		n, err := w.Write(responseBytes)
		if err != nil {
			s.logger.Error(err)
		} else {
			s.logger.Debugf("successfully send response(size %d) for request", n)
		}
	}

}

func (s *HttpServer) handleRequest(body []byte) ([]byte, error) {
	// read body(json-rpc request)
	if !json.Valid(body) {
		return nil, errors.NewParseError(fmt.Sprintf("received invalid json with length [%d]", len(body)))
	}

	var bodyMap map[string]*json.RawMessage
	err := json.Unmarshal(body, &bodyMap)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	rawId, ok := bodyMap["id"]
	if ok {
		// request
		var id string
		err := json.Unmarshal(*rawId, &id)
		if err != nil {
			return nil, errors.NewParseError(err.Error())
		}

		s.logger.Debugf("processing request with id [%s]", id)
		method, err := getMethod(bodyMap)
		if err != nil {
			return nil, err
		}

		descriptor, found := s.registry.Descriptor(method)
		if !found {
			return nil, errors.NewMethodNotFoundError(method)
		}
		args, err := descriptor.Marshaller.Unmarshal(*bodyMap["params"])
		if err != nil {
			return nil, errors.NewParseError(err.Error())
		}

		result, err := descriptor.MethodHandle.Call(args)

		responseBytes := descriptor.Marshaller.Marshal(id, result, err)

		return responseBytes, nil
	} else {
		// notification
		method, err := getMethod(bodyMap)
		if err != nil {
			s.logger.Error(err)
			return nil, nil
		}

		descriptor, found := s.registry.Descriptor(method)
		if !found {
			s.logger.Error(errors.NewMethodNotFoundError(method))
			return nil, nil
		}
		args, err := descriptor.Marshaller.Unmarshal(*bodyMap["params"])
		if err != nil {
			s.logger.Error(errors.NewParseError(err.Error()))
			return nil, nil
		}

		result, err := descriptor.MethodHandle.Call(args)
		if err != nil {
			s.logger.Errorf("notification [%s] return error %v", method, err)
		}
		if result != nil {
			s.logger.Warnf("notification [%s] return something", method)
		}
		return nil, nil
	}
}

func getMethod(bodyMap map[string]*json.RawMessage) (string, error) {
	rawMethod, ok := bodyMap["method"]
	if !ok {
		return "", errors.NewInvalidRequestError(fmt.Sprintf("[%s] wasn't provided", "method"))
	}

	var method string

	err := json.Unmarshal(*rawMethod, &method)
	if err != nil {
		return "", errors.NewParseError(err.Error())
	}

	return method, nil
}
