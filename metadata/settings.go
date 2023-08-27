package metadata

import (
	"godocx/simpletype"
	"godocx/style"
)

type Settings struct {
	Zoom                  simpletype.Zoom
	MirrorMargins         bool
	HideSpellingErrors    bool
	HideGrammaticalErrors bool
	TrackRevisions        bool
	DoNotTrackMoves       bool
	DoNotTrackFormatting  bool
	DocumentProtection    Protection
	EvendAndOddHeaders    bool
	ThemeFontLang         style.Language
	UpdateFields          bool
	DecimalSymbol         string // "."
	AutoHypenation        bool
	ConsecutiveHypenLimit int
	HypennationZone       float64
	DoNotHypenateCaps     bool
	//RevisionView TrackChangesView
	//ProofState \PhpOffice\PhpWord\ComplexType\ProofState
}
