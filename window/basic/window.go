// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package basic

import (
	g "github.com/Sergobot/Rocky/geometry"
	"github.com/Sergobot/Rocky/window/state"
)

// Window is used to be embedded in other, more specific window structs.
// It deals only with geometry, state and *nothing* more. But anyway there are all the
// methods required to be Window, so this is (almost) an abstract window.
// Please don't perform any operations on Window's members directly, use
// Window's methods instead.
type Window struct {
	geometry g.Rect
	state    state.State
}

// SetGeometry sets geometry (bounding box) of a window.
// Window impements nothing more than just setting and getting that geometry.
func (w *Window) SetGeometry(r g.Rect) {
	w.geometry = r
}

// Geometry returns geometry (bounding box) of a window.
func (w *Window) Geometry() g.Rect {
	return w.geometry
}

// Show makes a window visible. In some cases it may create a window itself.
// Please, call this method when reimplementing it in a more specific window struct.
func (w *Window) Show() {
	w.state = state.Shown
}

// Hide makes a window invisible but doesn't destroy it. Please, call
// this method when reimplementing it in a more specific window struct.
func (w *Window) Hide() {
	w.state = state.Hidden
}

// Destroy destoys a window. Destoyed windows are assumed to be just like newly created.
// Don't forget to call this method, if you're gonna reimplement it.
func (w *Window) Destroy() {
	w.state = state.NotInitialized

	// It's safe enough, right?
	w.geometry = *new(g.Rect)
}

// State returns current window state. It can be Shown, Hidden or even
// NotInitialized.
func (w *Window) State() state.State {
	return w.state
}

// TODO: Add support for layouts
func (w *Window) SetLayout() {

}

func (w *Window) Layout() {

}
