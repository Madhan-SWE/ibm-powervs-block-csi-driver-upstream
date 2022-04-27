// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsGetReader is a Reader for the PcloudCloudconnectionsGet structure.
type PcloudCloudconnectionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudconnectionsGetOK creates a PcloudCloudconnectionsGetOK with default headers values
func NewPcloudCloudconnectionsGetOK() *PcloudCloudconnectionsGetOK {
	return &PcloudCloudconnectionsGetOK{}
}

/* PcloudCloudconnectionsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsGetOK struct {
	Payload *models.CloudConnection
}

func (o *PcloudCloudconnectionsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsGetOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudconnectionsGetOK) GetPayload() *models.CloudConnection {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetBadRequest creates a PcloudCloudconnectionsGetBadRequest with default headers values
func NewPcloudCloudconnectionsGetBadRequest() *PcloudCloudconnectionsGetBadRequest {
	return &PcloudCloudconnectionsGetBadRequest{}
}

/* PcloudCloudconnectionsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsGetBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudconnectionsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetUnauthorized creates a PcloudCloudconnectionsGetUnauthorized with default headers values
func NewPcloudCloudconnectionsGetUnauthorized() *PcloudCloudconnectionsGetUnauthorized {
	return &PcloudCloudconnectionsGetUnauthorized{}
}

/* PcloudCloudconnectionsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudconnectionsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetNotFound creates a PcloudCloudconnectionsGetNotFound with default headers values
func NewPcloudCloudconnectionsGetNotFound() *PcloudCloudconnectionsGetNotFound {
	return &PcloudCloudconnectionsGetNotFound{}
}

/* PcloudCloudconnectionsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudCloudconnectionsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsGetInternalServerError creates a PcloudCloudconnectionsGetInternalServerError with default headers values
func NewPcloudCloudconnectionsGetInternalServerError() *PcloudCloudconnectionsGetInternalServerError {
	return &PcloudCloudconnectionsGetInternalServerError{}
}

/* PcloudCloudconnectionsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudconnectionsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
