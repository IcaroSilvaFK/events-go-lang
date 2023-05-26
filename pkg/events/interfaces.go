package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface) error
}

type EventDispatcherInterface interface {
	Register(name string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(name string, handler EventHandlerInterface) error
	Has(name string, handler EventHandlerInterface) bool
	Clear() error
}
