// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package gl33

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Variable to store initialization state
var initialized bool

// Init initializes OpenGL 3.3
func Init() error {
	if err := gl.Init(); err != nil {
		initialized = false
		return err
	}

	initialized = true

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("Successfully initialized OpenGL version", version)

	return nil
}

// Initialized returns true, if OpenGL 3.3 was successfully initialized earlier.
func Initialized() bool {
	return initialized
}
