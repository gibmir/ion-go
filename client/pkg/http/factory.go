package http

import (
	"errors"
	"net/http"

	api "github.com/gibmir/ion-go/api/core"
	"github.com/gibmir/ion-go/client/pkg/configuration"
	"github.com/gibmir/ion-go/processor"
	"github.com/sirupsen/logrus"
)

const (
	logrusMethodFieldKey = "method"
)

type RequestFactoryEnvironment interface {
	HttpClient() *http.Client
	DefaultUrl() string
	Processor() processor.Processor
}

type HttpRequestFactoryEnvironment struct {
	httpClient *http.Client
	defaultUrl string
	proc       processor.Processor
}

func (env *HttpRequestFactoryEnvironment) HttpClient() *http.Client {
	return env.httpClient
}

func (env *HttpRequestFactoryEnvironment) DefaultUrl() string {
	return env.defaultUrl
}

func (env *HttpRequestFactoryEnvironment) Processor() processor.Processor {
	return env.proc
}

func NewHttpRequestFactory(proc processor.Processor, config *configuration.Configuration) RequestFactoryEnvironment {
	return &HttpRequestFactoryEnvironment{
		httpClient: http.DefaultClient,
		defaultUrl: config.Url,
		proc:       proc,
	}
}

func NewRequest0[R any](
	factory RequestFactoryEnvironment,
	procedure *api.Describer0[R],
) (*HttpRequest0[R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}

	procedureDescription := procedure.Describe()
	request := HttpRequest0[R]{
		HttpRequest: &HttpRequest{
			methodName: procedureDescription.ProcedureName,
			log:        logrus.WithField(logrusMethodFieldKey, procedureDescription.ProcedureName).Logger,
			proc:       factory.Processor(),
			httpSender: &HttpSender{
				httpClient: factory.HttpClient(),
				url:        factory.DefaultUrl(),
			},
		},
	}
	return &request, nil
}

func NewRequest1[T, R any](
	factory RequestFactoryEnvironment,
	procedure *api.Describer1[T, R],
) (*HttpRequest1[T, R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}
	procedureDescription := procedure.Describe()
	request := HttpRequest1[T, R]{
		ArgumentName: procedureDescription.ArgNames[0],
		HttpRequest: &HttpRequest{
			methodName: procedureDescription.ProcedureName,
			log:        logrus.WithField(logrusMethodFieldKey, procedureDescription.ProcedureName).Logger,
			proc:       factory.Processor(),
			httpSender: &HttpSender{
				httpClient: factory.HttpClient(),
				url:        factory.DefaultUrl(),
			},
		},
	}
	return &request, nil
}

func NewRequest2[T1, T2, R any](
	factory RequestFactoryEnvironment,
	procedure *api.Describer2[T1, T2, R],
) (*HttpRequest2[T1, T2, R], error) {
	if factory == nil || procedure == nil {
		return nil, errors.New("factory or procedure are nil")
	}
	procedureDescription := procedure.Describe()
	request := HttpRequest2[T1, T2, R]{
		FirstArgumentName:  procedureDescription.ArgNames[0],
		SecondArgumentName: procedureDescription.ArgNames[1],
		HttpRequest: &HttpRequest{
			methodName: procedureDescription.ProcedureName,
			log:        logrus.WithField(logrusMethodFieldKey, procedureDescription.ProcedureName).Logger,
			proc:       factory.Processor(),
			httpSender: &HttpSender{
				httpClient: factory.HttpClient(),
				url:        factory.DefaultUrl(),
			},
		},
	}
	return &request, nil
}

func NewRequest3[T1, T2, T3, R any](
	factory RequestFactoryEnvironment,
	procedure *api.Describer3[T1, T2, T3, R],
) (*HttpRequest3[T1, T2, T3, R], error) {
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
			log:        logrus.WithField(logrusMethodFieldKey, procedureDescription.ProcedureName).Logger,
			proc:       factory.Processor(),
			httpSender: &HttpSender{
				httpClient: factory.HttpClient(),
				url:        factory.DefaultUrl(),
			},
		},
	}
	return &request, nil
}
