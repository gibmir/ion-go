package dto

import "github.com/gibmir/ion-go/api/errors"

const (
	JsonRpcProtocol = "2.0"
	IdKey           = "id"
	JsonRpcKey      = "jsonrpc"
	ResultKey       = "result"
	ErrorKey        = "error"
)

// Response represents json-rpc 2.0 response.
type Response[R any] struct {
	Id       string               `json:"id"`
	Protocol string               `json:"jsonrpc"`
	Result   R                    `json:"result,omitempty"`
	Error    *errors.JsonRpcError `json:"error,omitempty"`
}

// Batch
type BatchResponse struct {
}
