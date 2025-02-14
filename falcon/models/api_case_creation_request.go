// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APICaseCreationRequest api case creation request
//
// swagger:model api.CaseCreationRequest
type APICaseCreationRequest struct {

	// body
	// Required: true
	Body *string `json:"body"`

	// detections
	// Required: true
	Detections []*MessagesDetection `json:"detections"`

	// incidents
	// Required: true
	Incidents []*MessagesIncident `json:"incidents"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type *string `json:"type"`

	// user uuid
	// Required: true
	UserUUID *string `json:"user_uuid"`
}

// Validate validates this api case creation request
func (m *APICaseCreationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICaseCreationRequest) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *APICaseCreationRequest) validateDetections(formats strfmt.Registry) error {

	if err := validate.Required("detections", "body", m.Detections); err != nil {
		return err
	}

	for i := 0; i < len(m.Detections); i++ {
		if swag.IsZero(m.Detections[i]) { // not required
			continue
		}

		if m.Detections[i] != nil {
			if err := m.Detections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("detections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("detections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APICaseCreationRequest) validateIncidents(formats strfmt.Registry) error {

	if err := validate.Required("incidents", "body", m.Incidents); err != nil {
		return err
	}

	for i := 0; i < len(m.Incidents); i++ {
		if swag.IsZero(m.Incidents[i]) { // not required
			continue
		}

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APICaseCreationRequest) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *APICaseCreationRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *APICaseCreationRequest) validateUserUUID(formats strfmt.Registry) error {

	if err := validate.Required("user_uuid", "body", m.UserUUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api case creation request based on the context it is used
func (m *APICaseCreationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncidents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICaseCreationRequest) contextValidateDetections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Detections); i++ {

		if m.Detections[i] != nil {
			if err := m.Detections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("detections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("detections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APICaseCreationRequest) contextValidateIncidents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Incidents); i++ {

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APICaseCreationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICaseCreationRequest) UnmarshalBinary(b []byte) error {
	var res APICaseCreationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
