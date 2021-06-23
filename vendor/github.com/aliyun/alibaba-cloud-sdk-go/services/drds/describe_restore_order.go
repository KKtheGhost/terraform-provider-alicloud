package drds

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

// DescribeRestoreOrder invokes the drds.DescribeRestoreOrder API synchronously
func (client *Client) DescribeRestoreOrder(request *DescribeRestoreOrderRequest) (response *DescribeRestoreOrderResponse, err error) {
	response = CreateDescribeRestoreOrderResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRestoreOrderWithChan invokes the drds.DescribeRestoreOrder API asynchronously
func (client *Client) DescribeRestoreOrderWithChan(request *DescribeRestoreOrderRequest) (<-chan *DescribeRestoreOrderResponse, <-chan error) {
	responseChan := make(chan *DescribeRestoreOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRestoreOrder(request)
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

// DescribeRestoreOrderWithCallback invokes the drds.DescribeRestoreOrder API asynchronously
func (client *Client) DescribeRestoreOrderWithCallback(request *DescribeRestoreOrderRequest, callback func(response *DescribeRestoreOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRestoreOrderResponse
		var err error
		defer close(result)
		response, err = client.DescribeRestoreOrder(request)
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

// DescribeRestoreOrderRequest is the request struct for api DescribeRestoreOrder
type DescribeRestoreOrderRequest struct {
	*requests.RpcRequest
	BackupDbNames       string `position:"Query" name:"BackupDbNames"`
	BackupId            string `position:"Query" name:"BackupId"`
	BackupLevel         string `position:"Query" name:"BackupLevel"`
	DrdsInstanceId      string `position:"Query" name:"DrdsInstanceId"`
	PreferredBackupTime string `position:"Query" name:"PreferredBackupTime"`
	BackupMode          string `position:"Query" name:"BackupMode"`
}

// DescribeRestoreOrderResponse is the response struct for api DescribeRestoreOrder
type DescribeRestoreOrderResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	RestoreOrderDO RestoreOrderDO `json:"RestoreOrderDO" xml:"RestoreOrderDO"`
}

// CreateDescribeRestoreOrderRequest creates a request to invoke DescribeRestoreOrder API
func CreateDescribeRestoreOrderRequest() (request *DescribeRestoreOrderRequest) {
	request = &DescribeRestoreOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeRestoreOrder", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRestoreOrderResponse creates a response to parse from DescribeRestoreOrder response
func CreateDescribeRestoreOrderResponse() (response *DescribeRestoreOrderResponse) {
	response = &DescribeRestoreOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
