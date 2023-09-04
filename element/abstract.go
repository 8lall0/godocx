package element

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
	ElementIndex int
	ElementId    string
	WithoutP     bool
}

func (a *AbsEl) foo() {

}

// Orco zio questa sarà difficile...
