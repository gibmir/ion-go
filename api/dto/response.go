package dto

const (
	JsonRpcProtocol = "2.0"
	IdKey           = "id"
	JsonRpcKey      = "jsonrpc"
	ResultKey       = "result"
	ErrorKey        = "error"
)

// Positional/named Response
type Response[R any] struct {
	Id       string `json:"id"`
	Protocol string `json:"jsonrpc"`
	Result   R      `json:"result"`
	Error    Error  `json:"error,omitempty"`
}

// Notification response
type NotificationResponse struct {
	Protocol string      `json:"jsonrpc"`
	Result   interface{} `json:"result"`
}

// Batch
type BatchResponse struct {
	Responses []JsonRpcResponse
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JsonRpcResponse interface {
}
