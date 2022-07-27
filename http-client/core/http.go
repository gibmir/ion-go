package core


const (
	applicationJsonContentType = "application/json"
)

type HttpRequest struct {
	methodName string
	httpSender *HttpSender
}
