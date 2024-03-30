package event

import (
	"be101_golang/models/constant"
	"be101_golang/models/user"
)

type BookClass struct {
	Methods *EventFactory
}

func (b *BookClass) Trigger(user user.User) {
	b.Methods.Trigger(user, TriggerHelper(user, constant.Booking))
}
