package ebiten

import "errors"

var (
	errNoUpdateCallback = errors.New("no update callback specified for ebiten renderer")
	errNoRenderCallback = errors.New("no render callback specified for ebiten renderer")
)
