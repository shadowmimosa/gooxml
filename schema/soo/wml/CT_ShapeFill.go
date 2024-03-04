// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_ShapeFill struct {
	OnAttr        *string
	FocussizeAttr *string
}

func NewCT_ShapeFill() *CT_ShapeFill {
	return &CT_ShapeFill{}
}

func (m *CT_ShapeFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// 属性编码 -  文本
	if m.OnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "on"},
			Value: fmt.Sprintf("%v", *m.OnAttr)})
	}
	if m.FocussizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "focussize"},
			Value: fmt.Sprintf("%v", *m.FocussizeAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShapeFill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "on" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OnAttr = &parsed
			continue
		}
		if attr.Name.Local == "focussize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FocussizeAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ShapeFill: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ShapeFill and its children
func (m *CT_ShapeFill) Validate() error {
	return m.ValidateWithPath("CT_ShapeFill")
}

// ValidateWithPath validates the CT_ShapeFill and its children, prefixing error messages with path
func (m *CT_ShapeFill) ValidateWithPath(path string) error {

	return nil
}