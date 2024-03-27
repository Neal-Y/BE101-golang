package event

import (
	"be101_golang/models/user"
)

type BookClass struct {
	methods *EventFactory
}

func (b BookClass) GetEventName() string {
	return "booking"
}

func (b *BookClass) Trigger(user user.User) {
	b.methods.Trigger(user, b.GetEventName())
}
