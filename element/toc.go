package element

import "godocx/style"

// TODO controlla gli AbstractElement

type Toc struct {
	Style     style.Toc
	FontStyle style.Font
	MinDepth  int // 1
	MaxDepth  int // 9
}
