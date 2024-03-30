package event

import (
	"be101_golang/models/constant"
	"be101_golang/models/user"
)

type NewYear struct {
	Methods *EventFactory
}

func (n *NewYear) Trigger(user user.User) {
	message, err := user.GetPreferredLanguage().GetMessage(constant.Newyear)
	if err != nil {
		return
	}

	n.Methods.Trigger(user, message)
}
