package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DisableNatGatewayEcsMetric invokes the vpc.DisableNatGatewayEcsMetric API synchronously
func (client *Client) DisableNatGatewayEcsMetric(request *DisableNatGatewayEcsMetricRequest) (response *DisableNatGatewayEcsMetricResponse, err error) {
	response = CreateDisableNatGatewayEcsMetricResponse()
	err = client.DoAction(request, response)
	return
}

// DisableNatGatewayEcsMetricWithChan invokes the vpc.DisableNatGatewayEcsMetric API asynchronously
func (client *Client) DisableNatGatewayEcsMetricWithChan(request *DisableNatGatewayEcsMetricRequest) (<-chan *DisableNatGatewayEcsMetricResponse, <-chan error) {
	responseChan := make(chan *DisableNatGatewayEcsMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableNatGatewayEcsMetric(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DisableNatGatewayEcsMetricWithCallback invokes the vpc.DisableNatGatewayEcsMetric API asynchronously
func (client *Client) DisableNatGatewayEcsMetricWithCallback(request *DisableNatGatewayEcsMetricRequest, callback func(response *DisableNatGatewayEcsMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableNatGatewayEcsMetricResponse
		var err error
		defer close(result)
		response, err = client.DisableNatGatewayEcsMetric(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DisableNatGatewayEcsMetricRequest is the request struct for api DisableNatGatewayEcsMetric
type DisableNatGatewayEcsMetricRequest struct {
	*requests.RpcRequest
	DryRun       requests.Boolean `position:"Query" name:"DryRun"`
	NatGatewayId string           `position:"Query" name:"NatGatewayId"`
}

// DisableNatGatewayEcsMetricResponse is the response struct for api DisableNatGatewayEcsMetric
type DisableNatGatewayEcsMetricResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableNatGatewayEcsMetricRequest creates a request to invoke DisableNatGatewayEcsMetric API
func CreateDisableNatGatewayEcsMetricRequest() (request *DisableNatGatewayEcsMetricRequest) {
	request = &DisableNatGatewayEcsMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DisableNatGatewayEcsMetric", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableNatGatewayEcsMetricResponse creates a response to parse from DisableNatGatewayEcsMetric response
func CreateDisableNatGatewayEcsMetricResponse() (response *DisableNatGatewayEcsMetricResponse) {
	response = &DisableNatGatewayEcsMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}