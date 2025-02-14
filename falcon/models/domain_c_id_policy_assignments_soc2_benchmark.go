// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainCIDPolicyAssignmentsSoc2Benchmark domain c ID policy assignments soc2 benchmark
//
// swagger:model domain.CIDPolicyAssignments.soc2_benchmark
type DomainCIDPolicyAssignmentsSoc2Benchmark struct {

	// benchmark short
	// Required: true
	BenchmarkShort *string `json:"benchmark_short"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// recommendation number
	// Required: true
	RecommendationNumber *string `json:"recommendation_number"`
}

// Validate validates this domain c ID policy assignments soc2 benchmark
func (m *DomainCIDPolicyAssignmentsSoc2Benchmark) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBenchmarkShort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecommendationNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainCIDPolicyAssignmentsSoc2Benchmark) validateBenchmarkShort(formats strfmt.Registry) error {

	if err := validate.Required("benchmark_short", "body", m.BenchmarkShort); err != nil {
		return err
	}

	return nil
}

func (m *DomainCIDPolicyAssignmentsSoc2Benchmark) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainCIDPolicyAssignmentsSoc2Benchmark) validateRecommendationNumber(formats strfmt.Registry) error {

	if err := validate.Required("recommendation_number", "body", m.RecommendationNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain c ID policy assignments soc2 benchmark based on context it is used
func (m *DomainCIDPolicyAssignmentsSoc2Benchmark) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainCIDPolicyAssignmentsSoc2Benchmark) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCIDPolicyAssignmentsSoc2Benchmark) UnmarshalBinary(b []byte) error {
	var res DomainCIDPolicyAssignmentsSoc2Benchmark
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
