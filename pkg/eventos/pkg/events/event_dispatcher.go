package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	Handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{Handlers: make(map[string][]EventHandlerInterface)}
}

func (ed *EventDispatcher) Register(name string, handler EventHandlerInterface) error {
	if _, ok := ed.Handlers[name]; ok {
		for _, h := range ed.Handlers[name] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.Handlers[name] = append(ed.Handlers[name], handler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.Handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(name string, handler EventHandlerInterface) bool {
	if _, ok := ed.Handlers[name]; ok {
		for _, h := range ed.Handlers[name] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if _, ok := ed.Handlers[event.GetName()]; ok {
		for _, handler := range ed.Handlers[event.GetName()] {
			handler.Handle(event)
		}
	} else {
		return errors.New("No handlers registered for this event")
	}
	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface)  {
	if _, ok := ed.Handlers[eventName]; ok {
        for i, h := range ed.Handlers[eventName] {
            if h == handler {
                ed.Handlers[eventName] = append(ed.Handlers[eventName][:i], ed.Handlers[eventName][i+1:]...)
            }
        }
    }
}
