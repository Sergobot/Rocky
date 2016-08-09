// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package rocky

// Point represents a point in 2D space with X and Y coordinates
type Point struct {
	X, Y int
}

// Size represents size of a 2D object, like image or widget
type Size struct {
	W, H int
}

// Rect represents a rectangular structure, which has both size and coordinates
type Rect struct {
	Point
	Size
}
