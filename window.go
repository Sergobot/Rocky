// Copyright (c) Sergey Popov <sergobot@protonmail.com>
// See LICENSE for details.

package rocky

import (
    "log"
    "runtime"

    "github.com/go-gl/gl/v3.3-core/gl"
    "github.com/go-gl/glfw/v3.1/glfw"
)

// Use custom type to store window's state
type winState int
const (
    winClosed winState = iota
    winHidden winState = iota
    winShown winState = iota
)

// Keep GLFW and GL from re-initializing
var glfwInitialized bool
var glInitialized bool

// Window represents a wrapper of GLWF window to make it even easier to use
// Main functions:
// - Manage the window and OpenGL context in it
// - Nothing more for now.
type Window struct {
    // Basic window's parametres
    width, height int
    xPos, yPos int
    title string

    // GLFW window, controlled by this struct
    window *glfw.Window
    
    // Current state of the window. By default it is 'closed'
    state winState
}

func init() {
    // GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
    
    glfwInitialized = false
    glInitialized = false
}

// Init initializes or clears the window.
func (w *Window) Init() *Window {
    // width and height are 0 by default, so we set our own defaults
    *w = Window {
        width: 800,
        height: 600,

        title : "Rocky",
    }

	return w
}

// NewWindow returns an initialized window.
func NewWindow() *Window { return new(Window).Init() }

// initGLFW initializes GLFW to create a GLFW window later.
func (w *Window) initGLFW() {
    // If GLFW is already initialized, don't re-init it
    if glfwInitialized {
        log.Println("Prevent GLFW from re-initializing")
        return
    }
    
    // Initialize GLFW
    if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize GLFW:", err)
	}
    glfwInitialized = true

    // Set some necessary window properties
    // Rocky uses OpenGL 3.3 Core profile
    glfw.WindowHint(glfw.ContextVersionMajor, 3) 
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
    
    // Rocky windows aren't resizeable by user: usually that's not needed in a game
	glfw.WindowHint(glfw.Resizable, glfw.False)
}

// initGL initializes OpenGL context.
func (w *Window) initGL() {
    // If OpenGL is already initialized, don't re-init it
    if glInitialized {
        log.Println("Prevent OpenGL from re-initializing")
        return
    }
    
    // Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}
    glInitialized = true

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("Using OpenGL", version)
	
	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
}

// Show method shows the window.
func (w *Window) Show() {
	if w.state == winHidden {
        w.window.Show()
        return
    } else if w.state == winShown {
        log.Println("Prevented re-showing a window")        
        return
    }
    
    // First we initialize GLFW
    // Then we create a window
    // And only after that we initialize OpenGL context
    w.initGLFW()
    defer w.initGL()

	window, err := glfw.CreateWindow(w.width, w.height, w.title, nil, nil)
	if err != nil {
	    log.Fatalln("Failed to create GLFW window:", err)
	}
	window.MakeContextCurrent()
    
    // After creating a GLFW window we set w.window to it
    w.window = window
}

// Close closes the window and terminates GLFW.
func (w *Window) Close() {
    // glfw.Terminate does all the work: it also closes all the remaining windows
    glfw.Terminate()
    w.state = winClosed
    glfwInitialized, glInitialized = false, false
}

// Hide hides the window: it becomes invisible but still exists.
func (w *Window) Hide() {
    // Hide the window only if it's shown
    if glfwInitialized && w.state != winClosed {
        if w.state != winHidden {
            w.window.Hide()
            w.state = winHidden
        } else {
            log.Println("Prevented re-hiding a window")
        }
    } else {
        log.Println("Prevented hiding window, when GLFW isn't initialized")
    }
}

// SetPosition moves the window on the screen.
func (w *Window) SetPosition(x, y int) {
    // Move the window only if it's not closed
    if glfwInitialized && w.state != winClosed {
        w.window.SetPos(x, y)
    } else {
        log.Println("Prevented moving a window, when GLFW isn't initialized")
    }
    // Even if the window isn't moved, we save new parameters to set them when showing the window again
    w.xPos, w.yPos = x, y
}

// Position returns current coordinates of the window on the screen
func (w *Window) Position() (int, int) {
    return w.xPos, w.yPos
}

// SetSize method resizes the window.
func (w *Window) SetSize(width, height int) {
    if glfwInitialized && w.state != winClosed {
        w.window.SetSize(width, height)
    } else {
        log.Println("Prevented moving a window, when GLFW isn't initialized")
    }
    // Even if the window isn't resized, we save new parameters to set them when showing the window again
    w.width, w.height = width, height
}

// Size returns current width and height of the window
func (w *Window) Size() (int, int) {
    return w.width, w.height
}

// Update method updates all the window contents: redraws widgets, models, etc.
func (w *Window) Update() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    // TODO:
    // - Add support for widgets
    // - Add support for FPS counting
    
    w.window.SwapBuffers()
    glfw.PollEvents()
}

// ShouldClose returns true if window is going to be closed.
func (w *Window) ShouldClose() bool {
    if !glfwInitialized {
        // If the window doesn't exist, we need to stop drawing
        return false
    }        
    return w.window.ShouldClose()
}