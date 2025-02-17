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

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateAWSRegionCreatedCode is the HTTP code returned for type CreateAWSRegionCreated
const CreateAWSRegionCreatedCode int = 201

/*
CreateAWSRegionCreated Resource created

swagger:response createAWSRegionCreated
*/
type CreateAWSRegionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.AwsRegion `json:"body,omitempty"`
}

// NewCreateAWSRegionCreated creates CreateAWSRegionCreated with default headers values
func NewCreateAWSRegionCreated() *CreateAWSRegionCreated {

	return &CreateAWSRegionCreated{}
}

// WithPayload adds the payload to the create a w s region created response
func (o *CreateAWSRegionCreated) WithPayload(payload *models.AwsRegion) *CreateAWSRegionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create a w s region created response
func (o *CreateAWSRegionCreated) SetPayload(payload *models.AwsRegion) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAWSRegionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAWSRegionBadRequestCode is the HTTP code returned for type CreateAWSRegionBadRequest
const CreateAWSRegionBadRequestCode int = 400

/*
CreateAWSRegionBadRequest Bad request

swagger:response createAWSRegionBadRequest
*/
type CreateAWSRegionBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateAWSRegionBadRequest creates CreateAWSRegionBadRequest with default headers values
func NewCreateAWSRegionBadRequest() *CreateAWSRegionBadRequest {

	return &CreateAWSRegionBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create a w s region bad request response
func (o *CreateAWSRegionBadRequest) WithConfigurationVersion(configurationVersion string) *CreateAWSRegionBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create a w s region bad request response
func (o *CreateAWSRegionBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create a w s region bad request response
func (o *CreateAWSRegionBadRequest) WithPayload(payload *models.Error) *CreateAWSRegionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create a w s region bad request response
func (o *CreateAWSRegionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAWSRegionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAWSRegionConflictCode is the HTTP code returned for type CreateAWSRegionConflict
const CreateAWSRegionConflictCode int = 409

/*
CreateAWSRegionConflict The specified resource already exists

swagger:response createAWSRegionConflict
*/
type CreateAWSRegionConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateAWSRegionConflict creates CreateAWSRegionConflict with default headers values
func NewCreateAWSRegionConflict() *CreateAWSRegionConflict {

	return &CreateAWSRegionConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create a w s region conflict response
func (o *CreateAWSRegionConflict) WithConfigurationVersion(configurationVersion string) *CreateAWSRegionConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create a w s region conflict response
func (o *CreateAWSRegionConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create a w s region conflict response
func (o *CreateAWSRegionConflict) WithPayload(payload *models.Error) *CreateAWSRegionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create a w s region conflict response
func (o *CreateAWSRegionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAWSRegionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateAWSRegionDefault General Error

swagger:response createAWSRegionDefault
*/
type CreateAWSRegionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateAWSRegionDefault creates CreateAWSRegionDefault with default headers values
func NewCreateAWSRegionDefault(code int) *CreateAWSRegionDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateAWSRegionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create a w s region default response
func (o *CreateAWSRegionDefault) WithStatusCode(code int) *CreateAWSRegionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create a w s region default response
func (o *CreateAWSRegionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create a w s region default response
func (o *CreateAWSRegionDefault) WithConfigurationVersion(configurationVersion string) *CreateAWSRegionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create a w s region default response
func (o *CreateAWSRegionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create a w s region default response
func (o *CreateAWSRegionDefault) WithPayload(payload *models.Error) *CreateAWSRegionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create a w s region default response
func (o *CreateAWSRegionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAWSRegionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
