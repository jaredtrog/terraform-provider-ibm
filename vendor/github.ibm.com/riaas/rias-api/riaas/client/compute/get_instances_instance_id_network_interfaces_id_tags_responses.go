// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetInstancesInstanceIDNetworkInterfacesIDTagsReader is a Reader for the GetInstancesInstanceIDNetworkInterfacesIDTags structure.
type GetInstancesInstanceIDNetworkInterfacesIDTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesInstanceIDNetworkInterfacesIDTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInstancesInstanceIDNetworkInterfacesIDTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInstancesInstanceIDNetworkInterfacesIDTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetInstancesInstanceIDNetworkInterfacesIDTagsOK creates a GetInstancesInstanceIDNetworkInterfacesIDTagsOK with default headers values
func NewGetInstancesInstanceIDNetworkInterfacesIDTagsOK() *GetInstancesInstanceIDNetworkInterfacesIDTagsOK {
	return &GetInstancesInstanceIDNetworkInterfacesIDTagsOK{}
}

/*GetInstancesInstanceIDNetworkInterfacesIDTagsOK handles this case with default header values.

dummy
*/
type GetInstancesInstanceIDNetworkInterfacesIDTagsOK struct {
	Payload *models.GetInstancesInstanceIDNetworkInterfacesIDTagsOKBody
}

func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsOK) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/network_interfaces/{id}/tags][%d] getInstancesInstanceIdNetworkInterfacesIdTagsOK  %+v", 200, o.Payload)
}

func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetInstancesInstanceIDNetworkInterfacesIDTagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDNetworkInterfacesIDTagsNotFound creates a GetInstancesInstanceIDNetworkInterfacesIDTagsNotFound with default headers values
func NewGetInstancesInstanceIDNetworkInterfacesIDTagsNotFound() *GetInstancesInstanceIDNetworkInterfacesIDTagsNotFound {
	return &GetInstancesInstanceIDNetworkInterfacesIDTagsNotFound{}
}

/*GetInstancesInstanceIDNetworkInterfacesIDTagsNotFound handles this case with default header values.

error
*/
type GetInstancesInstanceIDNetworkInterfacesIDTagsNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/network_interfaces/{id}/tags][%d] getInstancesInstanceIdNetworkInterfacesIdTagsNotFound  %+v", 404, o.Payload)
}

func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDNetworkInterfacesIDTagsDefault creates a GetInstancesInstanceIDNetworkInterfacesIDTagsDefault with default headers values
func NewGetInstancesInstanceIDNetworkInterfacesIDTagsDefault(code int) *GetInstancesInstanceIDNetworkInterfacesIDTagsDefault {
	return &GetInstancesInstanceIDNetworkInterfacesIDTagsDefault{
		_statusCode: code,
	}
}

/*GetInstancesInstanceIDNetworkInterfacesIDTagsDefault handles this case with default header values.

unexpectederror
*/
type GetInstancesInstanceIDNetworkInterfacesIDTagsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get instances instance ID network interfaces ID tags default response
func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsDefault) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/network_interfaces/{id}/tags][%d] GetInstancesInstanceIDNetworkInterfacesIDTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstancesInstanceIDNetworkInterfacesIDTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}