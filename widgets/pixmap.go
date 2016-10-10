// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package widgets

import (
	g "github.com/Sergobot/Rocky/geometry"
	opengl33 "github.com/Sergobot/Rocky/opengl/gl33"
	widgets33 "github.com/Sergobot/Rocky/widgets/gl33"
)

// Pixmap is a wrapper for all the implemented OpenGL_version_dependent Pixmaps.
// These are located in gl** directories. For now there is only one version available,
// 3.3, so we have only one Pixmap implementation inside.
type Pixmap struct {
	p33 *widgets33.Pixmap
}

// NewPixmap returns an initialized Pixmap
func NewPixmap() *Pixmap {
	p := new(Pixmap)
	p.Init()
	return p
}

// Init initializes one of Pixmap members, according to separate OpenGL versions availability.
// If there is no OpenGL context yet (or there is, but no Pixmap is able to use it),
// this method does nothing and you will have to call it manually later, when a
// supported OpenGL version is initialized.
func (p *Pixmap) Init() {
	if opengl33.Initialized() {
		p.p33 = widgets33.NewPixmap()
	}
}

// LoadFromFile loads a texture from the given image
func (p *Pixmap) LoadFromFile(file string) {
	if p.p33 != nil {
		p.p33.LoadFromFile(file)
	}
}

// SetTexture sets texture to be used by Pixmap. Its purpose is to allow
// fast switching between textures.
func (p *Pixmap) SetTexture(t *opengl33.Texture) {
	if p.p33 != nil {
		p.p33.SetTexture(t)
	}
}

// GetReady initializes the Pixmap to be ready to Draw() function calls
func (p *Pixmap) GetReady() {
	if p.p33 != nil {
		p.p33.GetReady()
	}
}

// Draw draws Pixmap's contents to the screen
func (p *Pixmap) Draw() {
	if p.p33 != nil {
		p.p33.Draw()
	}
}

// SetSize sets the widget's size. You can call it manually or through SetGeometry
func (p *Pixmap) SetSize(s g.SizeF) {
	if p.p33 != nil {
		p.p33.SetSize(s)
	}
}

// Size returns current widget's size.
func (p *Pixmap) Size() g.SizeF {
	if p.p33 != nil {
		return p.p33.Size()
	}

	return g.SizeF{}
}

// SetPos sets position of a widget. It's used in SetGeometry but you can call
// it manually.
func (p *Pixmap) SetPos(pos g.PosF) {
	if p.p33 != nil {
		p.p33.SetPos(pos)
	}
}

// Pos returns current widget's position.
func (p *Pixmap) Pos() g.PosF {
	if p.p33 != nil {
		return p.p33.Pos()
	}

	return g.PosF{}
}

// SetGeometry sets the rectangle (or bounding box, if you want) of a widget.
// That means, widget will have same coordinates and size as a given rect.
func (p *Pixmap) SetGeometry(r g.RectF) {
	if p.p33 != nil {
		p.p33.SetGeometry(r)
	}
}

// Geometry returns current bounding box of a widget
func (p *Pixmap) Geometry() g.RectF {
	if p.p33 != nil {
		return p.p33.Geometry()
	}

	return g.RectF{}
}
