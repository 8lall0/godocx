package simpletype

type docProtectType string

const (
	DocProtectNone           docProtectType = "none"
	DocProtectReadOnly       docProtectType = "readOnly"       //Allow No Editing.
	DocProtectComments       docProtectType = "comments"       //Allow Editing of Comments.
	DocProtectTrackedChanges docProtectType = "trackedChanges" //Allow Editing With Revision Tracking.
	DocProtectForms          docProtectType = "forms"          //Allow Editing of Form Fields.
)
