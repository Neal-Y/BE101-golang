package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

const (
	Cancel   = "cancellation"
	Booking  = "booking"
	Register = "register"
)

type Event interface {
	AddNotifier(notifier notifier.Notifier)
	Trigger(user user.User)
}

type EventFactory struct {
	Notifier []notifier.Notifier
}

func (e *EventFactory) AddNotifier(n notifier.Notifier) {
	// 目前只允許添加三個Notifier，因為只有三個管道
	if len(e.Notifier) >= 3 {
		return
	}

	// 確保沒有重複
	for _, existingNotifier := range e.Notifier {
		if existingNotifier == n {
			return
		}
	}

	e.Notifier = append(e.Notifier, n)
}

func (e *EventFactory) Trigger(user user.User, eventName string) {
	for _, method := range e.Notifier {
		method.Notify(user, eventName)
	}
}
