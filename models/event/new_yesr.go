package event

import (
	"be101_golang/models/constant"
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type NewYear struct {
	Methods *EventFactory
}

func (n *NewYear) AddNotifier(notifier notifier.Notifier) {
	n.Methods.AddNotifier(notifier)
}

func (n *NewYear) Trigger(user user.User) {
	n.Methods.Trigger(user, TriggerHelper(user, constant.Newyear))
}
