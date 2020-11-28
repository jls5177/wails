package interfaces

import "github.com/jls5177/wails/lib/messages"

// EventManager is the event manager interface
type EventManager interface {
	PushEvent(*messages.EventData)
	Emit(eventName string, optionalData ...interface{})
	On(eventName string, callback func(...interface{}))
	Start(Renderer)
	Shutdown()
}
