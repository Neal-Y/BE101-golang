package event

import (
	"be101_golang/models/user"
)

type CancelClass struct {
	methods *EventFactory
}

func (c CancelClass) GetEventName() string {
	return Cancel
}

func (c *CancelClass) Trigger(user user.User) {
	c.methods.Trigger(user, c.GetEventName())
}
