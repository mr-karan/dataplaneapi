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

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetTCPResponseRuleOKCode is the HTTP code returned for type GetTCPResponseRuleOK
const GetTCPResponseRuleOKCode int = 200

/*
GetTCPResponseRuleOK Successful operation

swagger:response getTcpResponseRuleOK
*/
type GetTCPResponseRuleOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetTCPResponseRuleOKBody `json:"body,omitempty"`
}

// NewGetTCPResponseRuleOK creates GetTCPResponseRuleOK with default headers values
func NewGetTCPResponseRuleOK() *GetTCPResponseRuleOK {

	return &GetTCPResponseRuleOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Tcp response rule o k response
func (o *GetTCPResponseRuleOK) WithConfigurationVersion(configurationVersion string) *GetTCPResponseRuleOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Tcp response rule o k response
func (o *GetTCPResponseRuleOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Tcp response rule o k response
func (o *GetTCPResponseRuleOK) WithPayload(payload *GetTCPResponseRuleOKBody) *GetTCPResponseRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Tcp response rule o k response
func (o *GetTCPResponseRuleOK) SetPayload(payload *GetTCPResponseRuleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPResponseRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetTCPResponseRuleNotFoundCode is the HTTP code returned for type GetTCPResponseRuleNotFound
const GetTCPResponseRuleNotFoundCode int = 404

/*
GetTCPResponseRuleNotFound The specified resource was not found

swagger:response getTcpResponseRuleNotFound
*/
type GetTCPResponseRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTCPResponseRuleNotFound creates GetTCPResponseRuleNotFound with default headers values
func NewGetTCPResponseRuleNotFound() *GetTCPResponseRuleNotFound {

	return &GetTCPResponseRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get Tcp response rule not found response
func (o *GetTCPResponseRuleNotFound) WithConfigurationVersion(configurationVersion string) *GetTCPResponseRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Tcp response rule not found response
func (o *GetTCPResponseRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Tcp response rule not found response
func (o *GetTCPResponseRuleNotFound) WithPayload(payload *models.Error) *GetTCPResponseRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Tcp response rule not found response
func (o *GetTCPResponseRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPResponseRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetTCPResponseRuleDefault General Error

swagger:response getTcpResponseRuleDefault
*/
type GetTCPResponseRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTCPResponseRuleDefault creates GetTCPResponseRuleDefault with default headers values
func NewGetTCPResponseRuleDefault(code int) *GetTCPResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTCPResponseRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get TCP response rule default response
func (o *GetTCPResponseRuleDefault) WithStatusCode(code int) *GetTCPResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get TCP response rule default response
func (o *GetTCPResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get TCP response rule default response
func (o *GetTCPResponseRuleDefault) WithConfigurationVersion(configurationVersion string) *GetTCPResponseRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get TCP response rule default response
func (o *GetTCPResponseRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get TCP response rule default response
func (o *GetTCPResponseRuleDefault) WithPayload(payload *models.Error) *GetTCPResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get TCP response rule default response
func (o *GetTCPResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
