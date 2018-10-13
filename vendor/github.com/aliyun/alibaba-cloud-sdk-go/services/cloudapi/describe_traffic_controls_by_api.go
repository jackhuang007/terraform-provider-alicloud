package cloudapi

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

// DescribeTrafficControlsByApi invokes the cloudapi.DescribeTrafficControlsByApi API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describetrafficcontrolsbyapi.html
func (client *Client) DescribeTrafficControlsByApi(request *DescribeTrafficControlsByApiRequest) (response *DescribeTrafficControlsByApiResponse, err error) {
	response = CreateDescribeTrafficControlsByApiResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTrafficControlsByApiWithChan invokes the cloudapi.DescribeTrafficControlsByApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describetrafficcontrolsbyapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTrafficControlsByApiWithChan(request *DescribeTrafficControlsByApiRequest) (<-chan *DescribeTrafficControlsByApiResponse, <-chan error) {
	responseChan := make(chan *DescribeTrafficControlsByApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTrafficControlsByApi(request)
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

// DescribeTrafficControlsByApiWithCallback invokes the cloudapi.DescribeTrafficControlsByApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describetrafficcontrolsbyapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTrafficControlsByApiWithCallback(request *DescribeTrafficControlsByApiRequest, callback func(response *DescribeTrafficControlsByApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTrafficControlsByApiResponse
		var err error
		defer close(result)
		response, err = client.DescribeTrafficControlsByApi(request)
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

// DescribeTrafficControlsByApiRequest is the request struct for api DescribeTrafficControlsByApi
type DescribeTrafficControlsByApiRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// DescribeTrafficControlsByApiResponse is the response struct for api DescribeTrafficControlsByApi
type DescribeTrafficControlsByApiResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	TrafficControlItems TrafficControlItems `json:"TrafficControlItems" xml:"TrafficControlItems"`
}

// CreateDescribeTrafficControlsByApiRequest creates a request to invoke DescribeTrafficControlsByApi API
func CreateDescribeTrafficControlsByApiRequest() (request *DescribeTrafficControlsByApiRequest) {
	request = &DescribeTrafficControlsByApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeTrafficControlsByApi", "apigateway", "openAPI")
	return
}

// CreateDescribeTrafficControlsByApiResponse creates a response to parse from DescribeTrafficControlsByApi response
func CreateDescribeTrafficControlsByApiResponse() (response *DescribeTrafficControlsByApiResponse) {
	response = &DescribeTrafficControlsByApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}