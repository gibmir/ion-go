package dto

import ()

const (
	DefaultJsonRpcProtocolVersion = "2.0"
)

type PositionalRequest struct {
	Parameters []interface{} `json:"params,omitempty"`
	*Request   `json:",inline"`
}

type NamedRequest struct {
	Parameters map[string]interface{} `json:"params,omitempty"`
	*Request   `json:",inline"`
}

type Request struct {
	Id       string `json:"id,omitempty"`
	Method   string `json:"method"`
	Protocol string `json:"jsonrpc"`
}

type Parameters struct {
	NamedParameters      map[string]interface{} `json:",inline,omitempty"`
	PositionalParameters []interface{}          `json:",inline,omitempty"`
}

type BatchRequest struct {
}
