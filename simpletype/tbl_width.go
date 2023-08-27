package simpletype

type TblWidth string

const (
	TblWidthNIL     TblWidth = "nil"
	TblWidthAUTO    TblWidth = "auto" //Automatically Determined Width
	TblWidthPERCENT TblWidth = "pct"  //Width in Fiftieths of a Percent
	TblWidthTWIP    TblWidth = "dxa"  //Width in Twentieths of a Point

)
