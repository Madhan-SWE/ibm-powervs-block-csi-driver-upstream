// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPSECPolicyEncryptions IP s e c policy encryptions
// Example: ["aes-256-cbc","aes-192-cbc","aes-128-cbc","aes-256-gcm","aes-128-gcm","3des-cbc"]
//
// swagger:model IPSECPolicyEncryptions
type IPSECPolicyEncryptions []string

// Validate validates this IP s e c policy encryptions
func (m IPSECPolicyEncryptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP s e c policy encryptions based on context it is used
func (m IPSECPolicyEncryptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
