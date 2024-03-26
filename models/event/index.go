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

func (s *Signup) AddNotifier(n notifier.Notifier) {
	s.methods = append(s.methods, n)
}

func (s *Signup) Trigger(user user.User) {
	for _, method := range s.methods {
		method.Notify(user, s.GetEventName())
	}
}

func (b BookClass) GetEventName() string {
	return "booking"
}

func (b *BookClass) AddNotifier(n notifier.Notifier) {
	b.methods = append(b.methods, n)
}

func (b *BookClass) Trigger(user user.User) {
	for _, method := range b.methods {
		method.Notify(user, b.GetEventName())
	}
}

func (c CancelClass) GetEventName() string {
	return "cancellation"
}

func (c *CancelClass) AddNotifier(n notifier.Notifier) {
	c.methods = append(c.methods, n)
}

func (c *CancelClass) Trigger(user user.User) {
	for _, method := range c.methods {
		method.Notify(user, c.GetEventName())
	}
}
