package events

import (
	"errors"
	"sync"
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

func (ed *EventDispatcher) Clear() error {

	ed.handlers = make(map[string][]EventHandlerInterface)

	return nil
}

func (ed *EventDispatcher) Has(name string, handler EventHandlerInterface) bool {

	if handlers, ok := ed.handlers[name]; ok {
		for _, h := range handlers {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {

	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, h := range handlers {
			wg.Add(1)
			go h.Handle(event, wg)
		}
		wg.Wait() //! wait for all handlers to finish
	}

	return nil
}

func (ed *EventDispatcher) Remove(name string, handler EventHandlerInterface) error {

	if handlers, ok := ed.handlers[name]; ok {
		for i, h := range handlers {
			if h == handler {
				ed.handlers[name] = append(handlers[:i], handlers[i+1:]...)
				return nil
			}
		}
	}

	return nil
}
