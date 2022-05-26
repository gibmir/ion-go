package dto

// Positional/named Response
type Response struct{
	Id string `json:"id"`
	Protocol string `json:"jsonrpc"`
	Result interface{} `json:"result"`
}

// Notification response 
type NotificationResponse struct{
	Protocol string `json:"jsonrpc"`
	Result interface{} `json:"result"`
}

// Batch 
type BatchResponse struct{
	Responses []JsonRpcResponse
}

// Errors 
type ErrorResponse struct{
	Id string `json:"id"`
	Protocol string `json:"jsonrpc"`
	Error JsonRpcError `json:"error"`
}

type JsonRpcError struct{
	Code int8 `json:"code"`
	Message string `json:"message"`
}

type JsonRpcResponse interface{
}

