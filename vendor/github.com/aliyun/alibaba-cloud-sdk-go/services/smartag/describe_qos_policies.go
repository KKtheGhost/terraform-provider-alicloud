package smartag

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

// DescribeQosPolicies invokes the smartag.DescribeQosPolicies API synchronously
func (client *Client) DescribeQosPolicies(request *DescribeQosPoliciesRequest) (response *DescribeQosPoliciesResponse, err error) {
	response = CreateDescribeQosPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeQosPoliciesWithChan invokes the smartag.DescribeQosPolicies API asynchronously
func (client *Client) DescribeQosPoliciesWithChan(request *DescribeQosPoliciesRequest) (<-chan *DescribeQosPoliciesResponse, <-chan error) {
	responseChan := make(chan *DescribeQosPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeQosPolicies(request)
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

// DescribeQosPoliciesWithCallback invokes the smartag.DescribeQosPolicies API asynchronously
func (client *Client) DescribeQosPoliciesWithCallback(request *DescribeQosPoliciesRequest, callback func(response *DescribeQosPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeQosPoliciesResponse
		var err error
		defer close(result)
		response, err = client.DescribeQosPolicies(request)
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

// DescribeQosPoliciesRequest is the request struct for api DescribeQosPolicies
type DescribeQosPoliciesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	QosPolicyId          string           `position:"Query"`
	Description          string           `position:"Query"`
	PageNumber           requests.Integer `position:"Query"`
	PageSize             requests.Integer `position:"Query"`
	QosId                string           `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	Priority             requests.Integer `position:"Query"`
}

// DescribeQosPoliciesResponse is the response struct for api DescribeQosPolicies
type DescribeQosPoliciesResponse struct {
	*responses.BaseResponse
	TotalCount  int                              `json:"TotalCount" xml:"TotalCount"`
	RequestId   string                           `json:"RequestId" xml:"RequestId"`
	PageSize    int                              `json:"PageSize" xml:"PageSize"`
	PageNumber  int                              `json:"PageNumber" xml:"PageNumber"`
	QosPolicies QosPoliciesInDescribeQosPolicies `json:"QosPolicies" xml:"QosPolicies"`
}

// CreateDescribeQosPoliciesRequest creates a request to invoke DescribeQosPolicies API
func CreateDescribeQosPoliciesRequest() (request *DescribeQosPoliciesRequest) {
	request = &DescribeQosPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeQosPolicies", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeQosPoliciesResponse creates a response to parse from DescribeQosPolicies response
func CreateDescribeQosPoliciesResponse() (response *DescribeQosPoliciesResponse) {
	response = &DescribeQosPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
