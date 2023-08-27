package element

import "godocx/style"

type SdtType string

const (
	SdtPlainText    SdtType = "plainText"
	SdtComboBox     SdtType = "comboBox"
	SdtDropDownList SdtType = "dropDownList"
	SdtDate         SdtType = "date"
)

type Sdt struct {
	Element        AbstractElement
	Text           string
	FontStyle      style.Font
	ParagraphStyle style.Paragraph
	Type           SdtType
	Value          string
	ListItems      []struct{} // CheckBox/DropDown
	Alias          string
	Tag            string
}
