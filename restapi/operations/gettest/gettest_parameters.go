// Code generated by go-swagger; DO NOT EDIT.

package gettest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGettestParams creates a new GettestParams object
//
// There are no default values defined in the spec.
func NewGettestParams() GettestParams {

	return GettestParams{}
}

// GettestParams contains all the bound params for the gettest operation
// typically these are obtained from a http.Request
//
// swagger:parameters gettest
type GettestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*any enum
	  In: query
	*/
	Enum *string
	/*any number
	  In: query
	*/
	Num *int32
	/*any string
	  In: query
	*/
	Str *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGettestParams() beforehand.
func (o *GettestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEnum, qhkEnum, _ := qs.GetOK("enum")
	if err := o.bindEnum(qEnum, qhkEnum, route.Formats); err != nil {
		res = append(res, err)
	}

	qNum, qhkNum, _ := qs.GetOK("num")
	if err := o.bindNum(qNum, qhkNum, route.Formats); err != nil {
		res = append(res, err)
	}

	qStr, qhkStr, _ := qs.GetOK("str")
	if err := o.bindStr(qStr, qhkStr, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEnum binds and validates parameter Enum from query.
func (o *GettestParams) bindEnum(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Enum = &raw

	if err := o.validateEnum(formats); err != nil {
		return err
	}

	return nil
}

// validateEnum carries on validations for parameter Enum
func (o *GettestParams) validateEnum(formats strfmt.Registry) error {

	if err := validate.EnumCase("enum", "query", *o.Enum, []interface{}{"foo", "bar"}, true); err != nil {
		return err
	}

	return nil
}

// bindNum binds and validates parameter Num from query.
func (o *GettestParams) bindNum(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("num", "query", "int32", raw)
	}
	o.Num = &value

	return nil
}

// bindStr binds and validates parameter Str from query.
func (o *GettestParams) bindStr(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Str = &raw

	return nil
}
