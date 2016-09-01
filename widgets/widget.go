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
	// Usually it's called in AddWidget method of a window, so there is no need
	// to call it manually anywhere else.
	GetReady() error

	// Update updates widget's contents and reacts to events (if there are some).
	// Also, Draw() method is called as one of update steps.
	Update()

	// Draw draws something inside widget's space. By the way, it's called in
	// Update().
	Draw()

	// These two methods operate on a window, a widget is attached to. Should be
	// called in window's AddWidget or something like that.
	SetParentWindow(window.Window)
	ParentWindow() window.Window

	// Some basic methods to contol widget's size and position.
	SetGeometry(g.Rect)
	Geometry() g.Rect
}

// BasicWidget is the simpliest widget type, just a blank one.
// It implements only functions for work with size and position.
// Used to be embedded in another more specific widget struct.
type BasicWidget struct {
	// Pointer to a window, a widget is attached to. If there is none, this pointer
	// should be nil
	parentWindow *window.Window

	// Basic parameters: width, height and X/Y coordinates
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
	// If there is no parent window we just save new rect and return. When the widget
	// will be added to a window (through layout) and GetReady will be called, that
	// geometry will be fully applied.
	if bw.parentWindow == nil || !bw.ready {
		bw.geometry = r
		return
	}
	// If widget already has equal bounding box, we simply return
	if bw.Geometry() == r {
		return
	}

	// We have these variables to translate widget using matrices. Also, we need
	// them even if there should be no translation because of scaling side effect.
	var xTrans, yTrans, zTrans float32

	// Get width and height of the currently used window
	wWidth, wHeight := bw.ParentWindow().Geometry().W, bw.ParentWindow().Geometry().H

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
		// translate back
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

// GetReady initializes matrices and if there was attempt to resize/move widget
// that geometry change is applied here.
func (bw *BasicWidget) GetReady() {
	// Check if window is attached to a window. GetReady assumes SetParentWindow
	// was already called.
	if bw.ParentWindow() == nil {
		return
	}
	// Make sure matrices are initialized
	if blankMat := [16]float32{0}; bw.transMat == blankMat && bw.scaleMat == blankMat {
		bw.transMat, bw.scaleMat = mgl32.Ident4(), mgl32.Ident4()

		// Since widget is not actually resized before adding it to a window, we
		// reset widget's geometry and resize it again.
		var zeroRect g.Rect
		if bw.Geometry() != zeroRect {
			curGeom := bw.Geometry()
			// Reset geometry to a zero-valued one
			bw.geometry = zeroRect
			// Finally set geometry
			bw.SetGeometry(curGeom)
		}
	} else {
		// Matrices are already intialized, so GetReady was called earlier. That
		// usually means the widget was attached to a window and now it's being
		// re-attached.
		// TODO:
		// - Make it able to re-attach widget to a window.
	}
}

// Update does nothing: BasicWidget should not get any events or draw anything
func (bw *BasicWidget) Update() {}

// Draw does nothing: BasicWidget's always blank, we don't draw anything on it
func (bw *BasicWidget) Draw() {}

// SetParentWindow sets a parent window for a widget.
func (bw *BasicWidget) SetParentWindow(w *window.Window) {
	bw.parentWindow = w
}

// ParentWindow returns current parent window of a widget
func (bw *BasicWidget) ParentWindow() *window.Window {
	return bw.parentWindow
}
