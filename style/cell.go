package style

import "godocx/simpletype"

type TextDirection string

const (
	TEXT_DIR_LRTB  TextDirection = "lrTb"  // Left to Right, Top to Bottom
	TEXT_DIR_TBRL  TextDirection = "tbRl"  //Top to Bottom, Right to Left
	TEXT_DIR_BTLR  TextDirection = "btLr"  //Bottom to Top, Left to Right.
	TEXT_DIR_LRTBV TextDirection = "lrTbV" //Left to Right, Top to Bottom Rotated
	TEXT_DIR_TBRLV TextDirection = "tbRlV" //Top to Bottom, Right to Left Rotated
	TEXT_DIR_TBLRV TextDirection = "tbLrV" //Top to Bottom, Left to Right Rotated
)

type RowSpan string

const (
	VMERGE_RESTART  RowSpan = "restart"
	VMERGE_CONTINUE RowSpan = "continue"
)

// TODO vedi dove si usa
const DEFAULT_BORDER_COLOR = "000000"

type Cell struct {
	VAlign        string //(top, center, both, bottom)
	TextDirection TextDirection
	GridSpan      int
	VMerge        RowSpan //restart, continue
	Width         int
	Unit          simpletype.TblWidth //TblWidth::TWIP
	Shading       Shading
}
