// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

// Since GLFW does not support Android, we need to use different implementations.
// New() in this file returns window from glfw subpackage and no more, because GLFW
// is enough for Windows, Linux and macOS.
// +build !android

package window

import (
	g "github.com/Sergobot/Rocky/geometry"
	"github.com/Sergobot/Rocky/window/glfw"
	"github.com/Sergobot/Rocky/window/state"
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
	State() state.State

	// Sets a layout of widgets to the window. These widgets in the layout will be drawn
	// in Window.Update()
	SetLayout()
	Layout()
}

// New returns a newly created window. For now it returns only GLFW window, but later
// (when I will implement) it will be able to return other windows. For example,
// There may be a wrapper designed to work on Android.
func New() Window {
	return new(glfw.Window)
}
