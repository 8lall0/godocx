package style

import "github.com/shabbyrobe/xmlwriter"

type Spacing struct {
	Before   float64
	After    float64
	Line     float64
	LineRule string //lineRule = LineSpacingRule::AUTO
}

func (s Spacing) Write(w *xmlwriter.Writer) error {
	return nil
}
