package element

import "godocx/style"

type FormField struct {
	Element        AbstractElement
	Text           string
	FontStyle      style.Font
	ParagraphStyle style.Paragraph
	Type           string // textinput|checkbox|dropdown.
	Name           string
	Default        string
	Value          string
}
