package style

import "godocx/simpletype"

type Paragraph struct {
	Aliases             string // ['line-height' => 'lineHeight', 'line-spacing' => 'spacing'] TODO
	BasedOn             string // Normal
	Next                string
	Alignment           simpletype.Jc
	Indentation         Indentation
	Spacing             Spacing
	LineHeight          float64 //240
	WidowControl        bool    //true
	KeepNext            bool
	KeepLines           bool
	PageBreakBefore     bool
	NumStyle            string
	Level               int
	Tabs                []Tab
	Shading             Shading
	ContextualSpacing   bool
	BiDi                bool
	TextAlignment       string
	SuppressAutoHyphens bool
}

//TODO setLineHeight
