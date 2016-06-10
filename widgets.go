// Copyright (c) Sergey Popov <sergobot@protonmail.com>
// See LICENSE for details.

package rocky

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Widget is the interface every widget should implement.
// Widget is a piece of window's space, which has special view.
// Widgets are:
// - Buttons
// - Labels
// - Images
// - Text fields
// - And even more!
type Widget interface {
	// GetReady is used to prepare something before drawing.
	// Called only once, when the widget is added to a window
	GetReady()

	// Draws something in the widget's space
	Draw()

	SetSize(w, h int)
	Size() (int, int)

	SetPos(x, y int)
	Pos() (int, int)

	SetWindowSize(x, y int)
}

// BasicWidget is the simpliest widget type, just a blank widget.
// It implements only functions for work with size, position and event handling.
// Mention that we have to reimplement event handling function in inherited widgets,
// since BasicWidget has only blank .ProcessEvent() method.
// Used to be inherited by other widgets.
type BasicWidget struct {
	// Basic parameters
	xPos, yPos    int
	width, height int
	// Size of a window a widget is attached to
	wWidth, wHeight int

	// This matrix is responsible for widget moving
	transMat mgl32.Mat4
	// And this one - for widget resizing/scaling
	scaleMat mgl32.Mat4

	// Is true only if the widget is ready to be drawn
	ready bool

	// And nothing more!
}

// init is used to initialize widget's matrices and currently nothing more.
// If you have an idea about better solution, contribute please.
func (w *BasicWidget) init() *BasicWidget {
	w.transMat = mgl32.Ident4()
	w.scaleMat = mgl32.Ident4()
	return w
}

// NewBasicWidget is used to create initialized BasicWidget. For now, that's required
// to have identity matrices instead of null ones. As mentioned above, if you
// know better solution, contribute please.
func NewBasicWidget() *BasicWidget { return new(BasicWidget).init() }

// SetSize sets the widget's size
func (w *BasicWidget) SetSize(width, height int) {
	xScaleRatio := float32(width) / float32(w.width)
	yScaleRatio := float32(height) / float32(w.height)
	zScaleRatio := float32(1)

	// This is because of size effect of scaling: our widget moves and we need to move it back
	// First of all we convert current, "good" widget coordinates to use in OpenGL
	xCur := float32(2*w.xPos-w.wWidth) / float32(w.wWidth)
	yCur := float32(w.wHeight-2*w.yPos) / float32(w.wHeight)
	// Then we find new ("bad") ones
	xNew := xCur * xScaleRatio
	yNew := yCur * yScaleRatio

	// And, finally, apply transformations to matrices
	w.transMat = w.transMat.Mul4(mgl32.Translate3D(xCur-xNew, yCur-yNew, float32(0)))
	w.scaleMat = w.scaleMat.Mul4(mgl32.Scale3D(xScaleRatio, yScaleRatio, zScaleRatio))

	// Don't forget to update width and height in the widget itself
	w.width, w.height = width, height
}

// Size method returns the widget's current width and height
func (w *BasicWidget) Size() (int, int) {
	return w.width, w.height
}

// SetPos moves the widget on the window, widget is attached to
func (w *BasicWidget) SetPos(xPos, yPos int) {
	// First we convert pixels to OpenGL coordinates
	xTrans := 2 * float32(xPos-w.xPos) / float32(w.wWidth)
	yTrans := 2 * float32(w.yPos-yPos) / float32(w.wHeight)
	zTrans := float32(0) // We don't tranlsate z coordinate

	// Then multiply widget's translation matrix and one returned by 'Translate2D'
	w.transMat = w.transMat.Mul4(mgl32.Translate3D(xTrans, yTrans, zTrans))

	// And, finally, we assign widget's position to the new one
	w.xPos, w.yPos = xPos, yPos
}

// Pos method returns the widget's top-left corner current
// coordinates on the window, the widget is attached to
func (w *BasicWidget) Pos() (int, int) {
	return w.xPos, w.yPos
}

// SetWindowSize sets size of a window (that's needed in real GetReady methods.
// If you have an idea about better solution, contribute please)
func (w *BasicWidget) SetWindowSize(wWidth, wHeight int) {
	w.wWidth, w.wHeight = wWidth, wHeight
	w.width, w.height = wWidth, wHeight
}

// GetReady does nothing: BasicWidget has nothing to prepare for drawing
func (w *BasicWidget) GetReady() { w.ready = true }

// Draw does nothing: BasicWidget's always blank, we don't draw anything on it
func (w *BasicWidget) Draw() {}

// ProcessEvent does nothing: BasicWidget doesn't have to react on anything.
// In real ProcessEvent method() you may want to pass an event to another method
// using if/else or switch/case. That's a good practice.
func (w *BasicWidget) ProcessEvent(e Event) {}

// Pixmap is one of the simplest widgets, intended to draw raster images
// using OpenGL. It uses Texture struct for image loading and ShaderProgram
// for drawing. Also Pixmap inherits very basic methods from BasicWidget.
type Pixmap struct {
	BasicWidget

	texture       *Texture
	vao, vbo, ebo uint32
}

