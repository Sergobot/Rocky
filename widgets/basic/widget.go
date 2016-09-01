// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package basic

import g "github.com/Sergobot/Rocky/geometry"

// Widget is the simpliest widget type, just a blank one.
// It implements only functions for work with size and position.
// Used to be embedded in another more specific widget struct.
type Widget struct {
	// Basic parameters: width, height and X/Y coordinates
	geometry g.Rect
}

// SetGeometry sets the rectangle (or bounding box, if you want) of a widget.
// That means, widget will have same coordinates and size as a given rect.
func (w *Widget) SetGeometry(r g.Rect) {
	w.geometry = r
}

// Geometry returns current bounding box of a widget.
func (w *Widget) Geometry() g.Rect {
	return w.geometry
}

// GetReady initializes matrices and if there was attempt to resize/move widget
// that geometry change is applied here.
func (w *Widget) GetReady() {}

// Update does nothing: Widget should not get any events or draw anything
func (w *Widget) Update() {}

// Draw does nothing: Widget's always blank, we don't draw anything on it
func (w *Widget) Draw() {}
