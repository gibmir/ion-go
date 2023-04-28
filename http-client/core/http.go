package core

import (
	"github.com/gibmir/ion-go/processor"
	"github.com/sirupsen/logrus"
)

const (
	applicationJsonContentType = "application/json"
)

type HttpRequest struct {
	methodName string
	httpSender *HttpSender
	proc       processor.Processor
	log        *logrus.Logger
}
