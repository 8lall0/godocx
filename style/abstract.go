package style

type AbstractStyle struct {
	styleName string
	index     int
	aliases   []string
	isAuto    bool
}
