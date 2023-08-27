package style

type ChartOptions map[string]bool

type Chart struct {
	Width                 int  //1000000
	Height                int  //1000000
	Is3d                  bool //Is 3D; applies to pie, bar, line, area.
	Colors                []string
	Title                 string
	ShowLegend            bool
	LegendPosition        string //Possible values are "r", "t", "b", "l", "tr"
	CategoryLabelPosition string //nexTo, low, high
	ValueLabelPosition    string //nexTo, low, high
	CategoryAxisTitle     string
	ValueAxisTitle        string
	MajorTickmarkPos      string //Possible values are "in", "out", "cross", "none"
	ShowAxisLabels        bool
	ShowGridX             bool
	ShowGridY             bool
	Options               ChartOptions
}

var DefaultChartOptions = ChartOptions{
	"showVal":         true,  // value
	"showCatName":     true,  // category name
	"showLegendKey":   false, //show the cart legend
	"showSerName":     false, // series name
	"showPercent":     false,
	"showLeaderLines": false,
	"showBubbleSize":  false,
}
