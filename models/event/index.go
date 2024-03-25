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

type Signup struct {
	methods []notifier.Notifier
}

type BookClass struct {
	methods []notifier.Notifier
}

type CancelClass struct {
	methods []notifier.Notifier
}

func (s Signup) GetEventName() string {
	return "register"
}

func (s *Signup) AddNotifier(notifier notifier.Notifier) {
	s.methods = append(s.methods, notifier)
}

func (s *Signup) Trigger(user user.User) {
	for _, notifier := range s.methods {
		if msg, err := user.GetPreferredLanguage().GetMessage(s.GetEventName()); err == nil {
			notifier.Notify(user, msg)
		}
	}
}

func (b *BookClass) GetEventName() string {
	return "booking"
}

func (b *BookClass) AddNotifier(notifier notifier.Notifier) {
	b.methods = append(b.methods, notifier)
}

func (b *BookClass) Trigger(user user.User) {
	for _, notifier := range b.methods {
		if msg, err := user.GetPreferredLanguage().GetMessage(b.GetEventName()); err == nil {
			notifier.Notify(user, msg)
		}
	}
}

func (c *CancelClass) GetEventName() string {
	return "cancellation"
}

func (c *CancelClass) AddNotifier(notifier notifier.Notifier) {
	c.methods = append(c.methods, notifier)
}

func (c *CancelClass) Trigger(user user.User) {
	for _, notifier := range c.methods {
		if msg, err := user.GetPreferredLanguage().GetMessage(c.GetEventName()); err == nil {
			notifier.Notify(user, msg)
		}
	}
}
