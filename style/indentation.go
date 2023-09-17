package style

import "github.com/shabbyrobe/xmlwriter"

type Indentation struct {
	Left      float64
	Right     float64
	FirstLine float64
	Hanging   float64
}

func (i Indentation) Write(w *xmlwriter.Writer) error {
	return nil
}
