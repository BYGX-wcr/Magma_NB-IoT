// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PostNetworksNetworkIDTiersTierIDImagesReader is a Reader for the PostNetworksNetworkIDTiersTierIDImages structure.
type PostNetworksNetworkIDTiersTierIDImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDTiersTierIDImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostNetworksNetworkIDTiersTierIDImagesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDTiersTierIDImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDTiersTierIDImagesNoContent creates a PostNetworksNetworkIDTiersTierIDImagesNoContent with default headers values
func NewPostNetworksNetworkIDTiersTierIDImagesNoContent() *PostNetworksNetworkIDTiersTierIDImagesNoContent {
	return &PostNetworksNetworkIDTiersTierIDImagesNoContent{}
}

/*PostNetworksNetworkIDTiersTierIDImagesNoContent handles this case with default header values.

Success
*/
type PostNetworksNetworkIDTiersTierIDImagesNoContent struct {
}

func (o *PostNetworksNetworkIDTiersTierIDImagesNoContent) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/tiers/{tier_id}/images][%d] postNetworksNetworkIdTiersTierIdImagesNoContent ", 204)
}

func (o *PostNetworksNetworkIDTiersTierIDImagesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostNetworksNetworkIDTiersTierIDImagesDefault creates a PostNetworksNetworkIDTiersTierIDImagesDefault with default headers values
func NewPostNetworksNetworkIDTiersTierIDImagesDefault(code int) *PostNetworksNetworkIDTiersTierIDImagesDefault {
	return &PostNetworksNetworkIDTiersTierIDImagesDefault{
		_statusCode: code,
	}
}

/*PostNetworksNetworkIDTiersTierIDImagesDefault handles this case with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDTiersTierIDImagesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID tiers tier ID images default response
func (o *PostNetworksNetworkIDTiersTierIDImagesDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDTiersTierIDImagesDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/tiers/{tier_id}/images][%d] PostNetworksNetworkIDTiersTierIDImages default  %+v", o._statusCode, o.Payload)
}

func (o *PostNetworksNetworkIDTiersTierIDImagesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDTiersTierIDImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
