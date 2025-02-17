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

package mailer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetMailerEntriesOKCode is the HTTP code returned for type GetMailerEntriesOK
const GetMailerEntriesOKCode int = 200

/*
GetMailerEntriesOK Successful operation

swagger:response getMailerEntriesOK
*/
type GetMailerEntriesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetMailerEntriesOKBody `json:"body,omitempty"`
}

// NewGetMailerEntriesOK creates GetMailerEntriesOK with default headers values
func NewGetMailerEntriesOK() *GetMailerEntriesOK {

	return &GetMailerEntriesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get mailer entries o k response
func (o *GetMailerEntriesOK) WithConfigurationVersion(configurationVersion string) *GetMailerEntriesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get mailer entries o k response
func (o *GetMailerEntriesOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get mailer entries o k response
func (o *GetMailerEntriesOK) WithPayload(payload *GetMailerEntriesOKBody) *GetMailerEntriesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get mailer entries o k response
func (o *GetMailerEntriesOK) SetPayload(payload *GetMailerEntriesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMailerEntriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetMailerEntriesDefault General Error

swagger:response getMailerEntriesDefault
*/
type GetMailerEntriesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMailerEntriesDefault creates GetMailerEntriesDefault with default headers values
func NewGetMailerEntriesDefault(code int) *GetMailerEntriesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMailerEntriesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get mailer entries default response
func (o *GetMailerEntriesDefault) WithStatusCode(code int) *GetMailerEntriesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get mailer entries default response
func (o *GetMailerEntriesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get mailer entries default response
func (o *GetMailerEntriesDefault) WithConfigurationVersion(configurationVersion string) *GetMailerEntriesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get mailer entries default response
func (o *GetMailerEntriesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get mailer entries default response
func (o *GetMailerEntriesDefault) WithPayload(payload *models.Error) *GetMailerEntriesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get mailer entries default response
func (o *GetMailerEntriesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMailerEntriesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
