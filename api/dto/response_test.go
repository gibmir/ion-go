package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestResult struct {
	SomeNumber int `json:"x"`
}

func TestResponse_Success(t *testing.T) {
	a := assert.New(t)
	jsonString := `{
		"jsonrpc": "2.0",
		"id": "pepega",
		"result": {
			"x": 22
		}
	}`
	b := []byte(jsonString)
	response := Response[TestResult]{}
	err := json.Unmarshal(b, &response)
	a.Nil(err)
	a.Equal("pepega", response.Id)
	a.Equal("2.0", response.Protocol)
	a.Equal(22, response.Result.SomeNumber)
}
