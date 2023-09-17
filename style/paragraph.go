package style

import (
	"fmt"
	"github.com/shabbyrobe/xmlwriter"
	"godocx/simpletype"
)

/*
THE STUPIDEST THING EVER
  $styles = [
            'name' => $this->getStyleName(),
            'basedOn' => $this->getBasedOn(),
            'next' => $this->getNext(),
            'alignment' => $this->getAlignment(),
            'indentation' => $this->getIndentation(),
            'spacing' => $this->getSpace(),
            'pagination' => [
                'widowControl' => $this->hasWidowControl(),
                'keepNext' => $this->isKeepNext(),
                'keepLines' => $this->isKeepLines(),
                'pageBreak' => $this->hasPageBreakBefore(),
            ],
            'numbering' => [
                'style' => $this->getNumStyle(),
                'level' => $this->getNumLevel(),
            ],
            'tabs' => $this->getTabs(),
            'shading' => $this->getShading(),
            'contextualSpacing' => $this->hasContextualSpacing(),
            'bidi' => $this->isBidi(),
            'textAlignment' => $this->getTextAlignment(),
            'suppressAutoHyphens' => $this->hasSuppressAutoHyphens(),
        ];
*/

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
	Inline              bool
	WithoutPPR          bool
	Border              *Border
	AbstractStyle
}

func (p *Paragraph) Write(w *xmlwriter.Writer) error {
	if p.Inline && p.styleName != "" {
		if p.WithoutPPR == false {
			w.StartElem(xmlwriter.Elem{Name: "w:pPr"})
		}
		w.StartElem(xmlwriter.Elem{Name: "w:pStyle"})
		w.WriteAttr(xmlwriter.Attr{Name: "w:val", Value: p.styleName})
		w.EndElem("w:pStyle")
		if p.WithoutPPR == false {
			w.EndElem("w:pPr")
		}
	} else {
		p.writeStyle(w)
	}

	return nil
}

func (p *Paragraph) writeStyle(w *xmlwriter.Writer) error {
	//$styles = $style->getStyleValues();
	if p.WithoutPPR == false {
		w.StartElem(xmlwriter.Elem{Name: "w:pPr"})
	}

	// Style name
	if p.Inline && p.styleName != "" {
		w.WriteElem(xmlwriter.Elem{Name: "w:pStyle", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: p.styleName}}})
	}

	// Pagination
	if p.WidowControl == false {
		w.WriteElem(xmlwriter.Elem{Name: "w:widowControl", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: "0"}}})
	}
	if p.KeepNext {
		w.WriteElem(xmlwriter.Elem{Name: "w:keepNext", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: "1"}}})
	}
	if p.KeepNext {
		w.WriteElem(xmlwriter.Elem{Name: "w:keepLines", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: "1"}}})
	}
	if p.PageBreakBefore {
		w.WriteElem(xmlwriter.Elem{Name: "w:pageBreakBefore", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: "1"}}})
	}

	// Paragraph alignment
	if p.Alignment != "" {
		/*
			$paragraphAlignment = new ParagraphAlignment($styles['alignment']);
			$xmlWriter->startElement($paragraphAlignment->getName());
			foreach ($paragraphAlignment->getAttributes() as $attributeName => $attributeValue) {
				$xmlWriter->writeAttribute($attributeName, $attributeValue);
			}
			$xmlWriter->endElement();
		*/
	}

	//Right to left
	if p.BiDi {
		w.WriteElem(xmlwriter.Elem{Name: "w:bidi"})
	}

	//Paragraph contextualSpacing
	if p.ContextualSpacing {
		w.WriteElem(xmlwriter.Elem{Name: "w:contextualSpacing"})
	}

	//Paragraph textAlignment
	if p.TextAlignment != "" {
		w.WriteElem(xmlwriter.Elem{Name: "w:textAlignment", Attrs: []xmlwriter.Attr{{Name: "w:val", Value: p.TextAlignment}}})
	}

	// Hyphenation
	if p.SuppressAutoHyphens {
		w.WriteElem(xmlwriter.Elem{Name: "w:suppressAutoHyphens"})
	}

	// Child style: alignment, indentation, spacing, and shading
	p.Indentation.Write(w)
	p.Spacing.Write(w)
	p.Shading.Write(w)

	// Tabs
	p.writeTabs(w)
	// Numbering
	p.writeNumbering(w)

	// Border
	if p.Border != nil {
		w.StartElem(xmlwriter.Elem{Name: "w:pBdr"})
		//$styleWriter = new MarginBorder($xmlWriter);
		//$styleWriter->setSizes($style->getBorderSize());
		//$styleWriter->setStyles($style->getBorderStyle());
		//$styleWriter->setColors($style->getBorderColor());
		//$styleWriter->write();
		w.EndElem("w:pBdr")
	}

	if p.WithoutPPR == false {
		w.EndElem("w:pPr")
	}

	return nil
}

func (p *Paragraph) writeTabs(w *xmlwriter.Writer) error {
	if len(p.Tabs) == 0 {
		return nil
	}

	w.StartElem(xmlwriter.Elem{Name: "w:tabs"})
	for _, tab := range p.Tabs {
		tab.Write(w)
	}
	w.EndElem("w:tabs")

	return nil
}

func (p *Paragraph) writeNumbering(w *xmlwriter.Writer) error {
	if p.NumStyle == "" || p.Level == 0 {
		return nil
	}
	// TODO controlla se devo usare numbering
	w.StartElem(xmlwriter.Elem{Name: "w:numPr"})
	w.StartElem(xmlwriter.Elem{Name: "w:numId"})
	w.WriteAttr(xmlwriter.Attr{Name: "w:val", Value: fmt.Sprintf("%d", p.Index)})
	w.EndElem("w:numId")
	w.StartElem(xmlwriter.Elem{Name: "w:ilvl"})
	w.WriteAttr(xmlwriter.Attr{Name: "w:val", Value: fmt.Sprintf("%d", p.Level)})
	w.EndElem("w:ilvl")
	w.EndElem("w:numPr")

	w.StartElem(xmlwriter.Elem{Name: "w:outlineLvl"})
	w.WriteAttr(xmlwriter.Attr{Name: "w:val", Value: fmt.Sprintf("%d", p.Level)})
	w.EndElem("w:outlineLvl")

	return nil
}
