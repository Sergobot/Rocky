// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package opengl

import "github.com/Sergobot/Rocky/opengl/gl33"

// Texture is an interface for all the Texture struct in gl**/ subfolders.
// These struct help to assumed to manage single texture life:
// - Load from file
// - Check readiness for rendering
// - Bind/Unbind before/after rendering
// - Set texture unit to use multiple textures in one shader
// - Tell version of OpenGL it uses.
// Nothing more is required for now, so there is actually nothing more.
type Texture interface {
	// LoadFromFile loads texture from an image file
	LoadFromFile(string) error

	// These two methods manage binding texture for rendering
	Bind() error
	Unbind() error

	// Use these two methods to set texture unit to use multiple textures in a
	// single shader
	SetUnit(uint32)
	Unit() uint32

	// If texture is successfully loaded, Ready() will return true
	Ready() bool

	// Returns OpenGL version texture is using. Usually called in widgets to check
	// if an improper texture was passed.
	Version() string
}

// NewTexture creates a new texture, according to OpenGL version initialized.
// If there is none, it will return nil.
func NewTexture() Texture {
	if gl33.Initialized() {
		return new(gl33.Texture)
	}

	return nil
}
