package element

import "time"

type TrackChangeType string

const (
	INSERTED TrackChangeType = "INSERTED"
	DELETED  TrackChangeType = "DELETED"
)

type TrackChange struct {
	Container  string // "TrackChange"
	ChangeType TrackChangeType
	Author     string
	Date       time.Time
}
