// Copyright (c) Sergey Popov <sergobot@protonmail.com>
// See LICENSE for details.

package rocky

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

// EventType is a custom type to classify events
type EventType int

// Here are all the available (for now) event types
const (
	// NotAnEvent should be always equal to -1 to easily recognize empty events
	NotAnEvent           EventType = -1
	KeyEventType         EventType = iota
	MouseButtonEventType EventType = iota
	CursorEventType      EventType = iota
	ScrollEventType      EventType = iota
)

// EventReceiver is an interface, responsible for processing Events.
// There are lots of different possible struct, which may implement this enterface,
// for example: widgets, windows, audio players and so on.
type EventReceiver interface {
	ProcessEvent(e Event)
}

// Event is a base interface for all the events. You may create your own event struct
// and add more EventType's.
//
// All the receivers are getting notified using their .ProcessEvent(e Event) method,
// which is called in Event.Process().
type Event interface {
	// Type returns type of an event
	Type() EventType
	// SetReceiver sets receiver for an event. Receiver is used in Event.Process() to process
	// the event
	SetReceiver(er EventReceiver)
	// Receiver returns event's receiver
	Receiver() EventReceiver
	// Process should call some method of the event's receiver to process the event
	Process()
}

// BasicEvent is a base struct for all the other events. It deals with routine like
// storing event's type, receiver and setting/getting these.
type BasicEvent struct {
	receiver EventReceiver
}

// Type returns type of an event. Since BasicEvent doesn't contain any data except
// of event receiver, it's actually not an Event.
func (be *BasicEvent) Type() EventType {
	return NotAnEvent
}

// SetReceiver sets receiver for an event
func (be *BasicEvent) SetReceiver(er EventReceiver) {
	be.receiver = er
}

// Receiver method returns event's receiver
func (be *BasicEvent) Receiver() EventReceiver {
	return be.receiver
}

// Process is juts an alias for Event.receiver.ProcessEvent(Event)
func (be *BasicEvent) Process() {
	be.receiver.ProcessEvent(be)
}

// KeyEvent is a wrapper for data of GLFW's key callback function (KeyCallback(...)).
// It contains:
// - Key
// - Scancode
// - Action (Pressed, Released and so on)
// - Modifiers (Shift, Alt and so on).
type KeyEvent struct {
	BasicEvent
	key      glfw.Key
	scancode int
	action   glfw.Action
	mods     glfw.ModifierKey
}

// Type returns type of an event. For KeyEvent it is KeyEventType.
func (ke *KeyEvent) Type() EventType {
	return KeyEventType
}

// SetKey method sets KeyEvent's .key
func (ke *KeyEvent) SetKey(k glfw.Key) {
	ke.key = k
}

// Key method of KeyEvent struct returns its .key
func (ke *KeyEvent) Key() glfw.Key {
	return ke.key
}

// SetScancode method sets KeyEvent's .scancode
func (ke *KeyEvent) SetScancode(s int) {
	ke.scancode = s
}

// Scancode method of KeyEvent struct returns its .scancode
func (ke *KeyEvent) Scancode() int {
	return ke.scancode
}

// SetAction method sets KeyEvent's .action
func (ke *KeyEvent) SetAction(a glfw.Action) {
	ke.action = a
}

// Action method of KeyEvent struct returns its .action
func (ke *KeyEvent) Action() glfw.Action {
	return ke.action
}

// SetMods method sets KeyEvent's .mods
func (ke *KeyEvent) SetMods(m glfw.ModifierKey) {
	ke.mods = m
}

// Mods method of KeyEvent struct returns its .mods
func (ke *KeyEvent) Mods() glfw.ModifierKey {
	return ke.mods
}

// Fill wraps raw GLFW data in a single Event struct
func (ke *KeyEvent) Fill(k glfw.Key, s int, a glfw.Action, m glfw.ModifierKey) {
	ke.SetKey(k)
	ke.SetScancode(s)
	ke.SetAction(a)
	ke.SetMods(m)
}

// MouseButtonEvent is a wrapper for data of GLFW's mouse button callback function
// (MouseButtonCallback(...)). It contains:
// - Button (Left, Middle etc)
// - Action (Pressed or Released)
// - Modifiers (Shift, Alt and so on)
type MouseButtonEvent struct {
	BasicEvent
	button glfw.MouseButton
	action glfw.Action
	mods   glfw.ModifierKey
}

// Type returns type of an event. For MouseButtonEvent it is MouseButtonEventType.
func (mbe *MouseButtonEvent) Type() EventType {
	return MouseButtonEventType
}

// SetButton method sets MouseButtonEvent's .button
func (mbe *MouseButtonEvent) SetButton(b glfw.MouseButton) {
	mbe.button = b
}

// Button method of MouseButtonEvent struct returns its .button
func (mbe *MouseButtonEvent) Button() glfw.MouseButton {
	return mbe.button
}

// SetAction method sets MouseButtonEvent's .action
func (mbe *MouseButtonEvent) SetAction(a glfw.Action) {
	mbe.action = a
}

// Action method of MouseButtonEvent struct returns its .action
func (mbe *MouseButtonEvent) Action() glfw.Action {
	return mbe.action
}

// SetMods method sets MouseButtonEvent's .mods
func (mbe *MouseButtonEvent) SetMods(m glfw.ModifierKey) {
	mbe.mods = m
}

// Mods method of MouseButtonEvent struct returns its .mods
func (mbe *MouseButtonEvent) Mods() glfw.ModifierKey {
	return mbe.mods
}

// Fill wraps raw GLFW data in a single Event struct
func (mbe *MouseButtonEvent) Fill(b glfw.MouseButton, a glfw.Action, m glfw.ModifierKey) {
	mbe.SetButton(b)
	mbe.SetAction(a)
	mbe.SetMods(m)
}

// EventQueue strust is used to store events in queue. It has 2 methods:
// - PushEvent() - used to send an event
// - PullEvent() - used to get an event
type EventQueue struct {
	events []Event
}

// PushEvent appends an event to a queue
func (eq *EventQueue) PushEvent(e Event) {
	eq.events = append(eq.events, e)
}

// PullEvent removes the first event in a queue and returns it. If the queue
// is empty, it returns nil
func (eq *EventQueue) PullEvent() Event {
	if len(eq.events) > 0 {
		e := eq.events[0]
		eq.events = append(eq.events[:0], eq.events[1:]...)
		return e
	}

	return nil
}
