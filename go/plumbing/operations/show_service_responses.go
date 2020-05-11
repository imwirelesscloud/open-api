// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// ShowServiceReader is a Reader for the ShowService structure.
type ShowServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewShowServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowServiceOK creates a ShowServiceOK with default headers values
func NewShowServiceOK() *ShowServiceOK {
	return &ShowServiceOK{}
}

/*ShowServiceOK handles this case with default header values.

services
*/
type ShowServiceOK struct {
	Payload *models.ServiceInstance
}

func (o *ShowServiceOK) Error() string {
	return fmt.Sprintf("[GET /services/{addonName}][%d] showServiceOK  %+v", 200, o.Payload)
}

func (o *ShowServiceOK) GetPayload() *models.ServiceInstance {
	return o.Payload
}

func (o *ShowServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowServiceDefault creates a ShowServiceDefault with default headers values
func NewShowServiceDefault(code int) *ShowServiceDefault {
	return &ShowServiceDefault{
		_statusCode: code,
	}
}

/*ShowServiceDefault handles this case with default header values.

error
*/
type ShowServiceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the show service default response
func (o *ShowServiceDefault) Code() int {
	return o._statusCode
}

func (o *ShowServiceDefault) Error() string {
	return fmt.Sprintf("[GET /services/{addonName}][%d] showService default  %+v", o._statusCode, o.Payload)
}

func (o *ShowServiceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ShowServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
