package dto

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRequest_Success(t *testing.T) {

	request := PositionalRequest{
		Parameters: []interface{}{"argument"},
		Request: &Request{
			Method: "someName",
		},
	}
	jsonb, _ := json.Marshal(request)

	fmt.Print(string(jsonb))
}
