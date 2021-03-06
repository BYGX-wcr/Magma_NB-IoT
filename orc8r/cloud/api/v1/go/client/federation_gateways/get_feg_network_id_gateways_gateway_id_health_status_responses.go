// Code generated by go-swagger; DO NOT EDIT.

package federation_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetFegNetworkIDGatewaysGatewayIDHealthStatusReader is a Reader for the GetFegNetworkIDGatewaysGatewayIDHealthStatus structure.
type GetFegNetworkIDGatewaysGatewayIDHealthStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegNetworkIDGatewaysGatewayIDHealthStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegNetworkIDGatewaysGatewayIDHealthStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegNetworkIDGatewaysGatewayIDHealthStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegNetworkIDGatewaysGatewayIDHealthStatusOK creates a GetFegNetworkIDGatewaysGatewayIDHealthStatusOK with default headers values
func NewGetFegNetworkIDGatewaysGatewayIDHealthStatusOK() *GetFegNetworkIDGatewaysGatewayIDHealthStatusOK {
	return &GetFegNetworkIDGatewaysGatewayIDHealthStatusOK{}
}

/*GetFegNetworkIDGatewaysGatewayIDHealthStatusOK handles this case with default header values.

Health Status of Federation Gateway
*/
type GetFegNetworkIDGatewaysGatewayIDHealthStatusOK struct {
	Payload *models.FederationGatewayHealthStatus
}

func (o *GetFegNetworkIDGatewaysGatewayIDHealthStatusOK) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/gateways/{gateway_id}/health_status][%d] getFegNetworkIdGatewaysGatewayIdHealthStatusOK  %+v", 200, o.Payload)
}

func (o *GetFegNetworkIDGatewaysGatewayIDHealthStatusOK) GetPayload() *models.FederationGatewayHealthStatus {
	return o.Payload
}

func (o *GetFegNetworkIDGatewaysGatewayIDHealthStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FederationGatewayHealthStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegNetworkIDGatewaysGatewayIDHealthStatusDefault creates a GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault with default headers values
func NewGetFegNetworkIDGatewaysGatewayIDHealthStatusDefault(code int) *GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault {
	return &GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault{
		_statusCode: code,
	}
}

/*GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault handles this case with default header values.

Unexpected Error
*/
type GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg network ID gateways gateway ID health status default response
func (o *GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/gateways/{gateway_id}/health_status][%d] GetFegNetworkIDGatewaysGatewayIDHealthStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegNetworkIDGatewaysGatewayIDHealthStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
