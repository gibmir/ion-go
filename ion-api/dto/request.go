package dto

type NamedRequest struct{
	Id string `json:"id"`
	Method string `json:"method"`
	Parameters map[string]interface{} `json:"params"`
}

type PositionalRequest struct{
	Id string `json:"id"`
	Method string `json:"method"`
	Parameters []interface{} `json:"params"`
}

type NamedNotification struct{
	Method string `json:"method"`
	Parameters map[string]interface{} `json:"params"`
}

type PositionalNotification struct{
	Method string `json:"method"`
	Parameters []interface{} `json:"params"`
}

type BatchRequest struct{
	Requests []JsonRpcRequest
}

type JsonRpcRequest interface{
}
