package dto

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	testId                    = "testId"
	testSimpleResultValue     = "testSimpleResultValue"
	testCustomResultValue     = 22
	testIncorrectResponseJson = `testIncorrectJson`
)

var (
	testSimpleResultTypeResponseJsonString = fmt.Sprintf(`{
		"jsonrpc": "%s",
		"id": "%s",
		"result": "%s"
	}`, DefaultJsonRpcProtocolVersion, testId, testSimpleResultValue)

	testCustomResultTypeResponseJsonString = fmt.Sprintf(`{
		"jsonrpc": "%s",
		"id": "%s",
		"result": {
			"x": %d
		}
	}`, DefaultJsonRpcProtocolVersion, testId, testCustomResultValue)
)

type TestResult struct {
	SomeNumber int `json:"x"`
}

var _ = Describe("Response", func() {
	Describe("Unmarshal json response", func() {
		Context("with custom result type", func() {
			It("should be a correct response with custom result", func() {
				b := []byte(testCustomResultTypeResponseJsonString)
				response := Response[TestResult]{}

				err := json.Unmarshal(b, &response)

				Expect(err).Should(BeNil())
				Expect(response.Id).Should(Equal(testId))
				Expect(response.Result.SomeNumber).Should(Equal(testCustomResultValue))
				Expect(response.Protocol).Should(Equal(DefaultJsonRpcProtocolVersion))
				Expect(response.Error).Should(BeNil())
			})
		})
		Context("with simple result type", func() {
			It("should be a correct response with simple result", func() {
				b := []byte(testSimpleResultTypeResponseJsonString)
				response := Response[string]{}

				err := json.Unmarshal(b, &response)

				Expect(err).Should(BeNil())
				Expect(response.Id).Should(Equal(testId))
				Expect(response.Result).Should(Equal(testSimpleResultValue))
				Expect(response.Protocol).Should(Equal(DefaultJsonRpcProtocolVersion))
				Expect(response.Error).Should(BeNil())
			})
		})
		Context("with incorrect json", func() {
			It("should be an error(not panic)", func() {

				b := []byte(testIncorrectResponseJson)
				response := Response[string]{}

				err := json.Unmarshal(b, &response)

				Expect(err).ShouldNot(BeNil())
			})
		})
	})
})
