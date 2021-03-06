// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutWifiNetworkIDGatewaysGatewayIDDescriptionReader is a Reader for the PutWifiNetworkIDGatewaysGatewayIDDescription structure.
type PutWifiNetworkIDGatewaysGatewayIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWifiNetworkIDGatewaysGatewayIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutWifiNetworkIDGatewaysGatewayIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent creates a PutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent() *PutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent {
	return &PutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent{}
}

/*PutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent handles this case with default header values.

Success
*/
type PutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent struct {
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/description][%d] putWifiNetworkIdGatewaysGatewayIdDescriptionNoContent ", 204)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDescriptionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWifiNetworkIDGatewaysGatewayIDDescriptionDefault creates a PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDDescriptionDefault(code int) *PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault {
	return &PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault{
		_statusCode: code,
	}
}

/*PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault handles this case with default header values.

Unexpected Error
*/
type PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put wifi network ID gateways gateway ID description default response
func (o *PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/description][%d] PutWifiNetworkIDGatewaysGatewayIDDescription default  %+v", o._statusCode, o.Payload)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
