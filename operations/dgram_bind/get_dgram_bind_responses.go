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

package dgram_bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetDgramBindOKCode is the HTTP code returned for type GetDgramBindOK
const GetDgramBindOKCode int = 200

/*
GetDgramBindOK Successful operation

swagger:response getDgramBindOK
*/
type GetDgramBindOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetDgramBindOKBody `json:"body,omitempty"`
}

// NewGetDgramBindOK creates GetDgramBindOK with default headers values
func NewGetDgramBindOK() *GetDgramBindOK {

	return &GetDgramBindOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get dgram bind o k response
func (o *GetDgramBindOK) WithConfigurationVersion(configurationVersion string) *GetDgramBindOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get dgram bind o k response
func (o *GetDgramBindOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get dgram bind o k response
func (o *GetDgramBindOK) WithPayload(payload *GetDgramBindOKBody) *GetDgramBindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dgram bind o k response
func (o *GetDgramBindOK) SetPayload(payload *GetDgramBindOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDgramBindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetDgramBindNotFoundCode is the HTTP code returned for type GetDgramBindNotFound
const GetDgramBindNotFoundCode int = 404

/*
GetDgramBindNotFound The specified resource already exists

swagger:response getDgramBindNotFound
*/
type GetDgramBindNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDgramBindNotFound creates GetDgramBindNotFound with default headers values
func NewGetDgramBindNotFound() *GetDgramBindNotFound {

	return &GetDgramBindNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get dgram bind not found response
func (o *GetDgramBindNotFound) WithConfigurationVersion(configurationVersion string) *GetDgramBindNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get dgram bind not found response
func (o *GetDgramBindNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get dgram bind not found response
func (o *GetDgramBindNotFound) WithPayload(payload *models.Error) *GetDgramBindNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dgram bind not found response
func (o *GetDgramBindNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDgramBindNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetDgramBindDefault General Error

swagger:response getDgramBindDefault
*/
type GetDgramBindDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDgramBindDefault creates GetDgramBindDefault with default headers values
func NewGetDgramBindDefault(code int) *GetDgramBindDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDgramBindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get dgram bind default response
func (o *GetDgramBindDefault) WithStatusCode(code int) *GetDgramBindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get dgram bind default response
func (o *GetDgramBindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get dgram bind default response
func (o *GetDgramBindDefault) WithConfigurationVersion(configurationVersion string) *GetDgramBindDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get dgram bind default response
func (o *GetDgramBindDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get dgram bind default response
func (o *GetDgramBindDefault) WithPayload(payload *models.Error) *GetDgramBindDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dgram bind default response
func (o *GetDgramBindDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDgramBindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
