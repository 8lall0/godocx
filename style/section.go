package style

import "godocx/simpletype"

//TODO devi riorganizzare cazzo

type Orientation string

const (
	ORIENTATION_PORTRAIT  Orientation = "portrait"
	ORIENTATION_LANDSCAPE Orientation = "landscape"
)

type DefaultAnts string //must be int/float
const (
	DEFAULT_WIDTH          = 11905.511811024 // In twips.
	DEFAULT_HEIGHT         = 16837.79527559  // In twips.
	DEFAULT_MARGIN         = 1440            // In twips.
	DEFAULT_GUTTER         = 0               // In twips.
	DEFAULT_HEADER_HEIGHT  = 720             // In twips.
	DEFAULT_FOOTER_HEIGHT  = 720             // In twips.
	DEFAULT_COLUMN_COUNT   = 1
	DEFAULT_COLUMN_SPACING = 720
)

/**
 * Section break type.
 *
 * Options:
 * - nextPage: Next page section break
 * - nextColumn: Column section break
 * - continuous: Continuous section break
 * - evenPage: Even page section break
 * - oddPage: Odd page section break
 *
 * @var ?string
 */

type Section struct {
	Border      Border
	Orientation Orientation // Portrait
	Paper       Paper
	PageSizeW   float64
	PageSizeH   float64
	Margin      struct {
		Top    float64
		Left   float64
		Right  float64
		Bottom float64
	}
	Gutter             float64
	HeaderHeight       float64
	FooterHeight       float64
	PageNumberingStart int
	ColsNum            int
	ColsSpace          float64
	BreakType          string
	LineNumbering      LineNumbering
	VAlign             simpletype.VerticalJc
}
