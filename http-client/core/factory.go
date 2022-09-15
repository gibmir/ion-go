package core

import (
	"errors"
	"net/http"

	api "github.com/gibmir/ion-go/api/core"
	"github.com/gibmir/ion-go/http-client/configuration"
	"github.com/gibmir/ion-go/pool"
)

type RequestFactoryEnvironment interface {
	GetBufferPool() *pool.BufferPool
	GetHttpClient() *http.Client
	GetDefaultUrl() string
}

type HttpRequestFactoryEnvironment struct {
	bufferPool *pool.BufferPool
	httpClient *http.Client
	defaultUrl string
}

func (env *HttpRequestFactoryEnvironment) GetBufferPool() *pool.BufferPool {
	return env.bufferPool
}

func (env *HttpRequestFactoryEnvironment) GetHttpClient() *http.Client {
	return env.httpClient
}

func (env *HttpRequestFactoryEnvironment) GetDefaultUrl() string {
	return env.defaultUrl
}

func NewHttpRequestFactory(config *configuration.Configuration) *HttpRequestFactoryEnvironment {
	return &HttpRequestFactoryEnvironment{
		bufferPool: pool.NewBufferPool(config.GetPoolSize(), config.GetBufferLength()),
		httpClient: http.DefaultClient,
		defaultUrl: config.GetUrl(),
	}
}

func NewRequest0[R any](
	factory RequestFactoryEnvironment,
	procedure *api.Describer0[R],
) (Request0[R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}

	procedureDescription := procedure.Describe()
	request := HttpRequest0[R]{
		HttpRequest: &HttpRequest{
			methodName: procedureDescription.ProcedureName,
			httpSender: &HttpSender{
				bufferPool: factory.GetBufferPool(),
				httpClient: factory.GetHttpClient(),
				url:        factory.GetDefaultUrl(),
			},
		},
	}
	return &request, nil
}

func NewRequest1[T, R any](
	factory RequestFactoryEnvironment,
	procedure *api.Describer1[T, R],
) (Request1[T, R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}
	procedureDescription := procedure.Describe()
	request := HttpRequest1[T, R]{
		ArgumentName: procedureDescription.ArgNames[0],
		HttpRequest: &HttpRequest{
			methodName: procedureDescription.ProcedureName,
			httpSender: &HttpSender{
				bufferPool: factory.GetBufferPool(),
				httpClient: factory.GetHttpClient(),
				url:        factory.GetDefaultUrl(),
			},
		},
	}
	return &request, nil
}

func NewRequest2[T1, T2, R any](
	factory RequestFactoryEnvironment,
	procedure *api.Describer2[T1, T2, R],
) (Request2[T1, T2, R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}
	procedureDescription := procedure.Describe()
	request := HttpRequest2[T1, T2, R]{
		FirstArgumentName:  procedureDescription.ArgNames[0],
		SecondArgumentName: procedureDescription.ArgNames[1],
		HttpRequest: &HttpRequest{
			methodName: procedureDescription.ProcedureName,
			httpSender: &HttpSender{
				bufferPool: factory.GetBufferPool(),
				httpClient: factory.GetHttpClient(),
				url:        factory.GetDefaultUrl(),
			},
		},
	}
	return &request, nil
}

func NewRequest3[T1, T2, T3, R any](
	factory RequestFactoryEnvironment,
	procedure *api.Describer3[T1, T2, T3, R],
) (Request3[T1, T2, T3, R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}
	procedureDescription := procedure.Describe()
	request := HttpRequest3[T1, T2, T3, R]{
		FirstArgumentName:  procedureDescription.ArgNames[0],
		SecondArgumentName: procedureDescription.ArgNames[1],
		ThirdArgumentName:  procedureDescription.ArgNames[2],
		HttpRequest: &HttpRequest{
			methodName: procedureDescription.ProcedureName,
			httpSender: &HttpSender{
				bufferPool: factory.GetBufferPool(),
				httpClient: factory.GetHttpClient(),
				url:        factory.GetDefaultUrl(),
			},
		},
	}
	return &request, nil
}
