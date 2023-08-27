package element

import "godocx/style"

type Table struct {
	Element AbstractElement
	Style   style.Table
	Width   int
	Rows    []Row
}
