package ecs

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

// StartElasticityAssurance invokes the ecs.StartElasticityAssurance API synchronously
func (client *Client) StartElasticityAssurance(request *StartElasticityAssuranceRequest) (response *StartElasticityAssuranceResponse, err error) {
	response = CreateStartElasticityAssuranceResponse()
	err = client.DoAction(request, response)
	return
}

// StartElasticityAssuranceWithChan invokes the ecs.StartElasticityAssurance API asynchronously
func (client *Client) StartElasticityAssuranceWithChan(request *StartElasticityAssuranceRequest) (<-chan *StartElasticityAssuranceResponse, <-chan error) {
	responseChan := make(chan *StartElasticityAssuranceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartElasticityAssurance(request)
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

// StartElasticityAssuranceWithCallback invokes the ecs.StartElasticityAssurance API asynchronously
func (client *Client) StartElasticityAssuranceWithCallback(request *StartElasticityAssuranceRequest, callback func(response *StartElasticityAssuranceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartElasticityAssuranceResponse
		var err error
		defer close(result)
		response, err = client.StartElasticityAssurance(request)
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

// StartElasticityAssuranceRequest is the request struct for api StartElasticityAssurance
type StartElasticityAssuranceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PrivatePoolOptionsId string           `position:"Query" name:"PrivatePoolOptions.Id"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// StartElasticityAssuranceResponse is the response struct for api StartElasticityAssurance
type StartElasticityAssuranceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartElasticityAssuranceRequest creates a request to invoke StartElasticityAssurance API
func CreateStartElasticityAssuranceRequest() (request *StartElasticityAssuranceRequest) {
	request = &StartElasticityAssuranceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "StartElasticityAssurance", "", "")
	request.Method = requests.POST
	return
}

// CreateStartElasticityAssuranceResponse creates a response to parse from StartElasticityAssurance response
func CreateStartElasticityAssuranceResponse() (response *StartElasticityAssuranceResponse) {
	response = &StartElasticityAssuranceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}