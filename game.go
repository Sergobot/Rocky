// Copyright (c) Sergey Popov <sergobot@protonmail.com>
// See LICENSE for details.

package rocky

// Game struct is responsible for event queuing. Also, it's planned to add these
// features:
// - Command line args handling
// - Storing all the necessary variables (windows, settings and so on)
// - And even more! But later... :)
type Game struct {
	eq EventQueue

	shouldQuit bool
}

// Terminate terminates game. It should power off all the Rocky stuff, i.e. windows,
// widgets and so on. But for now it does none of these, except of setting g.shouldQuit
// to true.
func (g *Game) Terminate() {
	// TODO:
	// Do all the termination stuff here (close all the windows, empty event queue etc.)
	g.shouldQuit = true
}

// ShouldQuit returns true if g.Terminate was already called and Rocky has to stop
// everything it handles.
func (g *Game) ShouldQuit() bool {
	return g.shouldQuit
}

// Exec executes event loop (it's Game.loop() by default), in which all the events are sent to their receivers
func (g *Game) Exec() {
	go g.loop()
}

// loop is default event processing loop. It just launches goroutines for each event
// in queue
func (g *Game) loop() {
	for !g.ShouldQuit() {
		if e := g.eq.PullEvent(); e != nil && e.Type() != NotAnEvent {
			go e.Process()
		}
	}
}
