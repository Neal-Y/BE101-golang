package event

import (
	"be101_golang/models/user"
)

type CancelClass struct {
	EventFactory
}

func (c CancelClass) GetEventName() string {
	return "cancellation"
}

func (c *CancelClass) Trigger(user user.User) {
	c.EventFactory.Trigger(user, c.GetEventName())
}
