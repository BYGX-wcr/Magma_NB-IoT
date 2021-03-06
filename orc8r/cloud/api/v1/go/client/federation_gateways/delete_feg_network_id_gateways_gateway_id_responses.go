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

// DeleteFegNetworkIDGatewaysGatewayIDReader is a Reader for the DeleteFegNetworkIDGatewaysGatewayID structure.
type DeleteFegNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFegNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFegNetworkIDGatewaysGatewayIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteFegNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFegNetworkIDGatewaysGatewayIDNoContent creates a DeleteFegNetworkIDGatewaysGatewayIDNoContent with default headers values
func NewDeleteFegNetworkIDGatewaysGatewayIDNoContent() *DeleteFegNetworkIDGatewaysGatewayIDNoContent {
	return &DeleteFegNetworkIDGatewaysGatewayIDNoContent{}
}

/*DeleteFegNetworkIDGatewaysGatewayIDNoContent handles this case with default header values.

Success
*/
type DeleteFegNetworkIDGatewaysGatewayIDNoContent struct {
}

func (o *DeleteFegNetworkIDGatewaysGatewayIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /feg/{network_id}/gateways/{gateway_id}][%d] deleteFegNetworkIdGatewaysGatewayIdNoContent ", 204)
}

func (o *DeleteFegNetworkIDGatewaysGatewayIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFegNetworkIDGatewaysGatewayIDDefault creates a DeleteFegNetworkIDGatewaysGatewayIDDefault with default headers values
func NewDeleteFegNetworkIDGatewaysGatewayIDDefault(code int) *DeleteFegNetworkIDGatewaysGatewayIDDefault {
	return &DeleteFegNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/*DeleteFegNetworkIDGatewaysGatewayIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteFegNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete feg network ID gateways gateway ID default response
func (o *DeleteFegNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFegNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /feg/{network_id}/gateways/{gateway_id}][%d] DeleteFegNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFegNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFegNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
