package style

import (
	"godocx/complextype"
	"godocx/simpletype"
)

type TableLayout string

const (
	LAYOUT_AUTO  TableLayout = "autofit"
	LAYOUT_FIXED TableLayout = "fixed"
)

type Table struct {
	IsFirstRow    bool
	FirstRowStyle *Table
	CellMargin    struct {
		Top    int
		Left   int
		Right  int
		Bottom int
	}
	BorderInside struct {
		HSize        int
		HInsideColor string
		VSize        int
		VInsideColor string
	}
	Shading      Shading
	Alignment    string
	Width        float64
	Unit         simpletype.TblWidth // AUTO
	CellSpacing  float64
	Layout       TableLayout //LAYOUT_AUTO
	Position     TablePosition
	ColumnWidths []int
	BidiVisual   bool
	Indent       complextype.TblWidth
}

//TODO controlla il costruttore ma mi sa che sti coglioni hanno fatto solo casino
