package element

import "godocx/style"

type ChartType string

const (
	chartPie                  ChartType = "pie"
	chartDoughnut             ChartType = "doughnut"
	chartBar                  ChartType = "bar"
	chartColumn               ChartType = "column"
	chartLine                 ChartType = "line"
	chartArea                 ChartType = "area"
	chartScatter              ChartType = "scatter"
	chartRadar                ChartType = "radar"
	chartStackedBar           ChartType = "stacked_bar"
	chartPercentStackedBar    ChartType = "percent_stacked_bar"
	chartStackedColumn        ChartType = "stacked_column"
	chartPercentStackedColumn ChartType = "percent_stacked_column"
)

type ChartSeries struct {
	Categories string // Must be array
	Values     string // Must be array
	Name       string
}

type Chart struct {
	collectionRelation bool // true
	Type               ChartType
	Series             []ChartSeries
	Style              style.Chart
}
