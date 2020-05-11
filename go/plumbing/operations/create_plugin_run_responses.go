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

// CreatePluginRunReader is a Reader for the CreatePluginRun structure.
type CreatePluginRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePluginRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePluginRunCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePluginRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePluginRunCreated creates a CreatePluginRunCreated with default headers values
func NewCreatePluginRunCreated() *CreatePluginRunCreated {
	return &CreatePluginRunCreated{}
}

/*CreatePluginRunCreated handles this case with default header values.

CREATED
*/
type CreatePluginRunCreated struct {
	Payload *models.PluginRun
}

func (o *CreatePluginRunCreated) Error() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/plugin_runs][%d] createPluginRunCreated  %+v", 201, o.Payload)
}

func (o *CreatePluginRunCreated) GetPayload() *models.PluginRun {
	return o.Payload
}

func (o *CreatePluginRunCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PluginRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePluginRunDefault creates a CreatePluginRunDefault with default headers values
func NewCreatePluginRunDefault(code int) *CreatePluginRunDefault {
	return &CreatePluginRunDefault{
		_statusCode: code,
	}
}

/*CreatePluginRunDefault handles this case with default header values.

error
*/
type CreatePluginRunDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create plugin run default response
func (o *CreatePluginRunDefault) Code() int {
	return o._statusCode
}

func (o *CreatePluginRunDefault) Error() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/plugin_runs][%d] createPluginRun default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePluginRunDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePluginRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
