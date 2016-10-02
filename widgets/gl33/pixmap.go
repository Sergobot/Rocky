// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package gl33

import (
	"log"

	"github.com/go-gl/gl/v3.3/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/Sergobot/Rocky/opengl/gl33"
)

// Pixmap is one of the simplest widgets, intended to draw raster images
// using OpenGL. It uses Texture struct for image loading and ShaderProgram
// for drawing.
type Pixmap struct {
	// Embed Widget to have access to [Set]Geometry() and some other things like
	// Widget.ready
	Widget

	vao, vbo, ebo uint32
	texture       *gl33.Texture
}

// init is used to initialize pixmap's matrices and currently nothing more.
func (p *Pixmap) init() *Pixmap {
	p.transMat = mgl32.Ident4()
	p.scaleMat = mgl32.Ident4()
	return p
}

// NewPixmap is used to create initialized Pixmap. For now, that's required
// to only have identity matrices instead of null ones.
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
func (p *Pixmap) SetTexture(tex *gl33.Texture) {
	if tex.Ready() {
		p.texture = tex
	} else {
		log.Println("Prevented setting an empty texture to Pixmap")
	}
}

// GetReady initializes the Pixmap to be ready to Draw() function calls.
func (p *Pixmap) GetReady() {
	if p.ready {
		return
	}

	p.ready = false

	if !PixmapShaderProgram.Linked() {
		var vShader, fShader gl33.Shader
		if err := vShader.Compile(PixmapVertexShaderSrc, gl33.VertexShader); err != nil {
			log.Println("Failed to compile Pixmap vertex shader:", err)
			return
		}
		if err := fShader.Compile(PixmapFragmentShaderSrc, gl33.FragmentShader); err != nil {
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
	gl.BufferData(gl.ARRAY_BUFFER, len(widgetVertices)*4, gl.Ptr(widgetVertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &p.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, p.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(widgetIndices)*4, gl.Ptr(widgetIndices), gl.STATIC_DRAW)

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

// PixmapShaderProgram is default shader program for Pixmaps
var PixmapShaderProgram gl33.ShaderProgram

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
