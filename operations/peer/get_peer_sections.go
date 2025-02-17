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

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v4/models"
)

// GetPeerSectionsHandlerFunc turns a function with the right signature into a get peer sections handler
type GetPeerSectionsHandlerFunc func(GetPeerSectionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPeerSectionsHandlerFunc) Handle(params GetPeerSectionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetPeerSectionsHandler interface for that can handle valid get peer sections params
type GetPeerSectionsHandler interface {
	Handle(GetPeerSectionsParams, interface{}) middleware.Responder
}

// NewGetPeerSections creates a new http.Handler for the get peer sections operation
func NewGetPeerSections(ctx *middleware.Context, handler GetPeerSectionsHandler) *GetPeerSections {
	return &GetPeerSections{Context: ctx, Handler: handler}
}

/*
	GetPeerSections swagger:route GET /services/haproxy/configuration/peer_section Peer getPeerSections

Return an array of peer_section

Returns an array of all configured peer_section.
*/
type GetPeerSections struct {
	Context *middleware.Context
	Handler GetPeerSectionsHandler
}

func (o *GetPeerSections) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPeerSectionsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetPeerSectionsOKBody get peer sections o k body
//
// swagger:model GetPeerSectionsOKBody
type GetPeerSectionsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.PeerSections `json:"data"`
}

// Validate validates this get peer sections o k body
func (o *GetPeerSectionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerSectionsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPeerSectionsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getPeerSectionsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getPeerSectionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get peer sections o k body based on the context it is used
func (o *GetPeerSectionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerSectionsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getPeerSectionsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getPeerSectionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPeerSectionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPeerSectionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPeerSectionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
