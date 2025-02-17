// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package http_after_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetHTTPAfterResponseRulesOKCode is the HTTP code returned for type GetHTTPAfterResponseRulesOK
const GetHTTPAfterResponseRulesOKCode int = 200

/*
GetHTTPAfterResponseRulesOK Successful operation

swagger:response getHttpAfterResponseRulesOK
*/
type GetHTTPAfterResponseRulesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetHTTPAfterResponseRulesOKBody `json:"body,omitempty"`
}

// NewGetHTTPAfterResponseRulesOK creates GetHTTPAfterResponseRulesOK with default headers values
func NewGetHTTPAfterResponseRulesOK() *GetHTTPAfterResponseRulesOK {

	return &GetHTTPAfterResponseRulesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http after response rules o k response
func (o *GetHTTPAfterResponseRulesOK) WithConfigurationVersion(configurationVersion string) *GetHTTPAfterResponseRulesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http after response rules o k response
func (o *GetHTTPAfterResponseRulesOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http after response rules o k response
func (o *GetHTTPAfterResponseRulesOK) WithPayload(payload *GetHTTPAfterResponseRulesOKBody) *GetHTTPAfterResponseRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http after response rules o k response
func (o *GetHTTPAfterResponseRulesOK) SetPayload(payload *GetHTTPAfterResponseRulesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPAfterResponseRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetHTTPAfterResponseRulesDefault General Error

swagger:response getHttpAfterResponseRulesDefault
*/
type GetHTTPAfterResponseRulesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPAfterResponseRulesDefault creates GetHTTPAfterResponseRulesDefault with default headers values
func NewGetHTTPAfterResponseRulesDefault(code int) *GetHTTPAfterResponseRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHTTPAfterResponseRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get HTTP after response rules default response
func (o *GetHTTPAfterResponseRulesDefault) WithStatusCode(code int) *GetHTTPAfterResponseRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get HTTP after response rules default response
func (o *GetHTTPAfterResponseRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get HTTP after response rules default response
func (o *GetHTTPAfterResponseRulesDefault) WithConfigurationVersion(configurationVersion string) *GetHTTPAfterResponseRulesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get HTTP after response rules default response
func (o *GetHTTPAfterResponseRulesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get HTTP after response rules default response
func (o *GetHTTPAfterResponseRulesDefault) WithPayload(payload *models.Error) *GetHTTPAfterResponseRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get HTTP after response rules default response
func (o *GetHTTPAfterResponseRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPAfterResponseRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
