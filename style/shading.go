package style

import "github.com/shabbyrobe/xmlwriter"

type ShadingPattern string

const (
	PATTERN_CLEAR   ShadingPattern = "clear"      // No pattern
	PATTERN_SOLID   ShadingPattern = "solid"      // 100% fill pattern
	PATTERN_HSTRIPE ShadingPattern = "horzStripe" // Horizontal stripe pattern
	PATTERN_VSTRIPE ShadingPattern = "vertStripe" // Vertical stripe pattern
	PATTERN_DSTRIPE ShadingPattern = "diagStripe" // Diagonal stripe pattern
	PATTERN_HCROSS  ShadingPattern = "horzCross"  // Horizontal cross pattern
	PATTERN_DCROSS  ShadingPattern = "diagCross"  // Diagonal cross pattern
)

type Shading struct {
	Pattern ShadingPattern //PAttern_Clear
	Color   string
	Fill    string
}

func (s Shading) Write(w *xmlwriter.Writer) error {
	return nil
}
