package errors

import "fmt"

const (
	ParseErrorCode          = -32700
	InvalidRequestErrorCode = -32600
	MethodNotFoundErrorCode = -32601
	InvalidParamsErrorCode  = -32602
	InternalErrorCode       = -32603
)

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewParseError constructor.
//Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text
func NewParseError(message string) *JsonRpcError {
	return &JsonRpcError{
		Code:    ParseErrorCode,
		Message: message,
	}
}

// NewInvalidRequestError constructor.
//The JSON sent is not a valid Request object
func NewInvalidRequestError(message string) *JsonRpcError {
	return &JsonRpcError{
		Code:    InvalidRequestErrorCode,
		Message: message,
	}
}

// NewMethodNotFoundError constructor.
//The method does not exist/ is not available
func NewMethodNotFoundError(methodName string) *JsonRpcError {
	return &JsonRpcError{
		Code:    MethodNotFoundErrorCode,
		Message: fmt.Sprintf("[%s] not found", methodName),
	}
}

// NewInvalidParamsError constructor.
//Invalid method parameter(s)
func NewInvalidParamsError(message string) *JsonRpcError {
	return &JsonRpcError{
		Code:    InvalidParamsErrorCode,
		Message: message,
	}
}

// NewInternalError constructor.
//Internal JSON-RPC error
func NewInternalError(message string) *JsonRpcError {
	return &JsonRpcError{
		Code:    InternalErrorCode,
		Message: message,
	}
}

func (e *JsonRpcError) Error() string {
	return fmt.Sprintf("JsonRpcError: {Code: %d, Message: %s}",e.Code,e.Message)
}
