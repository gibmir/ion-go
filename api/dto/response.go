package dto

const (
	JsonRpcProtocol = "2.0"
	IdKey           = "id"
	JsonRpcKey      = "jsonrpc"
	ResultKey       = "result"
	ErrorKey        = "error"
)

// Response represents json-rpc 2.0 response.
type Response[R any] struct {
	Id       string `json:"id"`
	Protocol string `json:"jsonrpc"`
	Result   R      `json:"result,omitempty"`
	Error    *Error `json:"error,omitempty"`
}

// Batch
type BatchResponse struct {
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
