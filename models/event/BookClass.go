package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type BookClass struct {
	methods []notifier.Notifier
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
