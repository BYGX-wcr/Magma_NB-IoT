// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetWifiNetworkIDNameReader is a Reader for the GetWifiNetworkIDName structure.
type GetWifiNetworkIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWifiNetworkIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWifiNetworkIDNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWifiNetworkIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWifiNetworkIDNameOK creates a GetWifiNetworkIDNameOK with default headers values
func NewGetWifiNetworkIDNameOK() *GetWifiNetworkIDNameOK {
	return &GetWifiNetworkIDNameOK{}
}

/*GetWifiNetworkIDNameOK handles this case with default header values.

Name of the network
*/
type GetWifiNetworkIDNameOK struct {
	Payload models.NetworkName
}

func (o *GetWifiNetworkIDNameOK) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/name][%d] getWifiNetworkIdNameOK  %+v", 200, o.Payload)
}

func (o *GetWifiNetworkIDNameOK) GetPayload() models.NetworkName {
	return o.Payload
}

func (o *GetWifiNetworkIDNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWifiNetworkIDNameDefault creates a GetWifiNetworkIDNameDefault with default headers values
func NewGetWifiNetworkIDNameDefault(code int) *GetWifiNetworkIDNameDefault {
	return &GetWifiNetworkIDNameDefault{
		_statusCode: code,
	}
}

/*GetWifiNetworkIDNameDefault handles this case with default header values.

Unexpected Error
*/
type GetWifiNetworkIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wifi network ID name default response
func (o *GetWifiNetworkIDNameDefault) Code() int {
	return o._statusCode
}

func (o *GetWifiNetworkIDNameDefault) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/name][%d] GetWifiNetworkIDName default  %+v", o._statusCode, o.Payload)
}

func (o *GetWifiNetworkIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWifiNetworkIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
