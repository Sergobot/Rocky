// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package gl33

import (
	"github.com/go-gl/mathgl/mgl32"

	g "github.com/Sergobot/Rocky/geometry"
	"github.com/Sergobot/Rocky/opengl/gl33"
)

// Widget is the simpliest widget type, just a blank one.
// It implements only functions for work with size and position.
// Used to be embedded in another more specific widget struct.
type Widget struct {
	// Basic parameters: width, height and X/Y coordinates
	geometry g.RectF

	// Two matrices: one for scalings, another for translatons
	transMat, scaleMat mgl32.Mat4
}

// SetGeometry sets the rectangle (or bounding box, if you want) of a widget.
// That means, widget will have same coordinates and size as a given rect.
func (w *Widget) SetGeometry(r g.RectF) {
	w.SetSize(r.SizeF)
	w.SetPos(r.PosF)
}

// Geometry returns current bounding box of a widget.
func (w *Widget) Geometry() g.RectF {
	return w.geometry
}

// SetSize sets the widget's size. You can call it manually or through SetGeometry.
func (w *Widget) SetSize(s g.SizeF) {
	xScaleRatio := s.W / w.geometry.W
	yScaleRatio := s.H / w.geometry.H
	zScaleRatio := float32(1) // There is no need to scale Z coordinate

	// This is because of side effect of scaling: our widget moves and we need to
	// move it back.
	// First, retrieve current normilized viewport size
	vpSize := gl33.NormalizedViewportSize()
	// Next convert current, "good" widget coordinates to use in OpenGL
	xCur := (2*w.geometry.X - vpSize.W) / vpSize.W
	yCur := (vpSize.H - 2*w.geometry.Y) / vpSize.H
	// Then we find new ("bad") ones
	xNew := xCur * xScaleRatio
	yNew := yCur * yScaleRatio

	// And, finally, apply transformations to matrices
	w.transMat = w.transMat.Mul4(mgl32.Translate3D(xCur-xNew, yCur-yNew, float32(0)))
	w.scaleMat = w.scaleMat.Mul4(mgl32.Scale3D(xScaleRatio, yScaleRatio, zScaleRatio))

	// Don't forget to update size in the widget itself
	w.geometry.SizeF = s
}

// Size returns current widget's size.
func (w *Widget) Size() g.SizeF {
	return w.geometry.SizeF
}

// SetPos sets position of a widget. It's used in SetGeometry but you can call
// it manually.
func (w *Widget) SetPos(p g.PosF) {
	// Get normalized viewport size
	vpSize := gl33.NormalizedViewportSize()

	// Convert our own normalized coordinates to OpenGL ones and get the difference
	// between new and old coordinates.
	xTrans := (2*p.X - w.geometry.X) / vpSize.W
	yTrans := (2*w.geometry.Y - p.Y) / vpSize.H
	zTrans := float32(0) // We don't tranlsate z coordinate

	// Then multiply widget's translation matrix and one returned by 'Translate3D'
	w.transMat = w.transMat.Mul4(mgl32.Translate3D(xTrans, yTrans, zTrans))

	// And, finally, we assign widget's position to the new one
	w.geometry.PosF = p
}

// Pos returns current widget's position.
func (w *Widget) Pos() g.PosF {
	return w.geometry.PosF
}

// GetReady initializes matrices and if there was attempt to resize/move widget
// that geometry change is applied here.
func (w *Widget) GetReady() {}

// Update does nothing: Widget should not get any events or draw anything
func (w *Widget) Update() {}

// Draw does nothing: Widget's always blank, we don't draw anything on it
func (w *Widget) Draw() {}

// widgetVertices are default widget vertices, used in more advanced widget than
// the one implemented in this file
var widgetVertices = []float32{
	// Positions (X, Y, Z)	// Texture Coords (U, V)
	1.0, 1.0, 0.0, 1.0, 1.0, // Top Right
	1.0, -1.0, 0.0, 1.0, 0.0, // Bottom Right
	-1.0, -1.0, 0.0, 0.0, 0.0, // Bottom Left
	-1.0, 1.0, 0.0, 0.0, 1.0, // Top Left
}

// widgetIndices are default widget indices, used in more advanced widget than
// the one implemented in this file
var widgetIndices = []uint32{
	0, 1, 3,
	1, 2, 3,
}
