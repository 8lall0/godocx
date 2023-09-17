package style

type AbstractStyle struct {
	styleName string
	Index     int
	aliases   []string
	isAuto    bool
}
