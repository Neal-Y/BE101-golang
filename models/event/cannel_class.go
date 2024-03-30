package event

import (
	"be101_golang/models/constant"
	"be101_golang/models/user"
)

type CancelClass struct {
	Methods *EventFactory
}

func (c *CancelClass) Trigger(user user.User) {
	c.Methods.Trigger(user, TriggerHelper(user, constant.Cancel))
}
