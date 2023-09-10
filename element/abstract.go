package element

import "github.com/shabbyrobe/xmlwriter"

type AbstractElement struct {
	SectionId          int
	DocPart            string // Section|Header|Footer|Footnote|Endnote
	DocPartId          int
	ElementIndex       int
	ElementId          string
	RelationId         int
	NestedLevel        int // 0 = Not in a table; 1 = in a table; 2 = in a table inside another table, etc.
	Parent             any // Porcodio questo no no no no no
	TrackChange        TrackChange
	ParentContainer    string   //Parent container type.
	MediaRelation      bool     // Has media relation flag; true for Link, Image, and Object.
	CollectionRelation bool     // Is part of collection; true for Title, Footnote, Endnote, Chart, and Comment.
	CommentRangeStart  struct{} // COMMENT
	CommentRangeEnd    struct{} // COMMENT
}

type AbsEl struct {
	ElementIndex      int
	ElementId         string
	WithoutP          bool
	CommentRangeStart struct{} // COMMENT
	CommentRangeEnd   struct{} // COMMENT
}

func (a *AbsEl) WriteCommentRangeStart(w *xmlwriter.Writer) error {
	err := w.StartElem(xmlwriter.Elem{Name: "w:commentRangeStart"})

	//$this->xmlWriter->writeElementBlock('w:commentRangeStart', ['w:id' => $comment->getElementId()]);
	/*public function writeElementBlock($element, $attributes, $value = null): void
	{
		$this->startElement($element);
		if (!is_array($attributes)) {
	$attributes = [$attributes => $value];
	}
		foreach ($attributes as $attribute => $value) {
		$this->writeAttribute($attribute, $value);
	}
		$this->endElement();
	}*/
	return err
}
