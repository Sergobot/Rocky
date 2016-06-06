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

// EventReciever is an interface, responsible for processing Events.
// There are lots of different possible struct, which may implement this enterface,
// for example: widgets, windows, audio players and so on.
type EventReciever interface {
	ProcessEvent(e Event)
}

// Event is a base interface for all the events. You may create your own event struct
// and add more EventType's.
//
// All the recievers are getting notified using their .ProcessEvent(e Event) method,
// which is called in Event.Process().
type Event interface {
	// Type returns type of an event
	Type() EventType
	// SetReciever sets reciever for an event. Reciever is used in Event.Process() to process
	// the event
	SetReciever(er EventReciever)
	// Reciever returns event's reciever
	Reciever() EventReciever
	// Process should call some method of the event's reciever to process the event
	Process()
}

// BasicEvent is a base struct for all the other events. It deals with routine like
// storing event's type, reciever and setting/getting these.
type BasicEvent struct {
	eType    EventType
	reciever EventReciever
}

// Type returns type of an event
func (be *BasicEvent) Type() EventType {
	return be.eType
}

// SetReciever sets reciever for an event
func (be *BasicEvent) SetReciever(er EventReciever) {
	be.reciever = er
}

// Reciever method returns event's reciever
func (be *BasicEvent) Reciever() EventReciever {
	return be.reciever
}

// Process is juts an alias for Event.reciever.ProcessEvent(Event)
func (be *BasicEvent) Process() {
	be.reciever.ProcessEvent(be)
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
