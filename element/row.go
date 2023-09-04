package element

import (
	"godocx/style"
)

type Row struct {
	Height int
	Style  style.Row
	Cells  []Cell
}

func (r *Row) AddCell() {
	//$cell = new Cell($width, $style);
	//$cell->setParentContainer($this);
	//$this->cells[] = $cell;
}
