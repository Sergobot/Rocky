// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package rocky

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Window is basically a wrapper for glfw's Window struct
// For now it doesn't have any killer-features but those will be
// in future releases.
// What Window does:
// - Manage GLFW window and OpenGL context in it;
// - That's all for this moment.
type Window struct {
	window *glfw.Window
}

// Create creates a full-screen window with OpenGL context in it.
func (w *Window) Create() {
	if !glfwInitialized {
		log.Fatalln("Failed to create a window: GLFW is not initialized yet")
	} else if w.window != nil {
		log.Fatalln("Failed to create a window: it is already created.")
	}

	// First of all we get primary monitor's video mode to get some
	// monitor settings to apply
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()

	// Set some monitor=dependent options for
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
	initGL()

	// Configure global settings
	gl.Viewport(0, 0, int32(mode.Width), int32(mode.Height))
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
}

// Show shows an already created window or creates it and shows.
func (w *Window) Show() {
	if !glfwInitialized {
		log.Fatalln("Failed to show a window: GLFW is not initialized yet")
	} else if w.window == nil {
		w.Create()
		w.window.Show()
	} else {
		w.window.Show()
	}
}
