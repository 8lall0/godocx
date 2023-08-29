package complextype

import "godocx/simpletype"

type FootNoteRestart string

const (
	RESTART_NUMBER_CONTINUOUS   FootNoteRestart = "continuous"
	RESTART_NUMBER_EACH_SECTION FootNoteRestart = "eachSect"
	RESTART_NUMBER_EACH_PAGE    FootNoteRestart = "eachPage"
)

type FootNotePosition string

const (
	POSITION_PAGE_BOTTOM  FootNotePosition = "pageBottom"
	POSITION_BENEATH_TEXT FootNotePosition = "beneathText"
	POSITION_SECTION_END  FootNotePosition = "sectEnd"
	POSITION_DOC_END      FootNotePosition = "docEnd"
)

type FootNoteProperties struct {
	Position   FootNotePosition
	NumFmt     simpletype.NumberFormat
	NumStart   float64
	NumRestart FootNoteRestart
}
