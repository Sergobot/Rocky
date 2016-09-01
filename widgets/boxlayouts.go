// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package widgets

import g "github.com/Sergobot/Rocky/geometry"

// HBoxLayout makes it much easier to arrange widgets as a horisontal line
// just like that:
//  ___ ___ ___ ___
// |___|___|___|___|
// Width is the same for all the widgets and is calculated based on widgets count
// and layout holder size. Height is equal to layout's height, which is equal to
// the layout holder's one.
type HBoxLayout struct {
	// Embedding BasicLayout helps us a lot: we only need to reimplement Activate()
	// method and nothing more to create a custom layout.
	BasicLayout
}

// Activate does all the arrangement work for a layout: it calculates required size
// for each widget and resizes them.
func (hbl *HBoxLayout) Activate() {
	var widgetRect g.Rect

	count := len(hbl.Widgets())
	width := hbl.Geometry().W / count

	widgetRect.W = width
	widgetRect.H = hbl.Geometry().H
	widgetRect.X = hbl.Geometry().X
	widgetRect.Y = hbl.Geometry().Y

	for _, v := range hbl.Widgets() {
		v.SetGeometry(widgetRect)
		widgetRect.X += width
	}
}

// VBoxLayout is very similar to HBoxLayout but it arranges widgets vertically,
// not horizontally. Just like that:
//  ___
// |___|
// |___|
// |___|
// |___|
// Height is the same for all the widgets and is calculated based on widgets count
// and layout holder size. Width is equal to layout's width, which is equal to the
// layout holder's one.
type VBoxLayout struct {
	// Embedding BasicLayout helps us a lot: we only need to reimplement Activate()
	// method and nothing more to create a custom layout.
	BasicLayout
}

// Activate does all the arrangement work for a layout: it calculates required size
// for each widget and resizes them.
func (vbl *VBoxLayout) Activate() {
	var widgetRect g.Rect

	count := len(vbl.Widgets())
	height := vbl.Geometry().H / count

	widgetRect.W = vbl.Geometry().W
	widgetRect.H = height
	widgetRect.X = vbl.Geometry().X
	widgetRect.Y = vbl.Geometry().Y

	for _, v := range vbl.Widgets() {
		v.SetGeometry(widgetRect)
		widgetRect.Y += height
	}
}
