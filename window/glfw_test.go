// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package rocky

import "testing"

func TestInitGLFW(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestInitGLFW in short mode")
	}
	if err := initGLFW(); err != nil {
		t.Fatal("Failed to intialize GLFW:", err)
	}
}

func TestTerminateGLFW(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestTerminateGLFW in short mode")
	}
	// Usually, when running this test case GLFW is already initialized.
	// However, it's safe to call initGLFW again.
	if err := initGLFW(); err != nil {
		t.Fatal("Failed to intialize GLFW:", err)
	}
	terminateGLFW()
}
