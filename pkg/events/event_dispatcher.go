package events

import (
	"errors"
)

var (
	ErrorHandlerAlreadyRegistered = errors.New("handler already registered")
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(name string, handler EventHandlerInterface) error {

	if _, ok := ed.handlers[name]; ok {
		for _, h := range ed.handlers[name] {
			if handler == h {
				return ErrorHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[name] = append(ed.handlers[name], handler)

	return nil

}
