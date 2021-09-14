// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2ListSupportedOpenshiftVersionsReader is a Reader for the V2ListSupportedOpenshiftVersions structure.
type V2ListSupportedOpenshiftVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListSupportedOpenshiftVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListSupportedOpenshiftVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewV2ListSupportedOpenshiftVersionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListSupportedOpenshiftVersionsOK creates a V2ListSupportedOpenshiftVersionsOK with default headers values
func NewV2ListSupportedOpenshiftVersionsOK() *V2ListSupportedOpenshiftVersionsOK {
	return &V2ListSupportedOpenshiftVersionsOK{}
}

/*V2ListSupportedOpenshiftVersionsOK handles this case with default header values.

Success.
*/
type V2ListSupportedOpenshiftVersionsOK struct {
	Payload models.OpenshiftVersions
}

func (o *V2ListSupportedOpenshiftVersionsOK) Error() string {
	return fmt.Sprintf("[GET /v2/openshift-versions][%d] v2ListSupportedOpenshiftVersionsOK  %+v", 200, o.Payload)
}

func (o *V2ListSupportedOpenshiftVersionsOK) GetPayload() models.OpenshiftVersions {
	return o.Payload
}

func (o *V2ListSupportedOpenshiftVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListSupportedOpenshiftVersionsServiceUnavailable creates a V2ListSupportedOpenshiftVersionsServiceUnavailable with default headers values
func NewV2ListSupportedOpenshiftVersionsServiceUnavailable() *V2ListSupportedOpenshiftVersionsServiceUnavailable {
	return &V2ListSupportedOpenshiftVersionsServiceUnavailable{}
}

/*V2ListSupportedOpenshiftVersionsServiceUnavailable handles this case with default header values.

Unavailable.
*/
type V2ListSupportedOpenshiftVersionsServiceUnavailable struct {
	Payload *models.Error
}

func (o *V2ListSupportedOpenshiftVersionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/openshift-versions][%d] v2ListSupportedOpenshiftVersionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2ListSupportedOpenshiftVersionsServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListSupportedOpenshiftVersionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
