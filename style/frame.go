package style

import (
	"fmt"
	"github.com/shabbyrobe/xmlwriter"
	"godocx/simpletype"
	"math"
)

type UnitLength string

const (
	UNIT_PT UnitLength = "pt" // Mostly for shapes
	UNIT_PX UnitLength = "px"
)

type Positioning string

const (
	POS_ABSOLUTE Positioning = "absolute"
	POS_RELATIVE Positioning = "relative"
)

type HorVerValue string

const (
	POS_CENTER  HorVerValue = "center"
	POS_LEFT    HorVerValue = "left"
	POS_RIGHT   HorVerValue = "right"
	POS_TOP     HorVerValue = "top"
	POS_BOTTOM  HorVerValue = "bottom"
	POS_INSIDE  HorVerValue = "inside"
	POS_OUTSIDE HorVerValue = "outside"
)

type PositioningRelative string

const (
	POS_RELTO_MARGIN  PositioningRelative = "margin"
	POS_RELTO_PAGE    PositioningRelative = "page"
	POS_RELTO_COLUMN  PositioningRelative = "column"             // horizontal only
	POS_RELTO_CHAR    PositioningRelative = "char"               // horizontal only
	POS_RELTO_TEXT    PositioningRelative = "text"               // vertical only
	POS_RELTO_LINE    PositioningRelative = "line"               // vertical only
	POS_RELTO_LMARGIN PositioningRelative = "left-margin-area"   // horizontal only
	POS_RELTO_RMARGIN PositioningRelative = "right-margin-area"  // horizontal only
	POS_RELTO_TMARGIN PositioningRelative = "top-margin-area"    // vertical only
	POS_RELTO_BMARGIN PositioningRelative = "bottom-margin-area" // vertical only
	POS_RELTO_IMARGIN PositioningRelative = "inner-margin-area"
	POS_RELTO_OMARGIN PositioningRelative = "outer-margin-area"
)

type WrapType string

const (
	WRAP_INLINE    WrapType = "inline"
	WRAP_SQUARE    WrapType = "square"
	WRAP_TIGHT     WrapType = "tight"
	WRAP_THROUGH   WrapType = "through"
	WRAP_TOPBOTTOM WrapType = "topAndBottom"
	WRAP_BEHIND    WrapType = "behind"
	WRAP_INFRONT   WrapType = "infront"
)

type Frame struct {
	Alignment          simpletype.Jc
	Unit               UnitLength
	Width              float64
	Height             float64
	Left               float64
	Top                float64
	Pos                Positioning
	HPos               HorVerValue
	HPosRelTo          PositioningRelative
	VPos               HorVerValue
	VPosRelTo          PositioningRelative
	Wrap               WrapType
	WrapDistanceTop    float64
	WrapDistanceBottom float64
	WrapDistanceLeft   float64
	WrapDistanceRight  float64
	Position           int
}

func (f *Frame) writeWrap(w *xmlwriter.Writer) {
	if f.Wrap == "" {
		return
	}

	w.StartElem(xmlwriter.Elem{Name: "w10:wrap"})
	w.WriteAttr(xmlwriter.Attr{Name: "type", Value: string(f.Wrap)})

	if f.Pos == POS_ABSOLUTE {
		w.WriteAttr(
			xmlwriter.Attr{Name: "anchorx", Value: "page"},
			xmlwriter.Attr{Name: "anchory", Value: "page"},
		)
	} else if f.Pos == POS_RELATIVE {
		if f.HPosRelTo == POS_RELTO_MARGIN ||
			f.HPosRelTo == POS_RELTO_TMARGIN ||
			f.HPosRelTo == POS_RELTO_LMARGIN {
			w.WriteAttr(xmlwriter.Attr{Name: "anchorx", Value: "margin"})
		} else if f.HPosRelTo == POS_RELTO_PAGE ||
			f.HPosRelTo == POS_RELTO_BMARGIN ||
			f.HPosRelTo == POS_RELTO_RMARGIN {
			w.WriteAttr(xmlwriter.Attr{Name: "anchorx", Value: "page"})
		}

		if f.VPosRelTo == POS_RELTO_MARGIN ||
			f.VPosRelTo == POS_RELTO_TMARGIN ||
			f.VPosRelTo == POS_RELTO_LMARGIN {
			w.WriteAttr(xmlwriter.Attr{Name: "anchory", Value: "margin"})
		} else if f.VPosRelTo == POS_RELTO_PAGE ||
			f.VPosRelTo == POS_RELTO_BMARGIN ||
			f.VPosRelTo == POS_RELTO_RMARGIN {
			w.WriteAttr(xmlwriter.Attr{Name: "anchory", Value: "page"})
		}

		w.EndElem("w10:wrap")
	}
}

func setStyles(key string, value float64, unit string) string {
	if value == 0 {
		return ""
	}

	return fmt.Sprintf("%s:%.2f%s;", key, value, unit)
}

func (f *Frame) Write(w *xmlwriter.Writer) error {
	var err error

	var styleAttrs string
	unit := string(f.Unit)

	styleAttrs += setStyles("width", f.Width, unit)
	styleAttrs += setStyles("height", f.Height, unit)
	styleAttrs += setStyles("margin-left", f.Left, unit)
	styleAttrs += setStyles("margin-top", f.Top, unit)
	styleAttrs += setStyles("mso-wrap-distance-top", f.WrapDistanceTop, unit)
	styleAttrs += setStyles("mso-wrap-distance-bottom", f.WrapDistanceBottom, unit)
	styleAttrs += setStyles("mso-wrap-distance-left", f.WrapDistanceLeft, unit)
	styleAttrs += setStyles("mso-wrap-distance-right", f.WrapDistanceRight, unit)

	if f.Position != 0 {
		styleAttrs += fmt.Sprintf("position:%d;", f.Position)
	}
	if f.HPos != "" {
		styleAttrs += fmt.Sprintf("mso-position-horizontal:%s;", f.HPos)
	}
	if f.VPos != "" {
		styleAttrs += fmt.Sprintf("mso-position-vertical:%s;", f.VPos)
	}
	if f.HPosRelTo != "" {
		styleAttrs += fmt.Sprintf("mso-position-horizontal-relative:%s;", f.HPosRelTo)
	}
	if f.VPosRelTo != "" {
		styleAttrs += fmt.Sprintf("mso-position-vertical-relative:%s;", f.VPosRelTo)
	}

	if f.Wrap == WRAP_INFRONT {
		styleAttrs += fmt.Sprintf("z-index:%d;", math.MaxInt)
	} else if f.Wrap == WRAP_BEHIND {
		styleAttrs += fmt.Sprintf("z-index:%d;", -math.MaxInt)
	}

	err = w.WriteAttr(xmlwriter.Attr{Name: "style", Value: styleAttrs})

	f.writeWrap(w)

	return err
}

type ImageFrame Frame // this works also with images!!!
