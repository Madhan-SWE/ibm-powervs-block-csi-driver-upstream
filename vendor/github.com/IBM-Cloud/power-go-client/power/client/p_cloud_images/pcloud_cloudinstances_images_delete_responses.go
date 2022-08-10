// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesImagesDeleteReader is a Reader for the PcloudCloudinstancesImagesDelete structure.
type PcloudCloudinstancesImagesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesImagesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesImagesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesImagesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesImagesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudinstancesImagesDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesImagesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesImagesDeleteOK creates a PcloudCloudinstancesImagesDeleteOK with default headers values
func NewPcloudCloudinstancesImagesDeleteOK() *PcloudCloudinstancesImagesDeleteOK {
	return &PcloudCloudinstancesImagesDeleteOK{}
}

/* PcloudCloudinstancesImagesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesImagesDeleteOK struct {
	Payload models.Object
}

func (o *PcloudCloudinstancesImagesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudinstancesImagesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteBadRequest creates a PcloudCloudinstancesImagesDeleteBadRequest with default headers values
func NewPcloudCloudinstancesImagesDeleteBadRequest() *PcloudCloudinstancesImagesDeleteBadRequest {
	return &PcloudCloudinstancesImagesDeleteBadRequest{}
}

/* PcloudCloudinstancesImagesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesImagesDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudinstancesImagesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteUnauthorized creates a PcloudCloudinstancesImagesDeleteUnauthorized with default headers values
func NewPcloudCloudinstancesImagesDeleteUnauthorized() *PcloudCloudinstancesImagesDeleteUnauthorized {
	return &PcloudCloudinstancesImagesDeleteUnauthorized{}
}

/* PcloudCloudinstancesImagesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesImagesDeleteUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudinstancesImagesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteGone creates a PcloudCloudinstancesImagesDeleteGone with default headers values
func NewPcloudCloudinstancesImagesDeleteGone() *PcloudCloudinstancesImagesDeleteGone {
	return &PcloudCloudinstancesImagesDeleteGone{}
}

/* PcloudCloudinstancesImagesDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudinstancesImagesDeleteGone struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteGone  %+v", 410, o.Payload)
}
func (o *PcloudCloudinstancesImagesDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteInternalServerError creates a PcloudCloudinstancesImagesDeleteInternalServerError with default headers values
func NewPcloudCloudinstancesImagesDeleteInternalServerError() *PcloudCloudinstancesImagesDeleteInternalServerError {
	return &PcloudCloudinstancesImagesDeleteInternalServerError{}
}

/* PcloudCloudinstancesImagesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesImagesDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesImagesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudinstancesImagesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
