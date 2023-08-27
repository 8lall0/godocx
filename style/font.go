package style

type FontUnderlineType string

const (
	UNDERLINE_NONE            FontUnderlineType = "none"
	UNDERLINE_DASH            FontUnderlineType = "dash"
	UNDERLINE_DASHHEAVY       FontUnderlineType = "dashHeavy"
	UNDERLINE_DASHLONG        FontUnderlineType = "dashLong"
	UNDERLINE_DASHLONGHEAVY   FontUnderlineType = "dashLongHeavy"
	UNDERLINE_DOUBLE          FontUnderlineType = "dbl"
	UNDERLINE_DOTDASH         FontUnderlineType = "dotDash"
	UNDERLINE_DOTDASHHEAVY    FontUnderlineType = "dotDashHeavy"
	UNDERLINE_DOTDOTDASH      FontUnderlineType = "dotDotDash"
	UNDERLINE_DOTDOTDASHHEAVY FontUnderlineType = "dotDotDashHeavy"
	UNDERLINE_DOTTED          FontUnderlineType = "dotted"
	UNDERLINE_DOTTEDHEAVY     FontUnderlineType = "dottedHeavy"
	UNDERLINE_HEAVY           FontUnderlineType = "heavy"
	UNDERLINE_SINGLE          FontUnderlineType = "single"
	UNDERLINE_WAVY            FontUnderlineType = "wavy"
	UNDERLINE_WAVYDOUBLE      FontUnderlineType = "wavyDbl"
	UNDERLINE_WAVYHEAVY       FontUnderlineType = "wavyHeavy"
	UNDERLINE_WORDS           FontUnderlineType = "words"
)

type FontForegroundColor string

const (
	FGCOLOR_YELLOW      FontForegroundColor = "yellow"
	FGCOLOR_LIGHTGREEN  FontForegroundColor = "green"
	FGCOLOR_CYAN        FontForegroundColor = "cyan"
	FGCOLOR_MAGENTA     FontForegroundColor = "magenta"
	FGCOLOR_BLUE        FontForegroundColor = "blue"
	FGCOLOR_RED         FontForegroundColor = "red"
	FGCOLOR_DARKBLUE    FontForegroundColor = "darkBlue"
	FGCOLOR_DARKCYAN    FontForegroundColor = "darkCyan"
	FGCOLOR_DARKGREEN   FontForegroundColor = "darkGreen"
	FGCOLOR_DARKMAGENTA FontForegroundColor = "darkMagenta"
	FGCOLOR_DARKRED     FontForegroundColor = "darkRed"
	FGCOLOR_DARKYELLOW  FontForegroundColor = "darkYellow"
	FGCOLOR_DARKGRAY    FontForegroundColor = "darkGray"
	FGCOLOR_LIGHTGRAY   FontForegroundColor = "lightGray"
	FGCOLOR_BLACK       FontForegroundColor = "black"
)

var aliases = map[string]string{
	"line-height":    "lineHeight",
	"letter-spacing": "spacing",
}

type Font struct {
	Type                string
	Name                string
	Hint                string
	Size                float64
	Color               string
	Bold                bool
	Italic              bool
	Underline           FontUnderlineType //UNDERLINE_NONE
	SuperScript         bool
	SubScript           bool
	Strikethrough       bool
	DoubleStrikethrough bool
	SmallCaps           bool
	AllCaps             bool
	FgColor             FontForegroundColor
	Scale               int //0-600 (percent)
	Spacing             float64
	Kerning             float64
	Shading             Shading
	Rtl                 bool
	NoProof             bool
	Language            Language
	Hidden              bool
	Position            int
	//Paragraph Paragraph
}

/*
 public function getStyleValues()
    {
        $styles = [
            'name' => $this->getStyleName(),
            'basic' => [
                'name' => $this->getName(),
                'size' => $this->getSize(),
                'color' => $this->getColor(),
                'hint' => $this->getHint(),
            ],
            'style' => [
                'bold' => $this->isBold(),
                'italic' => $this->isItalic(),
                'underline' => $this->getUnderline(),
                'strike' => $this->isStrikethrough(),
                'dStrike' => $this->isDoubleStrikethrough(),
                'super' => $this->isSuperScript(),
                'sub' => $this->isSubScript(),
                'smallCaps' => $this->isSmallCaps(),
                'allCaps' => $this->isAllCaps(),
                'fgColor' => $this->getFgColor(),
                'hidden' => $this->isHidden(),
            ],
            'spacing' => [
                'scale' => $this->getScale(),
                'spacing' => $this->getSpacing(),
                'kerning' => $this->getKerning(),
                'position' => $this->getPosition(),
            ],
            'paragraph' => $this->getParagraph(),
            'rtl' => $this->isRTL(),
            'shading' => $this->getShading(),
            'lang' => $this->getLang(),
        ];

        return $styles;
    }
*/
