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

// ListSiteBuildsReader is a Reader for the ListSiteBuilds structure.
type ListSiteBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSiteBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSiteBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSiteBuildsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSiteBuildsOK creates a ListSiteBuildsOK with default headers values
func NewListSiteBuildsOK() *ListSiteBuildsOK {
	return &ListSiteBuildsOK{}
}

/*ListSiteBuildsOK handles this case with default header values.

OK
*/
type ListSiteBuildsOK struct {
	Payload []*models.Build
}

func (o *ListSiteBuildsOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/builds][%d] listSiteBuildsOK  %+v", 200, o.Payload)
}

func (o *ListSiteBuildsOK) GetPayload() []*models.Build {
	return o.Payload
}

func (o *ListSiteBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSiteBuildsDefault creates a ListSiteBuildsDefault with default headers values
func NewListSiteBuildsDefault(code int) *ListSiteBuildsDefault {
	return &ListSiteBuildsDefault{
		_statusCode: code,
	}
}

/*ListSiteBuildsDefault handles this case with default header values.

error
*/
type ListSiteBuildsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list site builds default response
func (o *ListSiteBuildsDefault) Code() int {
	return o._statusCode
}

func (o *ListSiteBuildsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/builds][%d] listSiteBuilds default  %+v", o._statusCode, o.Payload)
}

func (o *ListSiteBuildsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSiteBuildsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
