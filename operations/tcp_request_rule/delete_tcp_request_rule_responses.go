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

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteTCPRequestRuleAcceptedCode is the HTTP code returned for type DeleteTCPRequestRuleAccepted
const DeleteTCPRequestRuleAcceptedCode int = 202

/*
DeleteTCPRequestRuleAccepted Configuration change accepted and reload requested

swagger:response deleteTcpRequestRuleAccepted
*/
type DeleteTCPRequestRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteTCPRequestRuleAccepted creates DeleteTCPRequestRuleAccepted with default headers values
func NewDeleteTCPRequestRuleAccepted() *DeleteTCPRequestRuleAccepted {

	return &DeleteTCPRequestRuleAccepted{}
}

// WithReloadID adds the reloadId to the delete Tcp request rule accepted response
func (o *DeleteTCPRequestRuleAccepted) WithReloadID(reloadID string) *DeleteTCPRequestRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete Tcp request rule accepted response
func (o *DeleteTCPRequestRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteTCPRequestRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteTCPRequestRuleNoContentCode is the HTTP code returned for type DeleteTCPRequestRuleNoContent
const DeleteTCPRequestRuleNoContentCode int = 204

/*
DeleteTCPRequestRuleNoContent TCP Request Rule deleted

swagger:response deleteTcpRequestRuleNoContent
*/
type DeleteTCPRequestRuleNoContent struct {
}

// NewDeleteTCPRequestRuleNoContent creates DeleteTCPRequestRuleNoContent with default headers values
func NewDeleteTCPRequestRuleNoContent() *DeleteTCPRequestRuleNoContent {

	return &DeleteTCPRequestRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteTCPRequestRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTCPRequestRuleNotFoundCode is the HTTP code returned for type DeleteTCPRequestRuleNotFound
const DeleteTCPRequestRuleNotFoundCode int = 404

/*
DeleteTCPRequestRuleNotFound The specified resource was not found

swagger:response deleteTcpRequestRuleNotFound
*/
type DeleteTCPRequestRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTCPRequestRuleNotFound creates DeleteTCPRequestRuleNotFound with default headers values
func NewDeleteTCPRequestRuleNotFound() *DeleteTCPRequestRuleNotFound {

	return &DeleteTCPRequestRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete Tcp request rule not found response
func (o *DeleteTCPRequestRuleNotFound) WithConfigurationVersion(configurationVersion string) *DeleteTCPRequestRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete Tcp request rule not found response
func (o *DeleteTCPRequestRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete Tcp request rule not found response
func (o *DeleteTCPRequestRuleNotFound) WithPayload(payload *models.Error) *DeleteTCPRequestRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Tcp request rule not found response
func (o *DeleteTCPRequestRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTCPRequestRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
DeleteTCPRequestRuleDefault General Error

swagger:response deleteTcpRequestRuleDefault
*/
type DeleteTCPRequestRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTCPRequestRuleDefault creates DeleteTCPRequestRuleDefault with default headers values
func NewDeleteTCPRequestRuleDefault(code int) *DeleteTCPRequestRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTCPRequestRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete TCP request rule default response
func (o *DeleteTCPRequestRuleDefault) WithStatusCode(code int) *DeleteTCPRequestRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete TCP request rule default response
func (o *DeleteTCPRequestRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete TCP request rule default response
func (o *DeleteTCPRequestRuleDefault) WithConfigurationVersion(configurationVersion string) *DeleteTCPRequestRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete TCP request rule default response
func (o *DeleteTCPRequestRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete TCP request rule default response
func (o *DeleteTCPRequestRuleDefault) WithPayload(payload *models.Error) *DeleteTCPRequestRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete TCP request rule default response
func (o *DeleteTCPRequestRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTCPRequestRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
