// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package rocky

import "fmt"

// Layout is an interface for objects, responsible for making widgets look nice
// together. Also, layouts should save developers from pain of setting pixel sizes
// to widgets. There are some basic layout examples:
// - VLayout - vertically arranges widgets one next another on screen,
// - HLayout - horizontally arranges widgets one next another on screen.
type Layout interface {
	// Obvious methods to add/remove widgets
	AddWidget(Widget)
	RemoveWidget(Widget) error

	// Sets bounding box for a layout. Layout can't grow bigger than that Rect.
	SetGeometry(Rect)
	Geometry() Rect

	// Makes Layout to update widgets' sizes. Usually, it's called automatically
	// but you can override that.
	Activate()

	// Returns slice of all the widgets attached to a layout. May be required when
	// drawing.
	Widgets() []Widget
}

// BasicLayout is a very basic layout struct, used to be embedded in other layouts.
type BasicLayout struct {
	// Bounding box of a layout
	geometry Rect

	widgets []Widget
}

// AddWidget adds a widget to a layout.
func (bl *BasicLayout) AddWidget(w Widget) {
	w.GetReady()
	bl.widgets = append(bl.widgets, w)

	bl.Activate()
}

// RemoveWidget removes a widget from a layout.
func (bl *BasicLayout) RemoveWidget(w Widget) error {
	removed := false
	for i, v := range bl.widgets {
		if w == v {
			bl.widgets = append(bl.widgets[:i], bl.widgets[i+1:]...)
			removed = true
		}
	}
	if !removed {
		return fmt.Errorf("Widget not found in layout")
	}

	bl.Activate()

	return nil
}

// SetGeometry sets geometry (bounding box) of a layout.
func (bl *BasicLayout) SetGeometry(r Rect) {
	bl.geometry = r
	bl.Activate()
}

// Geometry returns geometry (bounding box) of a layout.
func (bl *BasicLayout) Geometry() Rect {
	return bl.geometry
}

// Activate is called to recalculate size of widgets in a layout. You usually don't
// need to call this if you don't reimplement AddWidget/RemoveWidget methods.
func (bl *BasicLayout) Activate() {
	// Nothing, since this is only a basic layout, which's used to be embedded in
	// a more specific one.
}

// Widgets returns slice of widgets attached to a layout.
func (bl *BasicLayout) Widgets() []Widget {
	return bl.widgets
}
