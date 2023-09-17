package element

import (
	"errors"
	"fmt"
	xw "github.com/shabbyrobe/xmlwriter"
	"godocx/style"
)

type Line struct {
	AbsEl
	Style style.Line
}

func (l *Line) Write(w *xw.Writer) error {
	var err error
	if l.WithoutP == false {
		err = errors.Join(err,
			w.StartElem(xw.Elem{Name: "w:p"}),
			l.Style.WriteAlignment(w),
		)
	}
	err = errors.Join(err,
		l.writeCommentRangeStart(w),
		w.StartElem(xw.Elem{Name: "w:r"}),
		w.StartElem(xw.Elem{Name: "w:pict"}),
	)

	// TODO rimuovi
	l.ElementIndex = 1
	if l.ElementIndex == 1 {
		err = errors.Join(err,
			w.StartElem(xw.Elem{Name: "v:shapetype"}),
			w.WriteAttr(
				xw.Attr{Name: "id", Value: "_x0000_t32"},
				xw.Attr{Name: "coordsize", Value: "21600,21600"},
				xw.Attr{Name: "o:spt", Value: "32"},
				xw.Attr{Name: "o:oned", Value: "t"},
				xw.Attr{Name: "path", Value: "m,l21600,21600e"},
				xw.Attr{Name: "filled", Value: "f"},
			),

			w.StartElem(xw.Elem{Name: "v:path"}),
			w.WriteAttr(
				xw.Attr{Name: "arrowok", Value: "t"},
				xw.Attr{Name: "fillok", Value: "f"},
				xw.Attr{Name: "o:connecttype", Value: "none"},
			),
			w.EndAny(),
			w.StartElem(xw.Elem{Name: "o:lock"}),
			w.WriteAttr(
				xw.Attr{Name: "v:ext", Value: "edit"},
				xw.Attr{Name: "shapetype", Value: "t"},
			),
			w.EndAny(),
			w.EndAny(),
		)
	}

	err = errors.Join(err,
		w.StartElem(xw.Elem{Name: "v:shape"}),
		w.WriteAttr(
			xw.Attr{Name: "id", Value: fmt.Sprintf("_x0000_s1%03d", l.ElementIndex)},
			xw.Attr{Name: "type", Value: "#_x0000_t32"},
		),
		l.Style.Write(w),
		l.Style.WriteStroke(w),
		// Potenziale BUG!
		w.EndAll(),
	)

	return err
}
