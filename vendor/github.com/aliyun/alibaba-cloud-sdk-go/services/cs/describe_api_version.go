package cs

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

// DescribeApiVersion invokes the cs.DescribeApiVersion API synchronously
// api document: https://help.aliyun.com/api/cs/describeapiversion.html
func (client *Client) DescribeApiVersion(request *DescribeApiVersionRequest) (response *DescribeApiVersionResponse, err error) {
	response = CreateDescribeApiVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiVersionWithChan invokes the cs.DescribeApiVersion API asynchronously
// api document: https://help.aliyun.com/api/cs/describeapiversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiVersionWithChan(request *DescribeApiVersionRequest) (<-chan *DescribeApiVersionResponse, <-chan error) {
	responseChan := make(chan *DescribeApiVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiVersion(request)
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

// DescribeApiVersionWithCallback invokes the cs.DescribeApiVersion API asynchronously
// api document: https://help.aliyun.com/api/cs/describeapiversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiVersionWithCallback(request *DescribeApiVersionRequest, callback func(response *DescribeApiVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiVersionResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiVersion(request)
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

// DescribeApiVersionRequest is the request struct for api DescribeApiVersion
type DescribeApiVersionRequest struct {
	*requests.RoaRequest
}

// DescribeApiVersionResponse is the response struct for api DescribeApiVersion
type DescribeApiVersionResponse struct {
	*responses.BaseResponse
}

// CreateDescribeApiVersionRequest creates a request to invoke DescribeApiVersion API
func CreateDescribeApiVersionRequest() (request *DescribeApiVersionRequest) {
	request = &DescribeApiVersionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeApiVersion", "/version", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeApiVersionResponse creates a response to parse from DescribeApiVersion response
func CreateDescribeApiVersionResponse() (response *DescribeApiVersionResponse) {
	response = &DescribeApiVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}