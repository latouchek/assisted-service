// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2DownloadClusterManifestReader is a Reader for the V2DownloadClusterManifest structure.
type V2DownloadClusterManifestReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *V2DownloadClusterManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2DownloadClusterManifestOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2DownloadClusterManifestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2DownloadClusterManifestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2DownloadClusterManifestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2DownloadClusterManifestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2DownloadClusterManifestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2DownloadClusterManifestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2DownloadClusterManifestOK creates a V2DownloadClusterManifestOK with default headers values
func NewV2DownloadClusterManifestOK(writer io.Writer) *V2DownloadClusterManifestOK {
	return &V2DownloadClusterManifestOK{
		Payload: writer,
	}
}

/*V2DownloadClusterManifestOK handles this case with default header values.

Success.
*/
type V2DownloadClusterManifestOK struct {
	Payload io.Writer
}

func (o *V2DownloadClusterManifestOK) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests/files][%d] v2DownloadClusterManifestOK  %+v", 200, o.Payload)
}

func (o *V2DownloadClusterManifestOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *V2DownloadClusterManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadClusterManifestUnauthorized creates a V2DownloadClusterManifestUnauthorized with default headers values
func NewV2DownloadClusterManifestUnauthorized() *V2DownloadClusterManifestUnauthorized {
	return &V2DownloadClusterManifestUnauthorized{}
}

/*V2DownloadClusterManifestUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2DownloadClusterManifestUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2DownloadClusterManifestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests/files][%d] v2DownloadClusterManifestUnauthorized  %+v", 401, o.Payload)
}

func (o *V2DownloadClusterManifestUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2DownloadClusterManifestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadClusterManifestForbidden creates a V2DownloadClusterManifestForbidden with default headers values
func NewV2DownloadClusterManifestForbidden() *V2DownloadClusterManifestForbidden {
	return &V2DownloadClusterManifestForbidden{}
}

/*V2DownloadClusterManifestForbidden handles this case with default header values.

Forbidden.
*/
type V2DownloadClusterManifestForbidden struct {
	Payload *models.InfraError
}

func (o *V2DownloadClusterManifestForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests/files][%d] v2DownloadClusterManifestForbidden  %+v", 403, o.Payload)
}

func (o *V2DownloadClusterManifestForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2DownloadClusterManifestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadClusterManifestNotFound creates a V2DownloadClusterManifestNotFound with default headers values
func NewV2DownloadClusterManifestNotFound() *V2DownloadClusterManifestNotFound {
	return &V2DownloadClusterManifestNotFound{}
}

/*V2DownloadClusterManifestNotFound handles this case with default header values.

Error.
*/
type V2DownloadClusterManifestNotFound struct {
	Payload *models.Error
}

func (o *V2DownloadClusterManifestNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests/files][%d] v2DownloadClusterManifestNotFound  %+v", 404, o.Payload)
}

func (o *V2DownloadClusterManifestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadClusterManifestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadClusterManifestMethodNotAllowed creates a V2DownloadClusterManifestMethodNotAllowed with default headers values
func NewV2DownloadClusterManifestMethodNotAllowed() *V2DownloadClusterManifestMethodNotAllowed {
	return &V2DownloadClusterManifestMethodNotAllowed{}
}

/*V2DownloadClusterManifestMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2DownloadClusterManifestMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2DownloadClusterManifestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests/files][%d] v2DownloadClusterManifestMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2DownloadClusterManifestMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadClusterManifestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadClusterManifestConflict creates a V2DownloadClusterManifestConflict with default headers values
func NewV2DownloadClusterManifestConflict() *V2DownloadClusterManifestConflict {
	return &V2DownloadClusterManifestConflict{}
}

/*V2DownloadClusterManifestConflict handles this case with default header values.

Error.
*/
type V2DownloadClusterManifestConflict struct {
	Payload *models.Error
}

func (o *V2DownloadClusterManifestConflict) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests/files][%d] v2DownloadClusterManifestConflict  %+v", 409, o.Payload)
}

func (o *V2DownloadClusterManifestConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadClusterManifestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DownloadClusterManifestInternalServerError creates a V2DownloadClusterManifestInternalServerError with default headers values
func NewV2DownloadClusterManifestInternalServerError() *V2DownloadClusterManifestInternalServerError {
	return &V2DownloadClusterManifestInternalServerError{}
}

/*V2DownloadClusterManifestInternalServerError handles this case with default header values.

Error.
*/
type V2DownloadClusterManifestInternalServerError struct {
	Payload *models.Error
}

func (o *V2DownloadClusterManifestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/manifests/files][%d] v2DownloadClusterManifestInternalServerError  %+v", 500, o.Payload)
}

func (o *V2DownloadClusterManifestInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DownloadClusterManifestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
