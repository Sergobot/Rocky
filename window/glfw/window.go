// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package glfw

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/Sergobot/Rocky/window/base"
)

// Window is basically a wrapper for glfw's Window struct
// For now it doesn't have any killer-features but those will be
// in future releases.
// What Window does:
// - Manage GLFW window and OpenGL context in it;
// - That's all for this moment.
type Window struct {
	// Embed BasicWindow to match base.Window interface and to get some very basic
	// methods, like [Set]Geometry() and others. However, we still need to reimplement
	// some of them.
	base.BasicWindow

	window *glfw.Window
}

// create creates a full-screen window with OpenGL 3.3 context in it. It is usually
// called when Show() is run for the first time.
func (w *Window) create() {
	if !glfwInitialized {
		initGLFW()
	}

	// First of all we get primary monitor's video mode to get some
	// monitor settings to apply
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()

	// If there already was a window, it's better for us to clean-up
	// window hints first
	glfw.DefaultWindowHints()

	// Set some monitor-dependent options for
	glfw.WindowHint(glfw.RedBits, mode.RedBits)
	glfw.WindowHint(glfw.GreenBits, mode.GreenBits)
	glfw.WindowHint(glfw.BlueBits, mode.BlueBits)
	glfw.WindowHint(glfw.RefreshRate, mode.RefreshRate)

	// Set some necessary window properties
	// Rocky uses OpenGL 3.3 Core profile
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	// Rocky windows aren't resizeable by user: usually that's
	// not needed in a game. There is even no way to do that with
	// full-screen windows, so this line is just for confidence.
	glfw.WindowHint(glfw.Resizable, glfw.False)

	// Create a fullscreen window
	window, err := glfw.CreateWindow(mode.Width, mode.Height, "Rocky", nil, nil)
	if err != nil {
		log.Fatalln("Failed to create GLFW window:", err)
	}
	window.MakeContextCurrent()

	// After creating a GLFW window we set w.window to it
	w.window = window

	// Now we initialize OpenGL context in our window
	if err = initGL(); err != nil {
		log.Fatalln("Failed to initialize OpenGL context in a window:", err)
	}

	fbWidth, fbHeight := w.window.GetFramebufferSize()
	// Configure global settings
	gl.Viewport(0, 0, int32(fbWidth), int32(fbHeight))
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
}

// Show shows an already created window or creates it and shows.
func (w *Window) Show() {
	switch w.State() {
	case base.Hidden:
		w.window.Show()
	case base.NotInitialized:
		w.create()
	case base.Shown:
		// Do nothing, window is already shown
	default:
		log.Println("Unknown State detected, leaving window unchanged.")
	}
	w.BasicWindow.Show()
}

// Hide makes a window invisible but doesn't destroy it.
func (w *Window) Hide() {
	switch w.State() {
	case base.Shown:
		w.window.Hide()
	case base.Hidden:
		// Do nothing
	case base.NotInitialized:
		log.Println("Failed to hide a window: Window is not initialized.")
	default:
		log.Println("Unknown State detected, leaving window unchanged.")
	}
	w.BasicWindow.Hide()
}

// Destroy destoys a window. Destoyed windows are assumed to be just like newly created.
func (w *Window) Destroy() {
	// We don't need to destroy an already destroyed/not initialized window.
	if w.State() != base.NotInitialized {
		w.window.Destroy()
		w.BasicWindow.Destroy()
	}

}
