package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type Event interface {
	AddNotifier(notifier notifier.Notifier)
	Trigger(user user.User, message string)
	GetEventName() string
}

type BaseEvent struct {
	methods []notifier.Notifier
}

type Signup struct {
	BaseEvent
}

type BookClass struct {
	BaseEvent
}

type CancelClass struct {
	BaseEvent
}

/*
	signupEvent := Signup{}
	signupEvent.AddNotifier(emailNotifier)
	signupEvent.AddNotifier(smsNotifier)
	signupEvent.Trigger(user, signupEvent.GetEventName())

	or

	signupEvent.Trigger(user)
*/

func (b *BaseEvent) AddNotifier(notifier notifier.Notifier) {
	b.methods = append(b.methods, notifier)
}

func (b *BaseEvent) Trigger(user user.User, message string) {
	for _, notifier := range b.methods {
		if msg, err := user.GetPreferredLanguage().GetMessage(message); err == nil {
			notifier.Notify(user, msg)
		}
	}
}

func (s Signup) GetEventName() string {
	return "register"
}

func (b BookClass) GetEventName() string {
	return "booking"
}

func (c CancelClass) GetEventName() string {
	return "cancellation"
}
