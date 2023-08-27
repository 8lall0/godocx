package element

import "godocx/style"

type Link struct {
	Internal       bool
	MediaRelation  bool
	ParagraphStyle style.Paragraph
	FontStyle      style.Font
	Source         string
	Text           string
}
