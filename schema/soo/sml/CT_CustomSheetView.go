// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/qifengzhang007/gooxml"
	"github.com/qifengzhang007/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_CustomSheetView struct {
	// GUID
	GuidAttr string
	// Print Scale
	ScaleAttr *uint32
	// Color Id
	ColorIdAttr *uint32
	// Show Page Breaks
	ShowPageBreaksAttr *bool
	// Show Formulas
	ShowFormulasAttr *bool
	// Show Grid Lines
	ShowGridLinesAttr *bool
	// Show Headers
	ShowRowColAttr *bool
	// Show Outline Symbols
	OutlineSymbolsAttr *bool
	// Show Zero Values
	ZeroValuesAttr *bool
	// Fit To Page
	FitToPageAttr *bool
	// Print Area Defined
	PrintAreaAttr *bool
	// Filtered List
	FilterAttr *bool
	// Show AutoFitler Drop Down Controls
	ShowAutoFilterAttr *bool
	// Hidden Rows
	HiddenRowsAttr *bool
	// Hidden Columns
	HiddenColumnsAttr *bool
	// Visible State
	StateAttr ST_SheetState
	// Filter
	FilterUniqueAttr *bool
	// View Type
	ViewAttr ST_SheetViewType
	// Show Ruler
	ShowRulerAttr *bool
	// Top Left Visible Cell
	TopLeftCellAttr *string
	// Pane Split Information
	Pane *CT_Pane
	// Selection
	Selection *CT_Selection
	// Horizontal Page Breaks
	RowBreaks *CT_PageBreak
	// Vertical Page Breaks
	ColBreaks *CT_PageBreak
	// Page Margins
	PageMargins *CT_PageMargins
	// Print Options
	PrintOptions *CT_PrintOptions
	// Page Setup Settings
	PageSetup *CT_PageSetup
	// Header Footer Settings
	HeaderFooter *CT_HeaderFooter
	// AutoFilter Settings
	AutoFilter *CT_AutoFilter
	ExtLst     *CT_ExtensionList
}

func NewCT_CustomSheetView() *CT_CustomSheetView {
	ret := &CT_CustomSheetView{}
	ret.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_CustomSheetView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
		Value: fmt.Sprintf("%v", m.GuidAttr)})
	if m.ScaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "scale"},
			Value: fmt.Sprintf("%v", *m.ScaleAttr)})
	}
	if m.ColorIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "colorId"},
			Value: fmt.Sprintf("%v", *m.ColorIdAttr)})
	}
	if m.ShowPageBreaksAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showPageBreaks"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowPageBreaksAttr))})
	}
	if m.ShowFormulasAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showFormulas"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowFormulasAttr))})
	}
	if m.ShowGridLinesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showGridLines"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowGridLinesAttr))})
	}
	if m.ShowRowColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showRowCol"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowRowColAttr))})
	}
	if m.OutlineSymbolsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outlineSymbols"},
			Value: fmt.Sprintf("%d", b2i(*m.OutlineSymbolsAttr))})
	}
	if m.ZeroValuesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zeroValues"},
			Value: fmt.Sprintf("%d", b2i(*m.ZeroValuesAttr))})
	}
	if m.FitToPageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fitToPage"},
			Value: fmt.Sprintf("%d", b2i(*m.FitToPageAttr))})
	}
	if m.PrintAreaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "printArea"},
			Value: fmt.Sprintf("%d", b2i(*m.PrintAreaAttr))})
	}
	if m.FilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "filter"},
			Value: fmt.Sprintf("%d", b2i(*m.FilterAttr))})
	}
	if m.ShowAutoFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showAutoFilter"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowAutoFilterAttr))})
	}
	if m.HiddenRowsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hiddenRows"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenRowsAttr))})
	}
	if m.HiddenColumnsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hiddenColumns"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenColumnsAttr))})
	}
	if m.StateAttr != ST_SheetStateUnset {
		attr, err := m.StateAttr.MarshalXMLAttr(xml.Name{Local: "state"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FilterUniqueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "filterUnique"},
			Value: fmt.Sprintf("%d", b2i(*m.FilterUniqueAttr))})
	}
	if m.ViewAttr != ST_SheetViewTypeUnset {
		attr, err := m.ViewAttr.MarshalXMLAttr(xml.Name{Local: "view"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShowRulerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showRuler"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowRulerAttr))})
	}
	if m.TopLeftCellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "topLeftCell"},
			Value: fmt.Sprintf("%v", *m.TopLeftCellAttr)})
	}
	e.EncodeToken(start)
	if m.Pane != nil {
		sepane := xml.StartElement{Name: xml.Name{Local: "ma:pane"}}
		e.EncodeElement(m.Pane, sepane)
	}
	if m.Selection != nil {
		seselection := xml.StartElement{Name: xml.Name{Local: "ma:selection"}}
		e.EncodeElement(m.Selection, seselection)
	}
	if m.RowBreaks != nil {
		serowBreaks := xml.StartElement{Name: xml.Name{Local: "ma:rowBreaks"}}
		e.EncodeElement(m.RowBreaks, serowBreaks)
	}
	if m.ColBreaks != nil {
		secolBreaks := xml.StartElement{Name: xml.Name{Local: "ma:colBreaks"}}
		e.EncodeElement(m.ColBreaks, secolBreaks)
	}
	if m.PageMargins != nil {
		sepageMargins := xml.StartElement{Name: xml.Name{Local: "ma:pageMargins"}}
		e.EncodeElement(m.PageMargins, sepageMargins)
	}
	if m.PrintOptions != nil {
		seprintOptions := xml.StartElement{Name: xml.Name{Local: "ma:printOptions"}}
		e.EncodeElement(m.PrintOptions, seprintOptions)
	}
	if m.PageSetup != nil {
		sepageSetup := xml.StartElement{Name: xml.Name{Local: "ma:pageSetup"}}
		e.EncodeElement(m.PageSetup, sepageSetup)
	}
	if m.HeaderFooter != nil {
		seheaderFooter := xml.StartElement{Name: xml.Name{Local: "ma:headerFooter"}}
		e.EncodeElement(m.HeaderFooter, seheaderFooter)
	}
	if m.AutoFilter != nil {
		seautoFilter := xml.StartElement{Name: xml.Name{Local: "ma:autoFilter"}}
		e.EncodeElement(m.AutoFilter, seautoFilter)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomSheetView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "state" {
			m.StateAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = parsed
			continue
		}
		if attr.Name.Local == "filterUnique" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FilterUniqueAttr = &parsed
			continue
		}
		if attr.Name.Local == "colorId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ColorIdAttr = &pt
			continue
		}
		if attr.Name.Local == "view" {
			m.ViewAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "showFormulas" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowFormulasAttr = &parsed
			continue
		}
		if attr.Name.Local == "showRuler" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowRulerAttr = &parsed
			continue
		}
		if attr.Name.Local == "showRowCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowRowColAttr = &parsed
			continue
		}
		if attr.Name.Local == "topLeftCell" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TopLeftCellAttr = &parsed
			continue
		}
		if attr.Name.Local == "zeroValues" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ZeroValuesAttr = &parsed
			continue
		}
		if attr.Name.Local == "hiddenColumns" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenColumnsAttr = &parsed
			continue
		}
		if attr.Name.Local == "outlineSymbols" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineSymbolsAttr = &parsed
			continue
		}
		if attr.Name.Local == "filter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FilterAttr = &parsed
			continue
		}
		if attr.Name.Local == "scale" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ScaleAttr = &pt
			continue
		}
		if attr.Name.Local == "showGridLines" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowGridLinesAttr = &parsed
			continue
		}
		if attr.Name.Local == "hiddenRows" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenRowsAttr = &parsed
			continue
		}
		if attr.Name.Local == "fitToPage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FitToPageAttr = &parsed
			continue
		}
		if attr.Name.Local == "printArea" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PrintAreaAttr = &parsed
			continue
		}
		if attr.Name.Local == "showAutoFilter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAutoFilterAttr = &parsed
			continue
		}
		if attr.Name.Local == "showPageBreaks" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowPageBreaksAttr = &parsed
			continue
		}
	}
