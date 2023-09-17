package style

import (
	"errors"
	"fmt"
	"github.com/shabbyrobe/xmlwriter"
)

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

func (a *ArrowStyleDUMMY) String() string {
	return string(*a)
}

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
	Frame
	Flip          bool
	ConnectorType ConnectorType
	Weight        int
	Color         string
	Dash          DashStyle
	BeginArrow    ArrowStyleDUMMY
	EndArrow      ArrowStyleDUMMY
}

func (l *Line) WriteAlignment(w *xmlwriter.Writer) error {
	var err error
	err = w.StartElem(xmlwriter.Elem{Name: "w:pPr"})

	if l.Alignment != "" {
		w.StartElem(xmlwriter.Elem{Name: "w:jc"})
		w.WriteAttr(xmlwriter.Attr{
			Name:  "w:val",
			Value: string(l.Alignment),
		})
		w.EndElem("w:jc")
	}
	w.EndElem("w:pPr")

	return err
}

var dashConvertingMap = map[DashStyle]string{
	DASH_STYLE_DASH:              "dash",
	DASH_STYLE_ROUND_DOT:         "1 1",
	DASH_STYLE_SQUARE_DOT:        "1 1",
	DASH_STYLE_DASH_DOT:          "dashDot",
	DASH_STYLE_LONG_DASH:         "longDash",
	DASH_STYLE_LONG_DASH_DOT:     "longDashDot",
	DASH_STYLE_LONG_DASH_DOT_DOT: "longDashDotDot",
}

func (l *Line) WriteStroke(w *xmlwriter.Writer) error {
	var err error
	err = w.StartElem(xmlwriter.Elem{Name: "v:stroke"})

	if l.Weight != 0 {
		err = errors.Join(err,
			w.WriteAttr(xmlwriter.Attr{Name: "weight", Value: fmt.Sprintf("%dpt", l.Weight)}),
		)
	}

	if l.Color != "" {
		err = errors.Join(err,
			w.WriteAttr(xmlwriter.Attr{Name: "color", Value: l.Color}),
		)
	}

	if l.BeginArrow != "" {
		err = errors.Join(err,
			w.WriteAttr(
				xmlwriter.Attr{Name: "startarrow", Value: string(l.BeginArrow)},
			),
		)
	}

	if l.EndArrow != "" {
		err = errors.Join(err,
			w.WriteAttr(
				xmlwriter.Attr{Name: "endarrow", Value: string(l.EndArrow)},
			),
		)
	}

	if l.Dash != "" {
		err = errors.Join(err,
			w.WriteAttr(
				xmlwriter.Attr{Name: "dashstyle", Value: dashConvertingMap[l.Dash]},
			),
		)

		if l.Dash == DASH_STYLE_ROUND_DOT {
			err = errors.Join(err,
				w.WriteAttr(
					xmlwriter.Attr{Name: "endcap", Value: "round"},
				),
			)
		}
	}

	err = errors.Join(err, w.EndElem("v:stroke"))

	return err
}
