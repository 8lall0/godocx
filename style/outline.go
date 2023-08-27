package style

type LineStyle string

const (
	LINE_SINGLE             LineStyle = "single"
	LINE_THIN_THIN          LineStyle = "thinThin"
	LINE_THIN_THICK         LineStyle = "thinThick"
	LINE_THICK_THIN         LineStyle = "thickThin"
	LINE_THICK_BETWEEN_THIN LineStyle = "thickBetweenThin"
)

type EndcapStyle string

const (
	ENDCAP_FLAT   EndcapStyle = "flat"
	ENDCAP_SQUARE EndcapStyle = "square"
	ENDCAP_ROUND  EndcapStyle = "round"
)

type ArrowStyle string

const (
	ARROW_NONE    ArrowStyle = "none"
	ARROW_BLOCK   ArrowStyle = "block"
	ARROW_CLASSIC ArrowStyle = "classic"
	ARROW_OVAL    ArrowStyle = "oval"
	ARROW_DIAMOND ArrowStyle = "diamond"
	ARROW_OPEN    ArrowStyle = "open"
)

type Outline struct {
	Unit       string
	Weight     float64
	Color      string
	Dash       string
	LineStyle  LineStyle
	EndCap     EndcapStyle
	StartArrow ArrowStyle
	EndArrow   ArrowStyle
}
