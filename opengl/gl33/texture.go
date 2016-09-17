// Copyright (c) Sergey Popov <sergobot@protonmail.com>

package gl33

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	// Packages image/jpeg and image/png are not used explicitly in
	// the code below, but are imported for its initialization side-effect,
	// which allows image.Decode to understand PNG and JPEG formatted images.
	_ "image/jpeg"
	_ "image/png"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Texture struct contains information about an OpenGL texture:
// - OpenGL id of a texture
// - Texture unit the texture is using
// - Its width and height
// Also Texture can load a texture from image.
type Texture struct {
	// OpenGL texture ID
	texture uint32

	// OpenGL texture unit is required if there are more than one texture
	// used in a single shader.
	unit uint32

	// Is true if the texture is ready to be drawn
	imageLoaded bool

	// Size of the texture
	width, height int
}

// LoadFromFile loads texture from an image file.
func (t *Texture) LoadFromFile(file string) error {
	if t.imageLoaded {
		gl.DeleteTextures(1, &t.texture)
	}

	t.imageLoaded = false

	// Load an image from file
	imgFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Failed to load texture from %q: %v", file, err)
	}

	// Decode the newly loaded image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return fmt.Errorf("Failed to decode an image: %v", err)
	}

	// Convert the image to unified format - RGBA
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)

	// Set texture's width and height equal to the image's ones
	t.width, t.height = rgba.Bounds().Dx(), rgba.Bounds().Dy()

	// Now image is loaded and we can use it as an OpenGL texture
	t.imageLoaded = true

	gl.GenTextures(1, &t.texture)

	err = t.Bind()
	if err != nil {
		return err
	}

	// Set texture filtering
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	// Set texture parameters
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	// Create texture
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(t.width),  // t.width and t.height are equal to
		int32(t.height), // rgba.Bounds().Dx and rgba.Bounds().Dy
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	gl.GenerateMipmap(gl.TEXTURE_2D)

	// It's a good practice to unbind once we are done
	err = t.Unbind()

	return err
}

// Ready returns true if image is already loaded and the texture is ready to be used
func (t *Texture) Ready() bool {
	return t.imageLoaded
}

// Bind binds the texture for drawing.
func (t *Texture) Bind() error {
	if !t.imageLoaded {
		return fmt.Errorf("Prevented binding a nonexistent texture")
	}

	gl.ActiveTexture(gl.TEXTURE0 + t.unit)
	gl.BindTexture(gl.TEXTURE_2D, t.texture)
	return nil
}

// Unbind unbinds the texture. Use if drawing is finished.
func (t *Texture) Unbind() error {
	if !t.imageLoaded {
		return fmt.Errorf("Prevented unbinding a nonexistent texture")
	}

	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.ActiveTexture(0)
	return nil
}

// SetUnit sets a unit the texture is using while drawing.
// Required when few textures are used in a single shader.
func (t *Texture) SetUnit(unit uint32) {
	t.unit = unit
}

// Unit returns texture unit the texture is using.
func (t *Texture) Unit() uint32 {
	return t.unit
}
