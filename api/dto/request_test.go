package dto_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"encoding/json"

	"github.com/gibmir/ion-go/api/dto"
)

const (
	testId        = "testId"
	testMethod    = "testMethod"
	testParameter = "testParameter"
)

var _ = Describe("Request", func() {
	Describe("Marshalling request", func() {
		Context("with correct positional arguments", func() {
			It("should be a correct json", func() {
				positionalRequest := dto.PositionalRequest{
					Request: &dto.Request{
						Id:       testId,
						Method:   testMethod,
						Protocol: dto.DefaultJsonRpcProtocolVersion,
					},
					Parameters: []interface{}{testParameter},
				}
				bytes, err := json.Marshal(positionalRequest)
				Expect(err).Should(BeNil())

				json := string(bytes)
				expected := fmt.Sprintf(`{"params":["%s"],"id":"%s","method":"%s","jsonrpc":"%s"}`,
					testParameter, testId, testMethod, dto.DefaultJsonRpcProtocolVersion)
				Expect(json).Should(Equal(expected))
			})
		})
	})
})
