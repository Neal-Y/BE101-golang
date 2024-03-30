package event

import (
	"be101_golang/models/constant"
	"be101_golang/models/user"
)

type CancelClass struct {
	Methods *EventFactory
}

func (c *CancelClass) Trigger(user user.User) {
	if message, _ := user.GetPreferredLanguage().GetMessage(constant.Cancel); message != "" {
		c.Methods.Trigger(user, message)
	}
}
