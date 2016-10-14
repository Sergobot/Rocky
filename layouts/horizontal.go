// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package layouts

import g "github.com/Sergobot/Rocky/geometry"

// Horizontal makes it much easier to arrange widgets as a horisontal line
// just like that:
//  ___ ___ ___ ___
// |___|___|___|___|
//
// Width is the same for all the widgets and is calculated based on widgets count
// and layout size. Height is equal to layout's height, which is set by layout's
// holder and usually equal to its (layout holder's) height.
type Horizontal struct {
	// Embedding BasicLayout helps us a lot: we only need to reimplement Activate()
	// method and nothing more to create a custom layout.
	BasicLayout
}

// Activate does all the arrangement work for a layout: it calculates required size
// for each widget and resizes them.
func (hbl *Horizontal) Activate() {
	var widgetRect g.RectF

	count := len(hbl.Widgets())
	width := hbl.Geometry().W / float32(count)

	widgetRect.W = width
	widgetRect.H = hbl.Geometry().H
	widgetRect.X = hbl.Geometry().X
	widgetRect.Y = hbl.Geometry().Y

	for _, v := range hbl.Widgets() {
		v.SetGeometry(widgetRect)
		widgetRect.X += width
	}
}
