package core

import (
	"errors"
	"net/http"

	api "github.com/gibmir/ion-go/api/core"
	"github.com/gibmir/ion-go/http-client/configuration"
	"github.com/gibmir/ion-go/pool"
)

type HttpRequestFactory struct {
	bufferPool *pool.BufferPool
	httpClient *http.Client
	defaultUrl string
}

func NewHttpRequestFactory(config *configuration.Configuration) *HttpRequestFactory {
	return &HttpRequestFactory{
		bufferPool: pool.NewBufferPool(config.GetPoolSize(), config.GetBufferLength()),
		httpClient: http.DefaultClient,
		defaultUrl: config.GetUrl(),
	}
}

func NewRequest0[R any](factory *HttpRequestFactory, procedure api.JsonRemoteProcedure0[R]) (*HttpRequest0[R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}

	request := HttpRequest0[R]{
		HttpRequest: &HttpRequest{
			methodName: procedure.GetName(),
			httpSender: &HttpSender{
				bufferPool: factory.bufferPool,
				httpClient: factory.httpClient,
				url:        factory.defaultUrl,
			},
		},
	}
	return &request, nil
}

func NewRequest1[T, R any](factory *HttpRequestFactory, procedure api.JsonRemoteProcedure1[T, R]) (*HttpRequest1[T, R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}
	request := HttpRequest1[T, R]{
		HttpRequest: &HttpRequest{
			methodName: procedure.GetName(),
			httpSender: &HttpSender{
				bufferPool: factory.bufferPool,
				httpClient: factory.httpClient,
				url:        factory.defaultUrl,
			},
		},
	}
	return &request, nil
}

func NewRequest2[T1, T2, R any](factory *HttpRequestFactory, procedure api.JsonRemoteProcedure2[T1, T2, R]) (*HttpRequest2[T1, T2, R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}
	request := HttpRequest2[T1, T2, R]{
		HttpRequest: &HttpRequest{
			methodName: procedure.GetName(),
			httpSender: &HttpSender{
				bufferPool: factory.bufferPool,
				httpClient: factory.httpClient,
				url:        factory.defaultUrl,
			},
		},
	}
	return &request, nil
}

func NewRequest3[T1, T2, T3, R any](factory *HttpRequestFactory, procedure api.JsonRemoteProcedure3[T1, T2, T3, R]) (*HttpRequest3[T1, T2, T3, R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}
	request := HttpRequest3[T1, T2, T3, R]{
		HttpRequest: &HttpRequest{
			methodName: procedure.GetName(),
			httpSender: &HttpSender{
				bufferPool: factory.bufferPool,
				httpClient: factory.httpClient,
				url:        factory.defaultUrl,
			},
		},
	}
	return &request, nil
}
