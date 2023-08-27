package style

type LineNumberingRestart string

const (
	LINE_NUMBERING_CONTINUOUS  LineNumberingRestart = "continuous"
	LINE_NUMBERING_NEW_PAGE    LineNumberingRestart = "newPage"
	LINE_NUMBERING_NEW_SECTION LineNumberingRestart = "newSection"
)

type LineNumbering struct {
	Start     int //1
	Increment int //1
	Distance  float64
	Restart   LineNumberingRestart
}
