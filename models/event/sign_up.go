package event

import (
	"be101_golang/models/constant"
	"be101_golang/models/user"
)

type Signup struct {
	Methods *EventFactory
}

func (s *Signup) Trigger(user user.User) {
	if message, _ := user.GetPreferredLanguage().GetMessage(constant.Register); message != "" {
		s.Methods.Trigger(user, message)
	}
}
