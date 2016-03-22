// Copyright (c) Sergey Popov <sergobot@protonmail.com>
// See LICENSE for details.

package rocky

// WidgetInterface is the interface every widget should implement.
// Widget is a piece of window's space, which has special view
// Widgets are:
// - Buttons
// - Labels
// - Images
// - Text fields
// - And even more!
type WidgetInterface interface {
    GetReady()
    Draw()
    
    SetSize(w, h int)
    Size() (int, int)
    
    SetPosition(x, y int)
    Position() (int, int)
}

// Widget is the simpliest widget type; just a blank widget.
// It implements only functions for work with size and position.
// Used to be inherited by other widgets.
type Widget struct {
    // Basic parameters
    xPos, yPos int
    width, height int
    
    // And nothing more!
}

// SetSize sets the widget's size
func (w *Widget) SetSize(width, height int) {
    w.width, w.height = width, height
}

// Size method returns the widget's current width and height
func (w *Widget) Size() (int, int) {
    return w.width, w.height
}

// SetPosition moves the widget on the window, widget is attached to
func (w* Widget) SetPosition(xPos, yPos int) {
    w.xPos, w.yPos = xPos, yPos
}

// Position method returns the widget's top-left corner current
// coordinates on the window, the widget is attached to
func (w *Widget) Position() (int, int) {
    return w.xPos, w.yPos
}
// GetReady does nothing: widget has nothing to prepare for drawing
func (w *Widget) GetReady() {  }

// Draw does nothing: widget's always blank, we don't draw anything on it
func (w *Widget) Draw() {  }
