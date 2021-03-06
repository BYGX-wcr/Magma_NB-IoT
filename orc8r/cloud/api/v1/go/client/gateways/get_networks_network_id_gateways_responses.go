// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDGatewaysReader is a Reader for the GetNetworksNetworkIDGateways structure.
type GetNetworksNetworkIDGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDGatewaysOK creates a GetNetworksNetworkIDGatewaysOK with default headers values
func NewGetNetworksNetworkIDGatewaysOK() *GetNetworksNetworkIDGatewaysOK {
	return &GetNetworksNetworkIDGatewaysOK{}
}

/*GetNetworksNetworkIDGatewaysOK handles this case with default header values.

Map of all gateways inside the network by gatewayID
*/
type GetNetworksNetworkIDGatewaysOK struct {
	Payload map[string]models.MagmadGateway
}

func (o *GetNetworksNetworkIDGatewaysOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways][%d] getNetworksNetworkIdGatewaysOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysOK) GetPayload() map[string]models.MagmadGateway {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDGatewaysDefault creates a GetNetworksNetworkIDGatewaysDefault with default headers values
func NewGetNetworksNetworkIDGatewaysDefault(code int) *GetNetworksNetworkIDGatewaysDefault {
	return &GetNetworksNetworkIDGatewaysDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDGatewaysDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDGatewaysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID gateways default response
func (o *GetNetworksNetworkIDGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDGatewaysDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways][%d] GetNetworksNetworkIDGateways default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
