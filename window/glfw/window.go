// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package glfw

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/Sergobot/Rocky/opengl/gl33"
	"github.com/Sergobot/Rocky/window/basic"
	"github.com/Sergobot/Rocky/window/state"
)

// Window is basically a wrapper for glfw's Window struct
// For now it doesn't have any killer-features but those will be
// in future releases.
// What Window does:
// - Manage GLFW window and OpenGL context in it;
// - That's all for this moment.
type Window struct {
	// Embed basic.window to match basic.Window interface and to get some very basic
	// methods, like [Set]Geometry() and others. However, we still need to reimplement
	// some of them.
	basic.Window

	window *glfw.Window
}

// create creates a full-screen window with OpenGL 3.3 context in it. It is usually
// called when Show() is run for the first time.
func (w *Window) create() {
	if !glfwInitialized {
		if err := initGLFW(); err != nil {
			log.Fatalln("Failed to initialize GLFW:", err)
		}
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
	if err = gl33.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL context in a window:", err)
	}

	// Configure global settings
	fbWidth, fbHeight := w.window.GetFramebufferSize()
	gl33.SetViewport(0, 0, int32(fbWidth), int32(fbHeight))
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	// Since window is already shown during glfw.CreateWindow(), we need to set
	// appropriate State
	w.Window.Show()
}

// Show shows an already created window or creates it and shows.
func (w *Window) Show() {
	switch w.State() {
	case state.Hidden:
		w.window.Show()
	case state.NotInitialized:
		w.create()
	case state.Shown:
		// Do nothing, window is already shown
	default:
		log.Println("Unknown State detected, leaving window unchanged.")
	}
	w.Window.Show()
}

// Hide makes a window invisible but doesn't destroy it.
func (w *Window) Hide() {
	switch w.State() {
	case state.Shown:
		w.window.Hide()
	case state.Hidden:
		// Do nothing
	case state.NotInitialized:
		log.Println("Failed to hide a window: Window is not initialized.")
	default:
		log.Println("Unknown State detected, leaving window unchanged.")
	}
	w.Window.Hide()
}

// Destroy destoys a window. Destoyed windows are assumed to be just like newly created.
func (w *Window) Destroy() {
	// We don't need to destroy an already destroyed/not initialized window.
	if w.State() != state.NotInitialized {
		w.window.Destroy()
		w.Window.Destroy()
	}
}
