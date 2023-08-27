package element

import "godocx/style"

type Text struct {
	Element        AbstractElement
	Text           string
	FontStyle      style.Font
	ParagraphStyle style.Paragraph
}
