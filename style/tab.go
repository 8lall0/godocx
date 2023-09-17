package style

import "github.com/shabbyrobe/xmlwriter"

type TabType string

const (
	TAB_STOP_CLEAR   TabType = "clear"
	TAB_STOP_LEFT    TabType = "left"
	TAB_STOP_CENTER  TabType = "center"
	TAB_STOP_RIGHT   TabType = "right"
	TAB_STOP_DECIMAL TabType = "decimal"
	TAB_STOP_BAR     TabType = "bar"
	TAB_STOP_NUM     TabType = "num"
)

type TabLeaderType string

const (
	TAB_LEADER_NONE       TabLeaderType = "none"
	TAB_LEADER_DOT        TabLeaderType = "dot"
	TAB_LEADER_HYPHEN     TabLeaderType = "hyphen"
	TAB_LEADER_UNDERSCORE TabLeaderType = "underscore"
	TAB_LEADER_HEAVY      TabLeaderType = "heavy"
	TAB_LEADER_MIDDLEDOT  TabLeaderType = "middleDot"
)

type Tab struct {
	Type     TabType
	Leader   TabLeaderType
	Position float64
}

func (t *Tab) Write(w *xmlwriter.Writer) error {
	return nil
}
