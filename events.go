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
	Key      glfw.Key
	Scancode int
	Action   glfw.Action
	Mods     glfw.ModifierKey
}

// Type returns type of an event. For KeyEvent it is KeyEventType.
func (ke *KeyEvent) Type() EventType {
	return KeyEventType
}

// Fill wraps raw GLFW data in a single Event struct
func (ke *KeyEvent) Fill(k glfw.Key, s int, a glfw.Action, m glfw.ModifierKey) {
	ke.Key = k
	ke.Scancode = s
	ke.Action = a
	ke.Mods = m
}

// MouseButtonEvent is a wrapper for data of GLFW's mouse button callback function
// (MouseButtonCallback(...)). It contains:
// - Button (Left, Middle etc)
// - Action (Pressed or Released)
// - Modifiers (Shift, Alt and so on)
type MouseButtonEvent struct {
	BasicEvent
	Button glfw.MouseButton
	Action glfw.Action
	Mods   glfw.ModifierKey
}

// Type returns type of an event. For MouseButtonEvent it is MouseButtonEventType.
func (mbe *MouseButtonEvent) Type() EventType {
	return MouseButtonEventType
}

// Fill wraps raw GLFW data in a single Event struct
func (mbe *MouseButtonEvent) Fill(b glfw.MouseButton, a glfw.Action, m glfw.ModifierKey) {
	mbe.Button = b
	mbe.Action = a
	mbe.Mods = m
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
