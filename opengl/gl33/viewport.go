// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package gl33

import (
	"github.com/go-gl/gl/v3.3-core/gl"

	g "github.com/Sergobot/Rocky/geometry"
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

// NormalizedViewportSize returns SizeF struct containing current viewport size.
// This size isn't in pixels or anything like that. Bigger side is always 2.0,
// smaller is similar to viewport's one. That's made to avoid pain of dealing with
// pixels manually.
// For example, if viewport is exactly 1920 pixels wide and 1080 pixels tall,
// this function will return 2.0 for width and 2.0 * 1080 / 1920 = 1.125 for height.
// If the function is called on a phone with a FullHD screen in portrait rotation,
// it will return 2.0 * 1080 / 1920 = 1.125 for width and 2.0 for height.
func NormalizedViewportSize() g.SizeF {
	var w, h float32

	ratio := ViewportAspectRatio()
	if ratio >= 1.0 {
		// Width is bigger or equals height
		w = 2.0
		h = w / ratio
	} else {
		// Height is bigger than width.
		h = 2.0
		w = h * ratio
	}

	return g.SizeF{W: w, H: h}
}

// ViewportAspectRatio returns aspect ratio of viewport, calculated as Width / Height. Note,
// if SetViewport wasn't called before this function, AspectRatio will return 0.
func ViewportAspectRatio() float32 {
	return float32(vpW) / float32(vpH)
}
