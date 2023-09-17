package style

import (
	"fmt"
	"github.com/shabbyrobe/xmlwriter"
)

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
	Shading             *Shading
	Rtl                 bool
	NoProof             bool
	Language            *Language
	Hidden              bool
	Position            int
	Paragraph           Paragraph
	Inline              bool
}

func writeOnIf(w *xmlwriter.Writer, condition bool, tagName string) {
	var value string
	if condition {
		value = "1"
	} else {
		value = "0"
	}

	w.WriteElem(xmlwriter.Elem{Name: tagName, Attrs: []xmlwriter.Attr{{Name: "w:val", Value: value}}})
}

func (f *Font) Write(w *xmlwriter.Writer) error {
	if f.Inline && f.Name != "" {
		w.StartElem(xmlwriter.Elem{Name: "w:rPr"})
		w.StartElem(xmlwriter.Elem{Name: "w:rStyle"})
		w.WriteAttr(xmlwriter.Attr{Name: "w:val", Value: f.Name})
		w.EndElem("w:rStyle")
		if f.Rtl {
			w.WriteElem(xmlwriter.Elem{Name: "w:rtl"})
		}
		w.EndElem("w:rPr")
	} else {
		f.writeStyle(w)
	}

	return nil
}
func (f *Font) writeStyle(w *xmlwriter.Writer) error {
	w.StartElem(xmlwriter.Elem{Name: "w:rPr"})

	if f.Inline && f.Name != "" {
		w.WriteElem(xmlwriter.Elem{Name: "w:rStyle", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: f.Name}}})
	}

	// Font Name/Family
	if f.Name != "" {
		w.StartElem(xmlwriter.Elem{Name: "w:rFonts"})
		w.WriteAttr(
			xmlwriter.Attr{Name: "w:ascii", Value: f.Name},
			xmlwriter.Attr{Name: "w:hAnsi", Value: f.Name},
			xmlwriter.Attr{Name: "w:eastAsia", Value: f.Name},
			xmlwriter.Attr{Name: "w:cs", Value: f.Name},
		)
		if f.Hint != "" {
			w.WriteAttr(xmlwriter.Attr{Name: "w:hint", Value: f.Hint})
		}
		w.EndElem("w:rFonts")
	}

	if f.Language != nil {
		w.StartElem(xmlwriter.Elem{Name: "w:lang"})
		w.WriteAttr(
			xmlwriter.Attr{Name: "w:val", Value: f.Language.Latin},
			xmlwriter.Attr{Name: "w:eastAsia", Value: f.Language.EastAsia},
		)

		if f.Language.Bidirectional != "" {
			w.WriteAttr(xmlwriter.Attr{Name: "w:bidi", Value: f.Language.Bidirectional})
		} else if f.Rtl && f.Language.Bidirectional == "" && f.Language.Latin != "" {
			w.WriteAttr(xmlwriter.Attr{Name: "w:bidi", Value: f.Language.Latin})
		}
		w.EndElem("w:lang")
	}

	// Color
	w.StartElem(xmlwriter.Elem{Name: "w:color"})
	w.WriteAttr(xmlwriter.Attr{Name: "w:val", Value: f.Color})
	w.EndElem("w:color")

	// Size
	if f.Size != 0 {
		w.WriteElem(xmlwriter.Elem{Name: "w:sz", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: fmt.Sprintf("%f", f.Size)}}})
		w.WriteElem(xmlwriter.Elem{Name: "w:szCs", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: fmt.Sprintf("%f", f.Size)}}})
	}

	// Bold, italic
	writeOnIf(w, f.Bold, "w:b")
	writeOnIf(w, f.Bold, "w:bCs")
	writeOnIf(w, f.Italic, "w:i")
	writeOnIf(w, f.Italic, "w:iCs")

	// Strikethrough, double strikethrough
	writeOnIf(w, f.Strikethrough, "w:strike")
	writeOnIf(w, f.DoubleStrikethrough, "w:dstrike")

	// Small caps, all caps
	writeOnIf(w, f.SmallCaps, "w:smallCaps")
	writeOnIf(w, f.AllCaps, "w:caps")

	//Hidden text
	writeOnIf(w, f.Hidden, "w:vanish")

	// Underline
	if f.Underline != UNDERLINE_NONE {
		w.WriteElem(xmlwriter.Elem{Name: "w:u", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: string(f.Underline)}}})
	}

	if f.FgColor != "" {
		w.WriteElem(xmlwriter.Elem{Name: "w:highlight", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: string(f.FgColor)}}})
	}

	// Superscript/subscript
	if f.SuperScript {
		w.WriteElem(xmlwriter.Elem{Name: "w:vertAlign", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: "superscript"}}})
	}
	if f.SubScript {
		w.WriteElem(xmlwriter.Elem{Name: "w:vertAlign", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: "subscript"}}})
	}

	// Spacing
	if f.Scale != 0 {
		w.WriteElem(
			xmlwriter.Elem{Name: "w:w", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: fmt.Sprintf("%d", f.Scale)}}},
		)
	}
	if f.Spacing != 0 {
		w.WriteElem(
			xmlwriter.Elem{Name: "w:spacing", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: fmt.Sprintf("%f", f.Spacing)}}},
		)
	}
	if f.Kerning != 0 {
		w.WriteElem(
			xmlwriter.Elem{Name: "w:kern", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: fmt.Sprintf("%f", f.Kerning*2)}}},
		)
	}

	// noProof
	writeOnIf(w, f.NoProof, "w:noProof")

	if f.Shading != nil {
		// f.Shading.Write(w)
	}
	// RTL
	if f.Inline && f.Name == "" && f.Rtl {
		w.WriteElem(xmlwriter.Elem{Name: "w:rtl"})
	}

	// Position
	if f.Position != 0 {
		w.StartElem(xmlwriter.Elem{Name: "w:position"})
		w.WriteAttr(xmlwriter.Attr{Name: "w:val", Value: fmt.Sprintf("%d", f.Position)})
		w.EndElem("w:position")
	}

	w.EndElem("w:rPr")

	return nil
}

func (f *Font) WriteAlignment(w *xmlwriter.Writer) error {
	return nil
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
