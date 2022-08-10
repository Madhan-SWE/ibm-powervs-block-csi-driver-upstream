// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPSECPolicyAuthentications IP s e c policy authentications
// Example: ["hmac-sha-256-128","hmac-sha1-96","none"]
//
// swagger:model IPSECPolicyAuthentications
type IPSECPolicyAuthentications []string

// Validate validates this IP s e c policy authentications
func (m IPSECPolicyAuthentications) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP s e c policy authentications based on context it is used
func (m IPSECPolicyAuthentications) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
