package event

import (
	"be101_golang/models/user"
)

type BookClass struct {
	Methods *EventFactory
}

func (b *BookClass) Trigger(user user.User) {
	if message, _ := user.GetPreferredLanguage().GetMessage(Booking); message != "" {
		b.Methods.Trigger(user, message)
	}
}
