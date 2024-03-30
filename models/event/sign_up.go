package event

import (
	"be101_golang/models/constant"
	"be101_golang/models/user"
)

type Signup struct {
	Methods *EventFactory
}

func (s *Signup) Trigger(user user.User) {
	s.Methods.Trigger(user, TriggerHelper(user, constant.Register))
}
