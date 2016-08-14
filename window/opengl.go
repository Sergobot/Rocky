// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package window

import (
	"log"

	// For now Rocky supports only OpenGL 3.3 core.
	// In future releases support for 4.x and ES may appear.
	"github.com/go-gl/gl/v3.3-core/gl"
)

func initGL() error {
	if err := gl.Init(); err != nil {
		return err
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("Successfully initialized OpenGL version", version)

	return nil
}
