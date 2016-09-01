// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package base

import (
	g "github.com/Sergobot/Rocky/geometry"
)

// State variables contain possible states of a window.
type State int

// Most usual window states.
const (
	NotInitialized State = iota
	Shown          State = iota
	Hidden         State = iota
)

// Window represents holder of an OpenGL context. It can be resized and moved,
// shown, hidden and destoyed. Also, windows are assumed to be able
// to hold a layout of widgets.
// However, on mobile devices resizing and moving a widget is kinda impossible,
// so, in these implementations of Window, geometry method should be empty.
type Window interface {
	// Bery basic and obviouds methods to control window's Geometry
	Geometry() g.Rect
	SetGeometry(g.Rect)

	// Methods supposed to work on the state of a window.
	Show()
	Hide()
	Destroy()
	State() State

	// Sets a layout of widgets to the window. These widgets in the layout will be drawn
	// in Window.Update()
	SetLayout()
	Layout()
}

// BasicWindow is used to be embedded in other, more specific window structs.
// It deals only with geometry, state and *nothing* more. But anyway there are all the
// methods required to be Window, so this is (almost) an abstract window.
// Please don't perform any operations on BasicWindow's members directly, use
// BasicWindow's methods instead.
type BasicWindow struct {
	geometry g.Rect
	state    State
}

// SetGeometry sets geometry (bounding box) of a window.
// BasicWindow impements nothing more than just setting and getting that geometry.
func (bw *BasicWindow) SetGeometry(r g.Rect) {
	bw.geometry = r
}

// Geometry returns geometry (bounding box) of a window.
func (bw *BasicWindow) Geometry() g.Rect {
	return bw.geometry
}

// Show makes a window visible. In some cases it may create a window itself.
// Please, call this method when reimplementing it in a more specific window struct.
func (bw *BasicWindow) Show() {
	bw.state = Shown
}

// Hide makes a window invisible but doesn't destroy it. Please, call
// this method when reimplementing it in a more specific window struct.
func (bw *BasicWindow) Hide() {
	bw.state = Hidden
}

// Destroy destoys a window. Destoyed windows are assumed to be just like newly created.
// Don't forget to call this method, if you're gonna reimplement it.
func (bw *BasicWindow) Destroy() {
	bw.state = NotInitialized

	// It's safe enough, right?
	bw.geometry = *new(g.Rect)
}

// State returns current window state. It can be Shown, Hidden or even
// NotInitialized.
func (bw *BasicWindow) State() State {
	return bw.state
}

// TODO: Add support for layouts
func (bw *BasicWindow) SetLayout() {

}

func (bw *BasicWindow) Layout() {

}
