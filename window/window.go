// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package window

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	g "github.com/Sergobot/Rocky/geometry"
)

func init() {
	initGLFW()
	internal = new(Window)
	internal.Create()
}

// Set sets a window to be used by other Rocky structs. There can be only one currently
// used window.
func Set(w *Window) {
	// TODO:
	// Notify widgets and others about window change
	internal = w
}

// Get returns a window used by Rocky now. It's often called in widgets' GetReady()
// or SetGeometry()
func Get() *Window {
	return internal
}

// Window is basically a wrapper for glfw's Window struct
// For now it doesn't have any killer-features but those will be
// in future releases.
// What Window does:
// - Manage GLFW window and OpenGL context in it;
// - That's all for this moment.
type Window struct {
	window *glfw.Window

	// Rectangle, representing window's bounding box
	geometry g.Rect
}

var internal *Window

// Create creates a full-screen window with OpenGL context in it.
func (w *Window) Create() {
	if !glfwInitialized {
		initGLFW()
	} else if w.window != nil {
		log.Fatalln("Failed to create a window: it is already created.")
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

	// Configure global settings
	gl.Viewport(0, 0, int32(mode.Width), int32(mode.Height))
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
}

// Close destroys a window, so it gets necessary to re-create it first
func (w *Window) Close() {
	if !glfwInitialized {
		log.Fatalln("Failed to close a window: GLFW is not initialized yet.")
	} else if w.window == nil {
		log.Fatalln("Failed to close a window: window is not shown yet.")
	} else {
		// GLFW is initialized and window is shown, so we need to destroy it.
		w.window.Destroy()
		w.window = nil
	}
}

// Show shows an already created window or creates it and shows.
func (w *Window) Show() {
	// Since GLFW is initialized before anything else, checking
	// glfwInitialized here may look overkill. But it's not, because GLFW can
	// be terminated and in that case w.Create() re-inits it.
	if glfwInitialized {
		// "Zero value" for pointers is nil, so newly allocated window will be
		// created here
		if w.window == nil {
			// There is no way for now to create fullscreen window initially hidden.
			// Window will be shown during glfw.CreateWindow call in w.Create.
			w.Create()
		} else {
			// This block is required for cases when a window was closed or hidden.
			// Also, it's safe if window was already created and shown.
			w.window.Show()
		}
	} else {
		// GLFW isn't initialized yet and window isn't created either.
		if w.window == nil {
			// w.Create() does all the work for us:
			// it initializes GLFW and creates a window, which is initially shown.
			w.Create()
		} else {
			// This should not happen: not initialized GLFW and already created window
			// is kinda a bug. But here we try to recover.
			// NOTE: There is at least one case this may happen: GLFW was terminated and w.window
			// destroyed but wasn't set to nil. So, when terminating GLFW we need to destroy all
			// the active windows first using Close().

			log.Println("Rocky has run into an internal error: GLFW isn't initialized but a window is created")
			log.Println("Trying to recover")
			// First of all we try to init GLFW: it's safe even if GLFW is already initialized
			if err := initGLFW(); err != nil {
				log.Fatalln("Failed to initialize GLFW:", err)
			}

			// Next we set w.window to nil and re-create it
			w.window = nil
			w.Create()

			// Lines above should succeed. But if those didn't message below won't be printed
			log.Println("Rocky has fixed the problem. If not, please report this bug.")
		}
	}
}

// SetGeometry sets geometry (bounding box) of a window.
func (w *Window) SetGeometry(r g.Rect) {
	// TODO:
	// Make window able to resize
	w.geometry = r
}

// Geometry returns geometry (bounding box) of a window.
func (w *Window) Geometry() g.Rect {
	return w.geometry
}
