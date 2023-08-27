package metadata

import "time"

type Property string

const (
	PROPERTY_TYPE_BOOLEAN Property = "b"
	PROPERTY_TYPE_INTEGER Property = "i"
	PROPERTY_TYPE_FLOAT   Property = "f"
	PROPERTY_TYPE_DATE    Property = "d"
	PROPERTY_TYPE_STRING  Property = "s"
	PROPERTY_TYPE_UNKNOWN Property = "u"
)

type DocInfo struct {
	Creator        string
	LastModifiedBy string
	Created        time.Time
	Modified       time.Time
	Title          string
	Description    string
	Subject        string
	Keywords       string
	Category       string
	Company        string
	Manager        string
	//CustomProperties WHAT DE LA FUK
}
