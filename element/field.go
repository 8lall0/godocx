package element

import "godocx/style"

type Field struct {
	Type       string
	Text       TextRun
	Properties []string // Array di cosa???
	Options    []string // Array di cosa???
	FontStyle  style.Font
}
