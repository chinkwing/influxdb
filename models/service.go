package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Service service

swagger:model Service
*/
type Service struct {

	/* The key part of the key-value pair that makes up a tag. Used to identify a service type.
	 */
	TagKey string `json:"tagKey,omitempty"`

	/* The value part of the key-value pair that makes up a tag.
	 */
	TagValue string `json:"tagValue,omitempty"`

	/* Type of service
	 */
	Type string `json:"type,omitempty"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceTypeTypePropEnum []interface{}

// prop value enum
func (m *Service) validateTypeEnum(path, location string, value string) error {
	if serviceTypeTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["containers","kubernetes","host"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serviceTypeTypePropEnum = append(serviceTypeTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serviceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Service) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
