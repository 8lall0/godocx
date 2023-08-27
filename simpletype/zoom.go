package simpletype

type Zoom string

const (
	ZoomNone     Zoom = "none"
	ZoomFullPage Zoom = "fullPage" //Display One Full Page
	ZoomBestFit  Zoom = "bestFit"  //Display Page Width
	ZoomTextFit  Zoom = "textFit"  //Display Text Width
)
