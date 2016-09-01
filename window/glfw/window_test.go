// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package glfw

import "testing"

// Tests if a newly created window is show-able
func TestWindowShow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestWindowShow in short mode")
	}
	var w Window
	w.Show()
}

// Tests if a window created explicitly using Create() is show-able
func TestWindowCreateShow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestWindowCreateShow in short mode")
	}
	var w Window
	w.create()
	w.Show()

	w.Destroy()
}

// Tests if Window is able to re-init GLFW and show itself
func TestWindowShowWithGLFWTerminated(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestWindowShowWithGLFWTerminated in short mode")
	}
	var w Window
	terminateGLFW()
	w.Show()
}
