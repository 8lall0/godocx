package style

type ExtrusionType string

const (
	ExtrusionParallel    ExtrusionType = "parallel"
	ExtrusionPerspective ExtrusionType = "perspective"
)

type Extrusion struct {
	Type  ExtrusionType
	Color string
}
