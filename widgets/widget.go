// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package widgets

import g "github.com/Sergobot/Rocky/geometry"

// Widget is an interface for objects, each of which a separate piece of
// window's space. There are many possible examples of widgets:
// - Buttons
// - Labels
// - Images
// - Text inputs
// - And even more!
type Widget interface {
	// GetReady does everything about initialization before drawing.
	// Usually it's called in AddWidget method of a window, so there is no need
	// to call it manually anywhere else.
	GetReady() error

	// Update updates widget's contents and reacts to events (if there are some).
	// Also, Draw() method is called as one of update steps.
	Update()

	// Draw draws something inside widget's space. By the way, it's called in
	// Update().
	Draw()

	// Some basic methods to contol widget's size and position.
	SetGeometry(g.Rect)
	Geometry() g.Rect
}
