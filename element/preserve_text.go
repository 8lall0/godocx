package element

import "godocx/style"

type PreserveText struct {
	Element        AbstractElement
	Text           string
	FontStyle      style.Font
	ParagraphStyle style.Paragraph
}
