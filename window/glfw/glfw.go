// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package glfw

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var glfwInitialized bool

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func initGLFW() error {
	if err := glfw.Init(); err != nil {
		return err
	}

	major, minor, revision := glfw.GetVersion()
	log.Printf("Successfully initialized GLFW %d.%d.%d", major, minor, revision)
	glfwInitialized = true

	return nil
}

func terminateGLFW() {
	glfw.Terminate()
	glfwInitialized = false
	log.Println("Terminated GLFW")
}
