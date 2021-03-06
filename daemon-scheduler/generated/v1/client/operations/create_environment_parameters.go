// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/blox/blox/daemon-scheduler/generated/v1/models"
)

// NewCreateEnvironmentParams creates a new CreateEnvironmentParams object
// with the default values initialized.
func NewCreateEnvironmentParams() *CreateEnvironmentParams {
	var ()
	return &CreateEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEnvironmentParamsWithTimeout creates a new CreateEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateEnvironmentParamsWithTimeout(timeout time.Duration) *CreateEnvironmentParams {
	var ()
	return &CreateEnvironmentParams{

		timeout: timeout,
	}
}

// NewCreateEnvironmentParamsWithContext creates a new CreateEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateEnvironmentParamsWithContext(ctx context.Context) *CreateEnvironmentParams {
	var ()
	return &CreateEnvironmentParams{

		Context: ctx,
	}
}

/*CreateEnvironmentParams contains all the parameters to send to the API endpoint
for the create environment operation typically these are written to a http.Request
*/
type CreateEnvironmentParams struct {

	/*Body*/
	Body *models.CreateEnvironmentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create environment params
func (o *CreateEnvironmentParams) WithTimeout(timeout time.Duration) *CreateEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create environment params
func (o *CreateEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create environment params
func (o *CreateEnvironmentParams) WithContext(ctx context.Context) *CreateEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create environment params
func (o *CreateEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the create environment params
func (o *CreateEnvironmentParams) WithBody(body *models.CreateEnvironmentRequest) *CreateEnvironmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create environment params
func (o *CreateEnvironmentParams) SetBody(body *models.CreateEnvironmentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.CreateEnvironmentRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
