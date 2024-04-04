package event

import (
	"be101_golang/models/constant"
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type BookClass struct {
	Methods *EventFactory
}

func (b *BookClass) AddNotifier(notifier notifier.Notifier) {
	b.Methods.AddNotifier(notifier)
}

func (b *BookClass) Trigger(user user.User) {
	b.Methods.Trigger(user, TriggerHelper(user, constant.Booking))
}
