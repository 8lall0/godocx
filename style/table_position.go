package style

type VerticalAnchor string

const (
	VANCHOR_TEXT   VerticalAnchor = "text"   // Relative to vertical text extents
	VANCHOR_MARGIN VerticalAnchor = "margin" // Relative to margin
	VANCHOR_PAGE   VerticalAnchor = "page"   // Relative to page
)

type HorizontalAnchor string

const (
	HANCHOR_TEXT   HorizontalAnchor = "text"   // Relative to text extents
	HANCHOR_MARGIN HorizontalAnchor = "margin" // Relative to margin
	HANCHOR_PAGE   HorizontalAnchor = "page"   // Relative to page
)

type HorizontalAlignment string

const (
	XALIGN_LEFT    HorizontalAlignment = "left"    // Left aligned horizontally
	XALIGN_CENTER  HorizontalAlignment = "center"  // Centered horizontally
	XALIGN_RIGHT   HorizontalAlignment = "right"   // Right aligned horizontally
	XALIGN_INSIDE  HorizontalAlignment = "inside"  // Inside
	XALIGN_OUTSIDE HorizontalAlignment = "outside" // Outside
)

type VerticalAlignment string

const (
	YALIGN_INLINE  VerticalAlignment = "inline"  // In line with text
	YALIGN_TOP     VerticalAlignment = "top"     // Top
	YALIGN_CENTER  VerticalAlignment = "center"  // Centered vertically
	YALIGN_BOTTOM  VerticalAlignment = "bottom"  // Bottom
	YALIGN_INSIDE  VerticalAlignment = "inside"  // Inside Anchor Extents
	YALIGN_OUTSIDE VerticalAlignment = "outside" // Centered vertically
)

type TablePosition struct {
	FromText struct {
		Left   int
		Right  int
		Top    int
		Bottom int
	}
	Anchor struct {
		Vertical   VerticalAnchor
		Horizontal HorizontalAnchor
	}
	Alignment struct {
		Vertical   VerticalAlignment
		Horizontal HorizontalAlignment
	}

	TblpX int
	TblpY int
}
