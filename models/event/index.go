package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type Event interface {
	AddNotifier(notifier notifier.Notifier)
	Trigger(user user.User)
}

type EventFactory struct {
	Notifier map[string]notifier.Notifier
}

func NewEventFactory() *EventFactory {
	return &EventFactory{
		Notifier: make(map[string]notifier.Notifier),
	}
}

// map is key-value pair, so we don't need to concern duplicate key
func (e *EventFactory) AddNotifier(n notifier.Notifier) {
	e.Notifier[n.GetName()] = n
}

func (e *EventFactory) Trigger(user user.User, eventMessage string) {
	if eventMessage == "" {
		return
	}
	for _, method := range e.Notifier {
		method.Notify(user, eventMessage)
	}
}

func TriggerHelper(user user.User, eventName string) (message string) {
	message, err := user.GetPreferredLanguage().GetMessage(eventName)

	if err != nil && message == "" {
		return
	}

	return
}
