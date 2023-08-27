package simpletype

type borderType string

const (
	Single                 borderType = "single"                 //A single line
	DashDotStroked         borderType = "dashDotStroked"         //A line with a series of alternating thin and thick strokes
	Dashed                 borderType = "dashed"                 //A dashed line
	DashSmallGap           borderType = "dashSmallGap"           //A dashed line with small gaps
	DotDash                borderType = "dotDash"                //A line with alternating dots and dashes
	DotDotDash             borderType = "dotDotDash"             //A line with a repeating dot - dot - dash sequence
	Dotted                 borderType = "dotted"                 //A dotted line
	Double                 borderType = "double"                 //A double line
	DoubleWave             borderType = "doubleWave"             //A double wavy line
	Inset                  borderType = "inset"                  //An inset set of lines
	Nil                    borderType = "nil"                    //No border
	None                   borderType = "none"                   //No border
	Outset                 borderType = "outset"                 //An outset set of lines
	Thick                  borderType = "thick"                  //A single line
	ThickThinLargeGap      borderType = "thickThinLargeGap"      //A thick line contained within a thin line with a large-sized intermediate gap
	ThickThinMediumGap     borderType = "thickThinMediumGap"     //A thick line contained within a thin line with a medium-sized intermediate gap
	ThickThinSmallGap      borderType = "thickThinSmallGap"      //A thick line contained within a thin line with a small intermediate gap
	ThinThickLargeGap      borderType = "thinThickLargeGap"      //A thin line contained within a thick line with a large-sized intermediate gap
	ThinThickMediumGap     borderType = "thinThickMediumGap"     //A thick line contained within a thin line with a medium-sized intermediate gap
	ThinThickSmallGap      borderType = "thinThickSmallGap"      //A thick line contained within a thin line with a small intermediate gap
	ThinThickThinLargeGap  borderType = "thinThickThinLargeGap"  //A thin-thick-thin line with a large gap
	ThinThickThinMediumGap borderType = "thinThickThinMediumGap" //A thin-thick-thin line with a medium gap
	ThinThickThinSmallGap  borderType = "thinThickThinSmallGap"  //A thin-thick-thin line with a small gap
	ThreeDEmboss           borderType = "threeDEmboss"           //A three-staged gradient line, getting darker towards the paragraph
	ThreeDEngrave          borderType = "threeDEngrave"          //A three-staged gradient like, getting darker away from the paragraph
	Triple                 borderType = "triple"                 //A triple line
	Wave                   borderType = "wave"                   //A wavy line
)
