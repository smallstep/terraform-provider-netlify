// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/open-api/go/models"
)

// UploadDeployFunctionReader is a Reader for the UploadDeployFunction structure.
type UploadDeployFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadDeployFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadDeployFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUploadDeployFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadDeployFunctionOK creates a UploadDeployFunctionOK with default headers values
func NewUploadDeployFunctionOK() *UploadDeployFunctionOK {
	return &UploadDeployFunctionOK{}
}

/*UploadDeployFunctionOK handles this case with default header values.

OK
*/
type UploadDeployFunctionOK struct {
	Payload *models.Function
}

func (o *UploadDeployFunctionOK) Error() string {
	return fmt.Sprintf("[PUT /deploys/{deploy_id}/functions/{name}][%d] uploadDeployFunctionOK  %+v", 200, o.Payload)
}

func (o *UploadDeployFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadDeployFunctionDefault creates a UploadDeployFunctionDefault with default headers values
func NewUploadDeployFunctionDefault(code int) *UploadDeployFunctionDefault {
	return &UploadDeployFunctionDefault{
		_statusCode: code,
	}
}

/*UploadDeployFunctionDefault handles this case with default header values.

error
*/
type UploadDeployFunctionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the upload deploy function default response
func (o *UploadDeployFunctionDefault) Code() int {
	return o._statusCode
}

func (o *UploadDeployFunctionDefault) Error() string {
	return fmt.Sprintf("[PUT /deploys/{deploy_id}/functions/{name}][%d] uploadDeployFunction default  %+v", o._statusCode, o.Payload)
}

func (o *UploadDeployFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}