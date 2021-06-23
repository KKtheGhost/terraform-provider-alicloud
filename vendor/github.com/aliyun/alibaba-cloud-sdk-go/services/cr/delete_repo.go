package cr

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

// DeleteRepo invokes the cr.DeleteRepo API synchronously
func (client *Client) DeleteRepo(request *DeleteRepoRequest) (response *DeleteRepoResponse, err error) {
	response = CreateDeleteRepoResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRepoWithChan invokes the cr.DeleteRepo API asynchronously
func (client *Client) DeleteRepoWithChan(request *DeleteRepoRequest) (<-chan *DeleteRepoResponse, <-chan error) {
	responseChan := make(chan *DeleteRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRepo(request)
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

// DeleteRepoWithCallback invokes the cr.DeleteRepo API asynchronously
func (client *Client) DeleteRepoWithCallback(request *DeleteRepoRequest, callback func(response *DeleteRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRepoResponse
		var err error
		defer close(result)
		response, err = client.DeleteRepo(request)
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

// DeleteRepoRequest is the request struct for api DeleteRepo
type DeleteRepoRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

// DeleteRepoResponse is the response struct for api DeleteRepo
type DeleteRepoResponse struct {
	*responses.BaseResponse
}

// CreateDeleteRepoRequest creates a request to invoke DeleteRepo API
func CreateDeleteRepoRequest() (request *DeleteRepoRequest) {
	request = &DeleteRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "DeleteRepo", "/repos/[RepoNamespace]/[RepoName]", "acr", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteRepoResponse creates a response to parse from DeleteRepo response
func CreateDeleteRepoResponse() (response *DeleteRepoResponse) {
	response = &DeleteRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
