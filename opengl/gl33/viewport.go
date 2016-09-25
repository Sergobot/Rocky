// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package gl33

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

// These variables contain viewport's size in pixels. If there is any other way to get
// viewport's size - contribute please.
var (
	vpX int32
	vpY int32
	vpW int32
	vpH int32
)

// SetViewport calls gl.Viewport with given size and stores it.
// If there is a function that gets notified about framebuffer size change, call
// this method instead of manual direct calling gl.Viewport.
// We do this because [there is no way]/[I don't know how] to get viewport size
// directly from OpenGL. If you know a solution, contribute please.
func SetViewport(x, y, width, height int32) {
	vpX, vpY, vpW, vpH = x, y, width, height
	gl.Viewport(vpX, vpY, vpW, vpH)
}

// Viewport returns current size of the viewport. As it was mentioned above,
// contribute please, it you know better solution to get OpenGL viewport.
func Viewport() (int32, int32, int32, int32) {
	return vpX, vpY, vpW, vpH
}

// ViewportAspectRatio returns aspect ratio of viewport, calculated as Width / Height. Note,
// if SetViewport wasn't called before this function, AspectRatio will return 0.
func ViewportAspectRatio() float32 {
	return float32(vpW) / float32(vpH)
}