// init is used to initialize pixmap's matrices and currently nothing more.
// If you have an idea about better solution, contribute please.
func (p *Pixmap) init() *Pixmap {
	p.transMat = mgl32.Ident4()
	p.scaleMat = mgl32.Ident4()
	return p
}

// NewPixmap is used to create initialized Pixmap. For now, that's required
// to have identity matrices instead of null ones.
// As mentioned above, if you know better solution, contribute please.
func NewPixmap() *Pixmap { return new(Pixmap).init() }

// LoadFromFile loads a texture from the given image.
func (p *Pixmap) LoadFromFile(file string) {
	err := p.texture.LoadFromFile(file)
	if err != nil {
		log.Println("Failed to load image to a Pixmap:", err)
	}
}

// SetTexture sets texture to be used by Pixmap. Its purpose is to allow
// fast switching between textures.
func (p *Pixmap) SetTexture(tex *Texture) {
	if tex.Ready() {
		p.texture = tex
	} else {
		log.Println("Prevented setting a not ready texture to pixmap")
	}
}

// GetReady initializes the Pixmap to be ready to Draw() function calls.
func (p *Pixmap) GetReady() {
	if p.ready {
		return
	}

	p.ready = false

	if !PixmapShaderProgram.Linked() {
		var vShader, fShader Shader
		if err := vShader.Compile(PixmapVertexShaderSrc, VertexShader); err != nil {
			log.Println("Failed to compile Pixmap vertex shader:", err)
			return
		}
		if err := fShader.Compile(PixmapFragmentShaderSrc, FragmentShader); err != nil {
			log.Println("Failed to compile Pixmap fragment shader:", err)
			return
		}
		if err := PixmapShaderProgram.Link(vShader, fShader); err != nil {
			log.Println("Failed to link Pixmap shader program:", err)
			return
		}
	}

	// Here we generate VAO, VBO and EBO
	gl.GenVertexArrays(1, &p.vao)
	gl.BindVertexArray(p.vao)

	gl.GenBuffers(1, &p.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, p.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(WidgetVertices)*4, gl.Ptr(WidgetVertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &p.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, p.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(WidgetIndices)*4, gl.Ptr(WidgetIndices), gl.STATIC_DRAW)

	// And then we load vertices and indices to OpenGL pipeline.
	vertAttrib := uint32(gl.GetAttribLocation(PixmapShaderProgram.Program(), gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(PixmapShaderProgram.Program(), gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	gl.BindVertexArray(0)

	// When we have already loaded vertices and indices,
	// it's time for uniforms, like texture
	textureUniform := gl.GetUniformLocation(PixmapShaderProgram.Program(), gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	p.ready = true
}

// Draw draws Pixmap's contents to the screen
func (p *Pixmap) Draw() {
	if !p.ready {
		log.Println("Prevented drawing a not ready Pixmap")
		return
	}

	transMatUniform := gl.GetUniformLocation(PixmapShaderProgram.Program(), gl.Str("transMat\x00"))
	gl.UniformMatrix4fv(transMatUniform, 1, false, &p.transMat[0])
	scaleMatUniform := gl.GetUniformLocation(PixmapShaderProgram.Program(), gl.Str("scaleMat\x00"))
	gl.UniformMatrix4fv(scaleMatUniform, 1, false, &p.scaleMat[0])

	// Activate shader
	PixmapShaderProgram.Use()
	//Bind texture
	err := p.texture.Bind()
	if err != nil {
		log.Println("Failed to bind texture while drawing Pixmap:", err)
	}

	// Draw container
	gl.BindVertexArray(p.vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}

// WidgetVertices are default widget vertices
var WidgetVertices = []float32{
	// Positions (X, Y, Z)	// Texture Coords (U, V)
	1.0, 1.0, 0.0, 1.0, 1.0, // Top Right
	1.0, -1.0, 0.0, 1.0, 0.0, // Bottom Right
	-1.0, -1.0, 0.0, 0.0, 0.0, // Bottom Left
	-1.0, 1.0, 0.0, 0.0, 1.0, // Top Left
}

// WidgetIndices are default widget indices
var WidgetIndices = []uint32{
	0, 1, 3,
	1, 2, 3,
}

// PixmapShaderProgram is default shader program for Pixmaps
var PixmapShaderProgram ShaderProgram

// PixmapVertexShaderSrc is default vertex shader source for Pixmaps
var PixmapVertexShaderSrc = `
#version 330 core
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;
uniform mat4 transMat;
uniform mat4 scaleMat;
void main() {
    gl_Position = scaleMat * transMat * vec4(vert, 1.0f);
    fragTexCoord = vec2(vertTexCoord.x, 1.0 - vertTexCoord.y);
}
` + "\x00"

// PixmapFragmentShaderSrc is default fragment shader source for Pixmaps
var PixmapFragmentShaderSrc = `
#version 330 core
in vec2 fragTexCoord;
out vec4 color;
uniform sampler2D tex;
void main() {
    color = texture(tex, fragTexCoord).rgba;
}
` + "\x00"
