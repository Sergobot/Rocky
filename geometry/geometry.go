// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package geometry

// Point represents a point in 2D space with integral X and Y coordinates
type Point struct {
	X, Y int
}

// PointF represents a point in 2D space with fractional X and Y coordinates
type PointF struct {
	X, Y float32
}

// Size represents integral size of a 2D object, like image or widget
type Size struct {
	W, H int
}

// SizeF represents fractional size of a 2D object, like image or widget
type SizeF struct {
	W, H float32
}

// Pos makes more sense than Point in context of rects
type Pos Point

// PosF makes more sense than PointF in context of rects
type PosF PointF

// Rect represents a rectangular structure, which has both integral size and
// coordinates of its top-left corner.
type Rect struct {
	Pos
	Size
}

// RectF represents a rectangular structure, which has both fractional size and
// coordinates of its top-left corner
type RectF struct {
	PosF
	SizeF
}
