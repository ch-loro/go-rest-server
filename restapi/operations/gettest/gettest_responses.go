// Code generated by go-swagger; DO NOT EDIT.

package gettest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/tetsk/go-server/models"
)

// GettestOKCode is the HTTP code returned for type GettestOK
const GettestOKCode int = 200

/*GettestOK success

swagger:response gettestOK
*/
type GettestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Obj `json:"body,omitempty"`
}

// NewGettestOK creates GettestOK with default headers values
func NewGettestOK() *GettestOK {

	return &GettestOK{}
}

// WithPayload adds the payload to the gettest o k response
func (o *GettestOK) WithPayload(payload []*models.Obj) *GettestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the gettest o k response
func (o *GettestOK) SetPayload(payload []*models.Obj) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GettestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Obj, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GettestDefault Unexpected error

swagger:response gettestDefault
*/
type GettestDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGettestDefault creates GettestDefault with default headers values
func NewGettestDefault(code int) *GettestDefault {
	if code <= 0 {
		code = 500
	}

	return &GettestDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the gettest default response
func (o *GettestDefault) WithStatusCode(code int) *GettestDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the gettest default response
func (o *GettestDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the gettest default response
func (o *GettestDefault) WithPayload(payload *models.Error) *GettestDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the gettest default response
func (o *GettestDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GettestDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
