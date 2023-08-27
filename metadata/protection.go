package metadata

type Protection struct {
	Editing   string //none|readOnly|comments|trackedChanges|forms
	Password  string
	SpinCount int    //100000
	Algorithm string //see constants defined in \PhpOffice\PhpWord\Shared\Microsoft\PasswordEncoder
	Salt      string // Must be 16
}
