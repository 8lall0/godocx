package element

import "godocx/style"

type Image struct {
	source          string
	sourceType      string
	style           style.ImageFrame //imgstyle
	watermark       bool
	name            string
	imageType       string
	imageCreateFunc string
	imageFunc       int // WHAT DE LA FUK
	imageExtension  string
	imageQuality    int
	memoryImage     bool
	target          string
	mediaIndex      int
	mediaRelation   bool
}
