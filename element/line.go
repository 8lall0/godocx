package element

import (
	"errors"
	"fmt"
	xw "github.com/shabbyrobe/xmlwriter"
	"godocx/style"
	"io"
)

type Line struct {
	AbsEl
	LineStyle style.Line
}

func (l *Line) Write(writer io.Writer) error {
	w := xw.Open(writer)

	err := w.StartDoc(xw.Doc{})

	if l.WithoutP == false {
		err = errors.Join(err, w.StartElem(xw.Elem{Name: "w:p"}))
		//	$styleWriter->writeAlignment();
	}
	// $this->writeCommentRangeStart();
	err = errors.Join(err,
		w.StartElem(xw.Elem{Name: "w:r"}),
		w.StartElem(xw.Elem{Name: "w:pict"}),
	)

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
		//$styleWriter->write();
		//$styleWriter->writeStroke();
		w.EndAll(),
	)

	err = errors.Join(err, w.EndAllFlush())
	return err
}
