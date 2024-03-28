package event

import (
	"be101_golang/models/user"
)

type BookClass struct {
	methods *EventFactory
}

func (b *BookClass) Trigger(user user.User) {
	b.methods.Trigger(user, Booking)
}
