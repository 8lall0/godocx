package style

/**
 * Points.
 *
 * - Arc: startAngle endAngle; 0 = top center, moving clockwise
 * - Curve: from-x1,from-y1 to-x2,to-y2 control1-x,control1-y control2-x,control2-y
 * - Line: from-x1,from-y1 to-x2,to-y2
 * - Polyline: x1,y1 x2,y2 ...
 * - Rect and oval: Not applicable
 *
 * @var string
 */

type Shape struct {
	Points    string
	Roundness float64 //0 = straightest (rectangular); 1 = roundest (circle/oval).
	Frame
	Fill
	//Outline Outline
	Shadow
	Extrusion
}
