package element

import "godocx/style"

type ShapeType string

const (
	ShapeArc      ShapeType = "arc"
	ShapeCurve    ShapeType = "curve"
	ShapeLine     ShapeType = "line"
	ShapePolyLine ShapeType = "polyline"
	ShapeRect     ShapeType = "rect"
	ShapeOval     ShapeType = "oval"
)

type Shape struct {
	Type  ShapeType
	Style style.Shape
}
