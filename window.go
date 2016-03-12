// Copyright (c) Sergey Popov <sergobot@protonmail.com>
// See LICENSE for details.

package rocky

import (
    "log"
    "runtime"

    "github.com/go-gl/gl/v3.3-core/gl"
    "github.com/go-gl/glfw/v3.1/glfw"
)

// Window represents a wrapper of GLWF window to make it even easier to use
// Main functions:
// - Manage the window and OpenGL context in it
// - Nothing more for now
type Window struct {
    // Basic window's parametres
    width, height int
    xPos, yPos int
    title string

    // GLFW window, controlled by this struct
    window glfw.Window
    // Is true only if GLFW is already initialized
    initialized bool
}

func init() {
    // GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

// Init initializes or clears Window w
func (w *Window) Init() *Window {
    // width, height, xPos and yPos are 0 by default, so we set our own defaults
    *w = Window {
        width: 800,
        height: 600,

        xPos: 0,
        yPos: 0,

        title : "Rocky",
    }
    
    if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize GLFW:", err)
	}
	w.initialized = true

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(w.width, w.height, w.title, nil, nil)
	if err != nil {
	    log.Fatalln("Failed to create GLFW window:", err)
	}
	window.MakeContextCurrent()
	
	// Initialize Glow
	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}
	
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("Using OpenGL", version)
	
	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	
	return w
}

// New returns an initialized window
func New() *Window { return new(Window).Init() }

// Update method updates all the window contents: redraws widgets, models, etc.
func (w *Window) Update() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    // TODO:
    // Add support for widgets
    // Add support for FPS counting
    
    w.window.SwapBuffers()
    glfw.PollEvents()
}

// Resize method resizes the window
func (w *Window) Resize(width, height int) {
    w.width, w.height = width, height
    w.window.SetSize(width, height)
}

// ShouldClose returns true if window is going to be closed
func (w *Window) ShouldClose() bool {
    return w.window.ShouldClose()
}