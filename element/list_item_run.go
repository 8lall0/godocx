package element

import "godocx/style"

type ListItemRun struct {
	TextRun
	container string //ListItemRun
	Style     style.ListItem
	Depth     int
}

// Piuttosto duplica!!!
