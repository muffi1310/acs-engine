package billing

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DownloadURL is a secure URL that can be used to download a PDF invoice until
// the URL expires.
type DownloadURL struct {
	ExpiryTime *date.Time `json:"expiryTime,omitempty"`
	URL        *string    `json:"url,omitempty"`
}

// ErrorDetails is the details of the error.
type ErrorDetails struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// ErrorResponse is error response indicates that the service is not able to
// process the incoming request. The reason is provided in the error message.
type ErrorResponse struct {
	Error *ErrorDetails `json:"error,omitempty"`
}

// Invoice is an invoice resource can be used download a PDF version of an
// invoice.
type Invoice struct {
	autorest.Response  `json:"-"`
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	Type               *string `json:"type,omitempty"`
	*InvoiceProperties `json:"properties,omitempty"`
}

// InvoiceProperties is the properties of the invoice.
type InvoiceProperties struct {
	DownloadURL            *DownloadURL `json:"downloadUrl,omitempty"`
	InvoicePeriodStartDate *date.Date   `json:"invoicePeriodStartDate,omitempty"`
	InvoicePeriodEndDate   *date.Date   `json:"invoicePeriodEndDate,omitempty"`
	BillingPeriodIds       *[]string    `json:"billingPeriodIds,omitempty"`
}

// InvoicesListResult is result of listing invoices. It contains a list of
// available invoices in reverse chronological order.
type InvoicesListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Invoice `json:"value,omitempty"`
	NextLink          *string    `json:"nextLink,omitempty"`
}

// InvoicesListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client InvoicesListResult) InvoicesListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Operation is a Billing REST API operation.
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is the object that represents the operation.
type OperationDisplay struct {
	Provider  *string `json:"provider,omitempty"`
	Resource  *string `json:"resource,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult is result listing billing operations. It contains a list
// of operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Period is a billing period resource.
type Period struct {
	autorest.Response `json:"-"`
	ID                *string `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
	*PeriodProperties `json:"properties,omitempty"`
}

// PeriodProperties is the properties of the billing period.
type PeriodProperties struct {
	BillingPeriodStartDate *date.Date `json:"billingPeriodStartDate,omitempty"`
	BillingPeriodEndDate   *date.Date `json:"billingPeriodEndDate,omitempty"`
	InvoiceIds             *[]string  `json:"invoiceIds,omitempty"`
}

// PeriodsListResult is result of listing billing periods. It contains a list
// of available billing periods in reverse chronological order.
type PeriodsListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Period `json:"value,omitempty"`
	NextLink          *string   `json:"nextLink,omitempty"`
}

// PeriodsListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client PeriodsListResult) PeriodsListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Resource is the Resource model definition.
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}
