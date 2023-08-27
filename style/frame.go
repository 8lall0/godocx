package style

import "godocx/simpletype"

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
	Leftmost           float64
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

type ImageFrame Frame // this works also with images!!!
