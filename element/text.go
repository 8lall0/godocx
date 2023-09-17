package element

import (
	"errors"
	"github.com/shabbyrobe/xmlwriter"
	"godocx/style"
)

type Text struct {
	AbsEl
	Text           string
	FontStyle      style.Font
	ParagraphStyle style.Paragraph
}

func (t *Text) Write(w *xmlwriter.Writer) error {
	var err error

	if t.WithoutP == false {
		err = errors.Join(
			w.StartElem(xmlwriter.Elem{Name: "w:p"}),
			t.ParagraphStyle.Write(w),
		)
	}

	err = errors.Join(err,
		t.writeCommentRangeStart(w),
		t.writeOpeningTrackChange(w),
		w.StartElem(xmlwriter.Elem{Name: "w:r"}),
		t.FontStyle.Write(w),
	)

	if t.TrackChange.ChangeType == DELETED {
		err = errors.Join(err, w.StartElem(xmlwriter.Elem{Name: "w:delText"}))
	} else {
		err = errors.Join(err, w.StartElem(xmlwriter.Elem{Name: "w:t"}))
	}

	err = errors.Join(err,
		w.WriteAttr(xmlwriter.Attr{Name: "xml:space", Value: "preserve"}),
		w.WriteText(t.Text),
		w.EndAny(),
		w.EndElem("w:r"),
		t.writeClosingTrackChange(w),
	)

	if t.WithoutP == false {
		err = errors.Join(err, w.EndElem("w:p"))
	}

	return err
}

func (t *Text) writeOpeningTrackChange(w *xmlwriter.Writer) error {
	var err error
	if t.TrackChange == nil {
		return nil
	}
	if t.TrackChange.ChangeType == INSERTED {
		err = errors.Join(err, w.StartElem(xmlwriter.Elem{Name: "w:ins"}))
	} else if t.TrackChange.ChangeType == DELETED {
		err = errors.Join(err, w.StartElem(xmlwriter.Elem{Name: "w:del"}))
	}
	err = errors.Join(err,
		w.WriteAttr(
			xmlwriter.Attr{Name: "w:author", Value: t.TrackChange.Author},
			xmlwriter.Attr{Name: "w:id", Value: t.ElementId},
		),
	)
	if t.TrackChange.Date.IsZero() == false {
		err = errors.Join(err, w.WriteAttr(xmlwriter.Attr{Name: "w:date", Value: t.TrackChange.Date.Format("2006-01-02T15:04:05Z")}))
	}

	return err
}

func (t *Text) writeClosingTrackChange(w *xmlwriter.Writer) error {
	return w.EndElem("w:ins", "w:del")
}
