package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ApplicationGatewaysClient is the composite Swagger for Network Client
type ApplicationGatewaysClient struct {
	ManagementClient
}

// NewApplicationGatewaysClient creates an instance of the
// ApplicationGatewaysClient client.
func NewApplicationGatewaysClient(subscriptionID string) ApplicationGatewaysClient {
	return NewApplicationGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGatewaysClientWithBaseURI creates an instance of the
// ApplicationGatewaysClient client.
func NewApplicationGatewaysClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewaysClient {
	return ApplicationGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// BackendHealth gets the backend health of the specified application gateway
// in a resource group. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the application gateway. expand is expands BackendAddressPool
// and BackendHttpSettings referenced in backend health.
func (client ApplicationGatewaysClient) BackendHealth(resourceGroupName string, applicationGatewayName string, expand string, cancel <-chan struct{}) (<-chan ApplicationGatewayBackendHealth, <-chan error) {
	resultChan := make(chan ApplicationGatewayBackendHealth, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result ApplicationGatewayBackendHealth
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.BackendHealthPreparer(resourceGroupName, applicationGatewayName, expand, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "BackendHealth", nil, "Failure preparing request")
			return
		}

		resp, err := client.BackendHealthSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "BackendHealth", resp, "Failure sending request")
			return
		}

		result, err = client.BackendHealthResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "BackendHealth", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// BackendHealthPreparer prepares the BackendHealth request.
func (client ApplicationGatewaysClient) BackendHealthPreparer(resourceGroupName string, applicationGatewayName string, expand string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/backendhealth", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// BackendHealthSender sends the BackendHealth request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) BackendHealthSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// BackendHealthResponder handles the response to the BackendHealth request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) BackendHealthResponder(resp *http.Response) (result ApplicationGatewayBackendHealth, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates the specified application gateway. This
// method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the application gateway. parameters is parameters supplied to
// the create or update application gateway operation.
func (client ApplicationGatewaysClient) CreateOrUpdate(resourceGroupName string, applicationGatewayName string, parameters ApplicationGateway, cancel <-chan struct{}) (<-chan ApplicationGateway, <-chan error) {
	resultChan := make(chan ApplicationGateway, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplicationGatewayPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration.Enabled", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration.RuleSetType", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration.RuleSetVersion", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "network.ApplicationGatewaysClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result ApplicationGateway
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, applicationGatewayName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplicationGatewaysClient) CreateOrUpdatePreparer(resourceGroupName string, applicationGatewayName string, parameters ApplicationGateway, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result ApplicationGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified application gateway. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the application gateway.
func (client ApplicationGatewaysClient) Delete(resourceGroupName string, applicationGatewayName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, applicationGatewayName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client ApplicationGatewaysClient) DeletePreparer(resourceGroupName string, applicationGatewayName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified application gateway.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the application gateway.
func (client ApplicationGatewaysClient) Get(resourceGroupName string, applicationGatewayName string) (result ApplicationGateway, err error) {
	req, err := client.GetPreparer(resourceGroupName, applicationGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationGatewaysClient) GetPreparer(resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) GetResponder(resp *http.Response) (result ApplicationGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all application gateways in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client ApplicationGatewaysClient) List(resourceGroupName string) (result ApplicationGatewayListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationGatewaysClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) ListResponder(resp *http.Response) (result ApplicationGatewayListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ApplicationGatewaysClient) ListNextResults(lastResults ApplicationGatewayListResult) (result ApplicationGatewayListResult, err error) {
	req, err := lastResults.ApplicationGatewayListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListAll gets all the application gateways in a subscription.
func (client ApplicationGatewaysClient) ListAll() (result ApplicationGatewayListResult, err error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", resp, "Failure sending request")
		return
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client ApplicationGatewaysClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) ListAllResponder(resp *http.Response) (result ApplicationGatewayListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client ApplicationGatewaysClient) ListAllNextResults(lastResults ApplicationGatewayListResult) (result ApplicationGatewayListResult, err error) {
	req, err := lastResults.ApplicationGatewayListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", resp, "Failure sending next results request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", resp, "Failure responding to next results request")
	}

	return
}

// ListAvailableWafRuleSets lists all available web application firewall rule
// sets.
func (client ApplicationGatewaysClient) ListAvailableWafRuleSets() (result ApplicationGatewayAvailableWafRuleSetsResult, err error) {
	req, err := client.ListAvailableWafRuleSetsPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAvailableWafRuleSets", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAvailableWafRuleSetsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAvailableWafRuleSets", resp, "Failure sending request")
		return
	}

	result, err = client.ListAvailableWafRuleSetsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAvailableWafRuleSets", resp, "Failure responding to request")
	}

	return
}

// ListAvailableWafRuleSetsPreparer prepares the ListAvailableWafRuleSets request.
func (client ApplicationGatewaysClient) ListAvailableWafRuleSetsPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAvailableWafRuleSetsSender sends the ListAvailableWafRuleSets request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) ListAvailableWafRuleSetsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAvailableWafRuleSetsResponder handles the response to the ListAvailableWafRuleSets request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) ListAvailableWafRuleSetsResponder(resp *http.Response) (result ApplicationGatewayAvailableWafRuleSetsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start starts the specified application gateway. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the application gateway.
func (client ApplicationGatewaysClient) Start(resourceGroupName string, applicationGatewayName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.StartPreparer(resourceGroupName, applicationGatewayName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Start", nil, "Failure preparing request")
			return
		}

		resp, err := client.StartSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Start", resp, "Failure sending request")
			return
		}

		result, err = client.StartResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Start", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// StartPreparer prepares the Start request.
func (client ApplicationGatewaysClient) StartPreparer(resourceGroupName string, applicationGatewayName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) StartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops the specified application gateway in a resource group. This
// method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the application gateway.
func (client ApplicationGatewaysClient) Stop(resourceGroupName string, applicationGatewayName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.StopPreparer(resourceGroupName, applicationGatewayName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Stop", nil, "Failure preparing request")
			return
		}

		resp, err := client.StopSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Stop", resp, "Failure sending request")
			return
		}

		result, err = client.StopResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Stop", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// StopPreparer prepares the Stop request.
func (client ApplicationGatewaysClient) StopPreparer(resourceGroupName string, applicationGatewayName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) StopSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
