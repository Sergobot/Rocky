// Copyright (c) Sergey Popov <sergobot@protonmail.com>
// See LICENSE for details.

package rocky

// EventType is a custom type to classify events
type EventType int

// Here are all the available (for now) event types
const (
	// NotAnEvent should be always equal to 0, which is default
	// value of Event.Type
	NotAnEvent           EventType = 0
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
	eType    EventType
	receiver EventReceiver
}

// Type returns type of an event
func (be *BasicEvent) Type() EventType {
	return be.eType
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
