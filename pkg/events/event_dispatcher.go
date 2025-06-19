package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}
func (eventDispatcher *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if eventDispatcher.Has(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}
	eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName], handler)
	return nil
}

func (eventDispatcher *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	_, ok := eventDispatcher.handlers[eventName]
	if !ok {
		return false
	}
	for _, h := range eventDispatcher.handlers[eventName] {
		if h == handler {
			return true
		}
	}
	return false
}

func (eventDispatcher *EventDispatcher) Clear() {
	eventDispatcher.handlers = make(map[string][]EventHandlerInterface)
}
