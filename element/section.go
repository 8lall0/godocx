package element

import "godocx/style"

type Section struct {
	container          string //Section
	Style              style.Section
	Headers            []Header //1-indexed
	Footers            []Footer //1-indexed
	FootNoteProperties struct{} //FootnoteProperties, è un complextype
}
