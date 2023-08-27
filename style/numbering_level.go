package style

import "godocx/simpletype"

type NumberingLevel struct {
	Level      int //0 to 8
	Start      int //1
	Restart    int
	PStyle     string
	Suffix     string //tab|space|nothin
	Text       string
	Alignment  simpletype.Jc
	Left       int
	Hanging    int
	TabPos     int
	FontFamily string
	Hint       string //default|eastAsia|cs
	//Format //NumberFormat
}