lCT_CustomSheetView:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pane"}:
				m.Pane = NewCT_Pane()
				if err := d.DecodeElement(m.Pane, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "selection"}:
				m.Selection = NewCT_Selection()
				if err := d.DecodeElement(m.Selection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rowBreaks"}:
				m.RowBreaks = NewCT_PageBreak()
				if err := d.DecodeElement(m.RowBreaks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "colBreaks"}:
				m.ColBreaks = NewCT_PageBreak()
				if err := d.DecodeElement(m.ColBreaks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pageMargins"}:
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "printOptions"}:
				m.PrintOptions = NewCT_PrintOptions()
				if err := d.DecodeElement(m.PrintOptions, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pageSetup"}:
				m.PageSetup = NewCT_PageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "headerFooter"}:
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "autoFilter"}:
				m.AutoFilter = NewCT_AutoFilter()
				if err := d.DecodeElement(m.AutoFilter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_CustomSheetView %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomSheetView
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomSheetView and its children
func (m *CT_CustomSheetView) Validate() error {
	return m.ValidateWithPath("CT_CustomSheetView")
}

// ValidateWithPath validates the CT_CustomSheetView and its children, prefixing error messages with path
func (m *CT_CustomSheetView) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.GuidAttr) {
		return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.GuidAttr)
	}
	if err := m.StateAttr.ValidateWithPath(path + "/StateAttr"); err != nil {
		return err
	}
	if err := m.ViewAttr.ValidateWithPath(path + "/ViewAttr"); err != nil {
		return err
	}
	if m.Pane != nil {
		if err := m.Pane.ValidateWithPath(path + "/Pane"); err != nil {
			return err
		}
	}
	if m.Selection != nil {
		if err := m.Selection.ValidateWithPath(path + "/Selection"); err != nil {
			return err
		}
	}
	if m.RowBreaks != nil {
		if err := m.RowBreaks.ValidateWithPath(path + "/RowBreaks"); err != nil {
			return err
		}
	}
	if m.ColBreaks != nil {
		if err := m.ColBreaks.ValidateWithPath(path + "/ColBreaks"); err != nil {
			return err
		}
	}
	if m.PageMargins != nil {
		if err := m.PageMargins.ValidateWithPath(path + "/PageMargins"); err != nil {
			return err
		}
	}
	if m.PrintOptions != nil {
		if err := m.PrintOptions.ValidateWithPath(path + "/PrintOptions"); err != nil {
			return err
		}
	}
	if m.PageSetup != nil {
		if err := m.PageSetup.ValidateWithPath(path + "/PageSetup"); err != nil {
			return err
		}
	}
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.ValidateWithPath(path + "/HeaderFooter"); err != nil {
			return err
		}
	}
	if m.AutoFilter != nil {
		if err := m.AutoFilter.ValidateWithPath(path + "/AutoFilter"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
