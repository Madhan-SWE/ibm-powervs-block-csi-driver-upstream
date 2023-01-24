// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPlacementgroupsPostReader is a Reader for the PcloudPlacementgroupsPost structure.
type PcloudPlacementgroupsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPlacementgroupsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPlacementgroupsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPlacementgroupsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPlacementgroupsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPlacementgroupsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPlacementgroupsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPlacementgroupsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPlacementgroupsPostOK creates a PcloudPlacementgroupsPostOK with default headers values
func NewPcloudPlacementgroupsPostOK() *PcloudPlacementgroupsPostOK {
	return &PcloudPlacementgroupsPostOK{}
}

/*
PcloudPlacementgroupsPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPlacementgroupsPostOK struct {
	Payload *models.PlacementGroup
}

// IsSuccess returns true when this pcloud placementgroups post o k response has a 2xx status code
func (o *PcloudPlacementgroupsPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud placementgroups post o k response has a 3xx status code
func (o *PcloudPlacementgroupsPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups post o k response has a 4xx status code
func (o *PcloudPlacementgroupsPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups post o k response has a 5xx status code
func (o *PcloudPlacementgroupsPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups post o k response a status code equal to that given
func (o *PcloudPlacementgroupsPostOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudPlacementgroupsPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostOK  %+v", 200, o.Payload)
}

func (o *PcloudPlacementgroupsPostOK) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostOK  %+v", 200, o.Payload)
}

func (o *PcloudPlacementgroupsPostOK) GetPayload() *models.PlacementGroup {
	return o.Payload
}

func (o *PcloudPlacementgroupsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsPostBadRequest creates a PcloudPlacementgroupsPostBadRequest with default headers values
func NewPcloudPlacementgroupsPostBadRequest() *PcloudPlacementgroupsPostBadRequest {
	return &PcloudPlacementgroupsPostBadRequest{}
}

/*
PcloudPlacementgroupsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPlacementgroupsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups post bad request response has a 2xx status code
func (o *PcloudPlacementgroupsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups post bad request response has a 3xx status code
func (o *PcloudPlacementgroupsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups post bad request response has a 4xx status code
func (o *PcloudPlacementgroupsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups post bad request response has a 5xx status code
func (o *PcloudPlacementgroupsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups post bad request response a status code equal to that given
func (o *PcloudPlacementgroupsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPlacementgroupsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPlacementgroupsPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPlacementgroupsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsPostForbidden creates a PcloudPlacementgroupsPostForbidden with default headers values
func NewPcloudPlacementgroupsPostForbidden() *PcloudPlacementgroupsPostForbidden {
	return &PcloudPlacementgroupsPostForbidden{}
}

/*
PcloudPlacementgroupsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPlacementgroupsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups post forbidden response has a 2xx status code
func (o *PcloudPlacementgroupsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups post forbidden response has a 3xx status code
func (o *PcloudPlacementgroupsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups post forbidden response has a 4xx status code
func (o *PcloudPlacementgroupsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups post forbidden response has a 5xx status code
func (o *PcloudPlacementgroupsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups post forbidden response a status code equal to that given
func (o *PcloudPlacementgroupsPostForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudPlacementgroupsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPlacementgroupsPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPlacementgroupsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsPostConflict creates a PcloudPlacementgroupsPostConflict with default headers values
func NewPcloudPlacementgroupsPostConflict() *PcloudPlacementgroupsPostConflict {
	return &PcloudPlacementgroupsPostConflict{}
}

/*
PcloudPlacementgroupsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPlacementgroupsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups post conflict response has a 2xx status code
func (o *PcloudPlacementgroupsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups post conflict response has a 3xx status code
func (o *PcloudPlacementgroupsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups post conflict response has a 4xx status code
func (o *PcloudPlacementgroupsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups post conflict response has a 5xx status code
func (o *PcloudPlacementgroupsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups post conflict response a status code equal to that given
func (o *PcloudPlacementgroupsPostConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PcloudPlacementgroupsPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPlacementgroupsPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPlacementgroupsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsPostUnprocessableEntity creates a PcloudPlacementgroupsPostUnprocessableEntity with default headers values
func NewPcloudPlacementgroupsPostUnprocessableEntity() *PcloudPlacementgroupsPostUnprocessableEntity {
	return &PcloudPlacementgroupsPostUnprocessableEntity{}
}

/*
PcloudPlacementgroupsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPlacementgroupsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups post unprocessable entity response has a 2xx status code
func (o *PcloudPlacementgroupsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups post unprocessable entity response has a 3xx status code
func (o *PcloudPlacementgroupsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups post unprocessable entity response has a 4xx status code
func (o *PcloudPlacementgroupsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups post unprocessable entity response has a 5xx status code
func (o *PcloudPlacementgroupsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups post unprocessable entity response a status code equal to that given
func (o *PcloudPlacementgroupsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudPlacementgroupsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPlacementgroupsPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPlacementgroupsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsPostInternalServerError creates a PcloudPlacementgroupsPostInternalServerError with default headers values
func NewPcloudPlacementgroupsPostInternalServerError() *PcloudPlacementgroupsPostInternalServerError {
	return &PcloudPlacementgroupsPostInternalServerError{}
}

/*
PcloudPlacementgroupsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPlacementgroupsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups post internal server error response has a 2xx status code
func (o *PcloudPlacementgroupsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups post internal server error response has a 3xx status code
func (o *PcloudPlacementgroupsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups post internal server error response has a 4xx status code
func (o *PcloudPlacementgroupsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups post internal server error response has a 5xx status code
func (o *PcloudPlacementgroupsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud placementgroups post internal server error response a status code equal to that given
func (o *PcloudPlacementgroupsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPlacementgroupsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPlacementgroupsPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups][%d] pcloudPlacementgroupsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPlacementgroupsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
