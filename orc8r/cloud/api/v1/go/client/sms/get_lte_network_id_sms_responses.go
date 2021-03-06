// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDSMSReader is a Reader for the GetLTENetworkIDSMS structure.
type GetLTENetworkIDSMSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDSMSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDSMSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDSMSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDSMSOK creates a GetLTENetworkIDSMSOK with default headers values
func NewGetLTENetworkIDSMSOK() *GetLTENetworkIDSMSOK {
	return &GetLTENetworkIDSMSOK{}
}

/*GetLTENetworkIDSMSOK handles this case with default header values.

List all SMS's in the system
*/
type GetLTENetworkIDSMSOK struct {
	Payload []*models.SMSMessage
}

func (o *GetLTENetworkIDSMSOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/sms][%d] getLteNetworkIdSmsOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDSMSOK) GetPayload() []*models.SMSMessage {
	return o.Payload
}

func (o *GetLTENetworkIDSMSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDSMSDefault creates a GetLTENetworkIDSMSDefault with default headers values
func NewGetLTENetworkIDSMSDefault(code int) *GetLTENetworkIDSMSDefault {
	return &GetLTENetworkIDSMSDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDSMSDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDSMSDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID SMS default response
func (o *GetLTENetworkIDSMSDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDSMSDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/sms][%d] GetLTENetworkIDSMS default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDSMSDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDSMSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
