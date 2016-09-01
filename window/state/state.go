// Copyright (c) 2016 Sergey Popov <sergobot@protonmail.com>

package state

// State variables contain possible states of a window.
type State int

// Most usual window states.
const (
	NotInitialized State = iota
	Shown          State = iota
	Hidden         State = iota
)
