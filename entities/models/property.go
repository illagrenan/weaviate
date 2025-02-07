//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Property property
//
// swagger:model Property
type Property struct {

	// Can be a reference to another type when it starts with a capital (for example Person), otherwise "string" or "int".
	DataType []string `json:"dataType"`

	// Description of the property.
	Description string `json:"description,omitempty"`

	// Optional. Should this property be indexed in the inverted index. Defaults to true. If you choose false, you will not be able to use this property in where filters. This property has no affect on vectorization decisions done by modules
	IndexInverted *bool `json:"indexInverted,omitempty"`

	// Configuratino specific to modules this Weaviate instance has installed
	ModuleConfig interface{} `json:"moduleConfig,omitempty"`

	// Name of the property as URI relative to the schema URL.
	Name string `json:"name,omitempty"`
}

// Validate validates this property
func (m *Property) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Property) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Property) UnmarshalBinary(b []byte) error {
	var res Property
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
