// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package widgets

import (
	g "github.com/Sergobot/Rocky/geometry"
	"github.com/Sergobot/Rocky/opengl"
	ogl33 "github.com/Sergobot/Rocky/opengl/gl33"
	wgts33 "github.com/Sergobot/Rocky/widgets/gl33"
)

// Pixmap is one of the simplest widgets, intended to draw raster images
// using OpenGL. It uses gl**.Texture struct for image loading and gl**.ShaderProgram
// for drawing.
type Pixmap interface {
	// Look in widget.go to learn more about these basic methods
	GetReady()
	Draw()

	SetSize(g.SizeF)
	Size() g.SizeF

	SetPos(g.PosF)
	Pos() g.PosF

	SetGeometry(g.RectF)
	Geometry() g.RectF

	// Pixmap-specific methods are going below

	// LoadFromFile loads a texture from the given image
	LoadFromFile(string)

	// SetTexture sets texture to be used by Pixmap. Its purpose is to allow
	// fast switching between textures.
	// Note, we are passing here an interface, so there will be a pointer passed,
	// not an actual value, so every Pixmap with the same texture pointer inside
	// will have the same image.
	SetTexture(opengl.Texture)
}

// NewPixmap returns a struct, which implements Pixmap interface defined above.
func NewPixmap() Pixmap {
	if ogl33.Initialized() {
		return wgts33.NewPixmap()
	}
	return nil
}
