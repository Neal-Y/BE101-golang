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
	// 目前只允許添加三個Notifier，因為只有三個管道
	if len(e.Notifier) >= 3 {
		return
	}
	e.Notifier[n.GetName()] = n
}

func (e *EventFactory) Trigger(user user.User, eventName string) {
	for _, method := range e.Notifier {
		method.Notify(user, eventName)
	}
}

func TriggerHelper(user user.User, event_name string) (message string) {
	message, err := user.GetPreferredLanguage().GetMessage(event_name)

	if err != nil && message == "" {
		return
	}

	return
}
