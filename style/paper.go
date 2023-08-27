package style

type Size struct {
	W    float64
	H    float64
	Unit string
}
type PaperSizes map[string]Size

var sizes PaperSizes = PaperSizes{
	"A3":     {297, 420, "mm"},
	"A4":     {210, 297, "mm"},
	"A5":     {148, 210, "mm"},
	"B5":     {176, 250, "mm"},
	"Folio":  {8.5, 13, "in"},
	"Legal":  {8.5, 14, "in"},
	"Letter": {8.5, 11, "in"},
}

type Paper struct {
	Size   Size
	Width  float64
	Height float64
}

// TODO check function SetSize
