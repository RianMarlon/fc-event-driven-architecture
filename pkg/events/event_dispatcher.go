package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (eventDispatcher *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := eventDispatcher.handlers[event.GetName()]; ok {
		maxWorkers := len(handlers)
		wg := sync.WaitGroup{}
		wg.Add(maxWorkers)
		for _, handler := range handlers {
			go func() {
				handler.Handle(event)
				defer wg.Done()

			}()
		}
		wg.Wait()
	}
	return nil
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

func (eventDispatcher *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	for i, h := range eventDispatcher.handlers[eventName] {
		if h == handler {
			eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName][:i], eventDispatcher.handlers[eventName][i+1:]...)
			return nil
		}
	}
	return errors.New("handler not registered")
}

func (eventDispatcher *EventDispatcher) Clear() error {
	eventDispatcher.handlers = make(map[string][]EventHandlerInterface)
	return nil
}
