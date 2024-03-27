package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type Event interface {
	AddNotifier(notifier notifier.Notifier)
	Trigger(user user.User)
	GetEventName() string
}

type EventFactory struct {
	Notifier []notifier.Notifier
}

func (e *EventFactory) AddNotifier(n notifier.Notifier) {
	e.Notifier = append(e.Notifier, n)
}

func (e *EventFactory) Trigger(user user.User, eventName string) {
	for _, method := range e.Notifier {
		method.Notify(user, eventName)
	}
}
