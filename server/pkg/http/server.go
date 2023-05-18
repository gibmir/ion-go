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
	log      *logrus.Entry
	registry *registry.Registry
}

func NewServer(log *logrus.Entry) *HttpServer {
	return &HttpServer{
		log:      log,
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
		s.log.Error(err)
		responseBytes := api.MarshalError(err)
		s.write(responseBytes, w, r)
		return
	}
	responseBytes, err := s.handleRequest(body)
	if err != nil {
		s.log.Error(err)
		responseBytes = api.MarshalError(err)
	}
	s.write(responseBytes, w, r)

}

func (s *HttpServer) write(responseBytes []byte, w http.ResponseWriter, r *http.Request) {
	if len(responseBytes) == 0 {
		s.log.Debugf("empty response for request method [%s] uri [%s] proto [%s]", r.Method, r.RequestURI, r.Proto)
		return
	}
	n, err := w.Write(responseBytes)
	if err != nil {
		s.log.Error(err)
	} else {
		s.log.Debugf("successfully send response(size %d) for request", n)
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

		s.log.Debugf("processing request with id [%s]", id)
		method, err := getMethod(bodyMap)
		if err != nil {
			return nil, err
		}

		descriptor, found := s.registry.Descriptor(method)
		if !found {
			return nil, errors.NewMethodNotFoundError(method)
		}
		var params []byte
		if bodyMap["params"] != nil {
			params = *bodyMap["params"]
		}
		args, err := descriptor.Marshaller.Unmarshal(params)
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
			s.log.Error(err)
			return nil, nil
		}

		descriptor, found := s.registry.Descriptor(method)
		if !found {
			s.log.Error(errors.NewMethodNotFoundError(method))
			return nil, nil
		}

		var params []byte
		if bodyMap["params"] != nil {
			params = *bodyMap["params"]
		}

		args, err := descriptor.Marshaller.Unmarshal(params)
		if err != nil {
			s.log.Error(errors.NewParseError(err.Error()))
			return nil, nil
		}

		result, err := descriptor.MethodHandle.Call(args)
		if err != nil {
			s.log.Errorf("notification [%s] return error %v", method, err)
		}
		if result != nil {
			s.log.Warnf("notification [%s] return something", method)
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
