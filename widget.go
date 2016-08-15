// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package rocky

import (
	"github.com/go-gl/mathgl/mgl32"

	g "github.com/Sergobot/Rocky/geometry"
	"github.com/Sergobot/Rocky/window"
)

// Widget is an interface for objects, each of which a separate piece of
// window's space. There are many possible examples of widgets:
// - Buttons
// - Labels
// - Images
// - Text inputs
// - And even more!
type Widget interface {
	// GetReady does everything about initialization before drawing.
	// It's called only once, when widget is going to be drawn for the first time.
	GetReady() error

	// Update updates widget's contents and reacts to events (if there are some).
	// Also, Draw() method is called as one of update steps.
	Update()

	// Draw draws something inside widget's space. By the way, it's called in
	// Update().
	Draw()

	// Some basic methods to contol widget's size and position.
	SetGeometry(g.Rect)
	Geometry() g.Rect
}

// BasicWidget is the simpliest widget type, just a blank one.
// It implements only functions for work with size and position.
// Used to be embedded in another more specific widget struct.
type BasicWidget struct {
	// Basic parameters
	geometry g.Rect

	// This matrix is responsible for widget moving
	transMat mgl32.Mat4
	// And this one - for widget resizing/scaling
	scaleMat mgl32.Mat4

	// Is true only if the widget is ready to be drawn
	ready bool
}

// SetGeometry sets the rectangle (or bounding box, if you want) of a widget.
// That means, widget will have same coordinates and size as a given rect.
func (bw *BasicWidget) SetGeometry(r g.Rect) {
	// If widget already has equal bounding box, we simply return
	if bw.Geometry() == r {
		return
	}

	// We have these variables to translate widget using matrices. Also, we need
	// them even if there should be no translation because of scaling side effect.
	var xTrans, yTrans, zTrans float32

	// Get width and height of the currently used window
	win := window.Get()
	wWidth, wHeight := win.Geometry().W, win.Geometry().H

	// First we apply size changes, if there are any
	if r.Size != bw.Geometry().Size {
		xScaleRatio := float32(r.W) / float32(bw.Geometry().W)
		yScaleRatio := float32(r.H) / float32(bw.Geometry().H)
		zScaleRatio := float32(1)

		bw.scaleMat = bw.scaleMat.Mul4(mgl32.Scale3D(xScaleRatio, yScaleRatio, zScaleRatio))

		// Then we apply pos changes: here and in the outer else block.
		// That's because of size effect of scaling: our widget moves and we need
		// to move it back.
		// First of all we convert current, "good" widget coordinates to use in OpenGL
		xCur := float32(2*bw.Geometry().X-wWidth) / float32(wWidth)
		yCur := float32(wHeight-2*bw.Geometry().Y) / float32(wHeight)

		// If pos should be kept the same, we find new "bad" coordinates and
		// translate
		xBad := xCur * xScaleRatio
		yBad := yCur * yScaleRatio

		// Handle case, when both size and pos are updated
		if r.Pos != bw.Geometry().Pos {
			xNew, yNew := float32(r.X), float32(r.Y)
			xTrans, yTrans = xNew-xBad, yNew-yBad
		} else {
			xTrans, yTrans = xCur-xBad, yCur-yBad
		}
	} else {
		// Changes affect only Pos
		xTrans = 2 * float32(r.X-bw.Geometry().X) / float32(wWidth)
		yTrans = 2 * float32(bw.Geometry().Y-r.Y) / float32(wHeight)
		zTrans = float32(0) // We don't tranlsate z coordinate
	}

	// And, finally, apply transformations to matrices
	bw.transMat = bw.transMat.Mul4(mgl32.Translate3D(xTrans, yTrans, zTrans))
}

// Geometry returns current bounding box of a widget.
func (bw *BasicWidget) Geometry() g.Rect {
	return bw.geometry
}

// GetReady does nothing: BasicWidget has nothing to prepare for drawing
func (bw *BasicWidget) GetReady() {}

// Update does nothing: BasicWidget should not get any events or draw anything
func (bw *BasicWidget) Update() {}

// Draw does nothing: BasicWidget's always blank, we don't draw anything on it
func (bw *BasicWidget) Draw() {}
