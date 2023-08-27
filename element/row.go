package element

import (
	"godocx/container"
	"godocx/style"
)

type Row struct {
	Height int
	Style  style.Row
	Cells  []container.Cell
}

func (r *Row) AddCell() {
	//$cell = new Cell($width, $style);
	//$cell->setParentContainer($this);
	//$this->cells[] = $cell;
}
