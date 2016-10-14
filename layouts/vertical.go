// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package layouts

import g "github.com/Sergobot/Rocky/geometry"

// Vertical is very similar to Horizontal but it arranges widgets vertically,
// not horizontally. Just like that:
//  ___
// |___|
// |___|
// |___|
// |___|
//
// Height is the same for all the widgets and is calculated based on widgets count
// and layout size. Width is equal to layout's width, which is set by layout's
// holder and usually equal to its (layout holder's) width.
type Vertical struct {
	// Embedding BasicLayout helps us a lot: we only need to reimplement Activate()
	// method and nothing more to create a custom layout.
	BasicLayout
}

// Activate does all the arrangement work for a layout: it calculates required size
// for each widget and resizes them.
func (vbl *Vertical) Activate() {
	var widgetRect g.RectF

	count := len(vbl.Widgets())
	height := vbl.Geometry().H / float32(count)

	widgetRect.W = vbl.Geometry().W
	widgetRect.H = height
	widgetRect.X = vbl.Geometry().X
	widgetRect.Y = vbl.Geometry().Y

	for _, v := range vbl.Widgets() {
		v.SetGeometry(widgetRect)
		widgetRect.Y += height
	}
}
