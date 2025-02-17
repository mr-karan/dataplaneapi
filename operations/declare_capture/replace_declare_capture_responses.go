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

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceDeclareCaptureOKCode is the HTTP code returned for type ReplaceDeclareCaptureOK
const ReplaceDeclareCaptureOKCode int = 200

/*
ReplaceDeclareCaptureOK Declare Capture replaced

swagger:response replaceDeclareCaptureOK
*/
type ReplaceDeclareCaptureOK struct {

	/*
	  In: Body
	*/
	Payload *models.Capture `json:"body,omitempty"`
}

// NewReplaceDeclareCaptureOK creates ReplaceDeclareCaptureOK with default headers values
func NewReplaceDeclareCaptureOK() *ReplaceDeclareCaptureOK {

	return &ReplaceDeclareCaptureOK{}
}

// WithPayload adds the payload to the replace declare capture o k response
func (o *ReplaceDeclareCaptureOK) WithPayload(payload *models.Capture) *ReplaceDeclareCaptureOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace declare capture o k response
func (o *ReplaceDeclareCaptureOK) SetPayload(payload *models.Capture) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDeclareCaptureOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceDeclareCaptureAcceptedCode is the HTTP code returned for type ReplaceDeclareCaptureAccepted
const ReplaceDeclareCaptureAcceptedCode int = 202

/*
ReplaceDeclareCaptureAccepted Configuration change accepted and reload requested

swagger:response replaceDeclareCaptureAccepted
*/
type ReplaceDeclareCaptureAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Capture `json:"body,omitempty"`
}

// NewReplaceDeclareCaptureAccepted creates ReplaceDeclareCaptureAccepted with default headers values
func NewReplaceDeclareCaptureAccepted() *ReplaceDeclareCaptureAccepted {

	return &ReplaceDeclareCaptureAccepted{}
}

// WithReloadID adds the reloadId to the replace declare capture accepted response
func (o *ReplaceDeclareCaptureAccepted) WithReloadID(reloadID string) *ReplaceDeclareCaptureAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace declare capture accepted response
func (o *ReplaceDeclareCaptureAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace declare capture accepted response
func (o *ReplaceDeclareCaptureAccepted) WithPayload(payload *models.Capture) *ReplaceDeclareCaptureAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace declare capture accepted response
func (o *ReplaceDeclareCaptureAccepted) SetPayload(payload *models.Capture) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDeclareCaptureAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceDeclareCaptureBadRequestCode is the HTTP code returned for type ReplaceDeclareCaptureBadRequest
const ReplaceDeclareCaptureBadRequestCode int = 400

/*
ReplaceDeclareCaptureBadRequest Bad request

swagger:response replaceDeclareCaptureBadRequest
*/
type ReplaceDeclareCaptureBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceDeclareCaptureBadRequest creates ReplaceDeclareCaptureBadRequest with default headers values
func NewReplaceDeclareCaptureBadRequest() *ReplaceDeclareCaptureBadRequest {

	return &ReplaceDeclareCaptureBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace declare capture bad request response
func (o *ReplaceDeclareCaptureBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceDeclareCaptureBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace declare capture bad request response
func (o *ReplaceDeclareCaptureBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace declare capture bad request response
func (o *ReplaceDeclareCaptureBadRequest) WithPayload(payload *models.Error) *ReplaceDeclareCaptureBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace declare capture bad request response
func (o *ReplaceDeclareCaptureBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDeclareCaptureBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceDeclareCaptureNotFoundCode is the HTTP code returned for type ReplaceDeclareCaptureNotFound
const ReplaceDeclareCaptureNotFoundCode int = 404

/*
ReplaceDeclareCaptureNotFound The specified resource was not found

swagger:response replaceDeclareCaptureNotFound
*/
type ReplaceDeclareCaptureNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceDeclareCaptureNotFound creates ReplaceDeclareCaptureNotFound with default headers values
func NewReplaceDeclareCaptureNotFound() *ReplaceDeclareCaptureNotFound {

	return &ReplaceDeclareCaptureNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace declare capture not found response
func (o *ReplaceDeclareCaptureNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceDeclareCaptureNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace declare capture not found response
func (o *ReplaceDeclareCaptureNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace declare capture not found response
func (o *ReplaceDeclareCaptureNotFound) WithPayload(payload *models.Error) *ReplaceDeclareCaptureNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace declare capture not found response
func (o *ReplaceDeclareCaptureNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDeclareCaptureNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
ReplaceDeclareCaptureDefault General Error

swagger:response replaceDeclareCaptureDefault
*/
type ReplaceDeclareCaptureDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceDeclareCaptureDefault creates ReplaceDeclareCaptureDefault with default headers values
func NewReplaceDeclareCaptureDefault(code int) *ReplaceDeclareCaptureDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceDeclareCaptureDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace declare capture default response
func (o *ReplaceDeclareCaptureDefault) WithStatusCode(code int) *ReplaceDeclareCaptureDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace declare capture default response
func (o *ReplaceDeclareCaptureDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace declare capture default response
func (o *ReplaceDeclareCaptureDefault) WithConfigurationVersion(configurationVersion string) *ReplaceDeclareCaptureDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace declare capture default response
func (o *ReplaceDeclareCaptureDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace declare capture default response
func (o *ReplaceDeclareCaptureDefault) WithPayload(payload *models.Error) *ReplaceDeclareCaptureDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace declare capture default response
func (o *ReplaceDeclareCaptureDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDeclareCaptureDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
