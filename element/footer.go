package element

type FooterType string

const (
	AUTO  = "default" // default and odd pages
	FIRST = "first"
	EVEN  = "even"
)

type Footer struct {
	container string     // Footer
	Type      FooterType //Auto
}
