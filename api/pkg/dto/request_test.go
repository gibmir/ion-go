package dto_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"encoding/json"

	"github.com/gibmir/ion-go/api/pkg/dto"
)

const (
	testId            = "testId"
	testMethod        = "testMethod"
	testParameter     = "testParameter"
	testParameterName = "testParameterName"
)

var _ = Describe("Request", func() {
	Describe("Marshalling request", func() {
		Context("with correct positional arguments", func() {
			It("should be a correct json with positional request", func() {
				positionalRequest := dto.Positional{
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
			It("should be a correct json with positional notification", func() {
				positionalNotification := dto.Positional{
					Request: &dto.Request{
						Method:   testMethod,
						Protocol: dto.DefaultJsonRpcProtocolVersion,
					},
					Parameters: []interface{}{testParameter},
				}
				bytes, err := json.Marshal(positionalNotification)
				Expect(err).Should(BeNil())

				json := string(bytes)
				expected := fmt.Sprintf(`{"params":["%s"],"method":"%s","jsonrpc":"%s"}`,
					testParameter, testMethod, dto.DefaultJsonRpcProtocolVersion)
				Expect(json).Should(Equal(expected))
			})
		})
		Context("with correct named arguments", func() {
			It("should be a correct json with named request", func() {
				namedRequest := dto.Named{
					Request: &dto.Request{
						Id:       testId,
						Method:   testMethod,
						Protocol: dto.DefaultJsonRpcProtocolVersion,
					},
					Parameters: map[string]interface{}{
						testParameterName: testParameter,
					},
				}
				bytes, err := json.Marshal(namedRequest)
				Expect(err).Should(BeNil())

				json := string(bytes)
				expected := fmt.Sprintf(`{"params":{"%s":"%s"},"id":"%s","method":"%s","jsonrpc":"%s"}`,
					testParameterName, testParameter, testId, testMethod,
					dto.DefaultJsonRpcProtocolVersion)

				Expect(json).Should(Equal(expected))
			})
			It("should be a correct json with named notification", func() {
				namedRequest := dto.Named{
					Request: &dto.Request{
						Method:   testMethod,
						Protocol: dto.DefaultJsonRpcProtocolVersion,
					},
					Parameters: map[string]interface{}{
						testParameterName: testParameter,
					},
				}
				bytes, err := json.Marshal(namedRequest)
				Expect(err).Should(BeNil())

				json := string(bytes)
				expected := fmt.Sprintf(`{"params":{"%s":"%s"},"method":"%s","jsonrpc":"%s"}`,
					testParameterName, testParameter, testMethod,
					dto.DefaultJsonRpcProtocolVersion)

				Expect(json).Should(Equal(expected))
			})
		})
	})
})
