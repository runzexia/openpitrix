// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeAppVersionsResponse openpitrix describe app versions response
// swagger:model openpitrixDescribeAppVersionsResponse
type OpenpitrixDescribeAppVersionsResponse struct {

	// app version set
	AppVersionSet OpenpitrixDescribeAppVersionsResponseAppVersionSet `json:"app_version_set"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix describe app versions response
func (m *OpenpitrixDescribeAppVersionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeAppVersionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeAppVersionsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeAppVersionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
