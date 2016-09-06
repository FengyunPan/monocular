package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/helm/monocular/pkg/swagger/models"
)

/*GetChartOK chart response

swagger:response getChartOK
*/
type GetChartOK struct {

	// In: body
	Payload GetChartOKBodyBody `json:"body,omitempty"`
}

// NewGetChartOK creates GetChartOK with default headers values
func NewGetChartOK() *GetChartOK {
	return &GetChartOK{}
}

// WithPayload adds the payload to the get chart o k response
func (o *GetChartOK) WithPayload(payload GetChartOKBodyBody) *GetChartOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chart o k response
func (o *GetChartOK) SetPayload(payload GetChartOKBodyBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChartOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetChartDefault unexpected error

swagger:response getChartDefault
*/
type GetChartDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChartDefault creates GetChartDefault with default headers values
func NewGetChartDefault(code int) *GetChartDefault {
	if code <= 0 {
		code = 500
	}

	return &GetChartDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get chart default response
func (o *GetChartDefault) WithStatusCode(code int) *GetChartDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get chart default response
func (o *GetChartDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get chart default response
func (o *GetChartDefault) WithPayload(payload *models.Error) *GetChartDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chart default response
func (o *GetChartDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChartDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}