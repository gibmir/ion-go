package dto

const (
	DefaultJsonRpcProtocolVersion = "2.0"
)

type Positional struct {
	// Parameters is procedure arguments array representation
	Parameters []interface{} `json:"params,omitempty"`
	*Request   `json:",inline"`
}

type Named struct {
	// Parameters is procedure arguments map representation
	Parameters map[string]interface{} `json:"params,omitempty"`
	*Request   `json:",inline"`
}

type Request struct {
	// Id represents request id. It should be equal to nil for notifications
	Id string `json:"id,omitempty"`
	// Method represents procedure name
	Method string `json:"method"`
	// Protocol represents json-rpc protocol version. Should be equal to 2.0
	Protocol string `json:"jsonrpc"`
}

type BatchRequest struct {
}
