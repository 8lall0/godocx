package style

type ConnectorType string

const CONNECTOR_TYPE_STRAIGHT ConnectorType = "straight"

type ArrowStyleDUMMY string

const (
	ARROW_STYLE_BLOCK   ArrowStyleDUMMY = "block"
	ARROW_STYLE_OPEN    ArrowStyleDUMMY = "open"
	ARROW_STYLE_CLASSIC ArrowStyleDUMMY = "classic"
	ARROW_STYLE_DIAMOND ArrowStyleDUMMY = "diamond"
	ARROW_STYLE_OVAL    ArrowStyleDUMMY = "oval"
)

type DashStyle string

const (
	DASH_STYLE_DASH              DashStyle = "dash"
	DASH_STYLE_ROUND_DOT         DashStyle = "rounddot"
	DASH_STYLE_SQUARE_DOT        DashStyle = "squaredot"
	DASH_STYLE_DASH_DOT          DashStyle = "dashdot"
	DASH_STYLE_LONG_DASH         DashStyle = "longdash"
	DASH_STYLE_LONG_DASH_DOT     DashStyle = "longdashdot"
	DASH_STYLE_LONG_DASH_DOT_DOT DashStyle = "longdashdotdot"
)

type Line struct {
	Flip          bool
	ConnectorType ConnectorType
	Weight        int
	Color         string
	Dash          DashStyle
	BeginArrow    ArrowStyle
	EndArrow      ArrowStyle
}
