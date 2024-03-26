package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type CancelClass struct {
	methods []notifier.Notifier
}

func (c CancelClass) GetEventName() string {
	return "cancellation"
}

func (c *CancelClass) AddNotifier(n notifier.Notifier) {
	c.methods = append(c.methods, n)
}

func (c *CancelClass) Trigger(user user.User) {
	for _, method := range c.methods {
		method.Notify(user, c.GetEventName())
	}
}
