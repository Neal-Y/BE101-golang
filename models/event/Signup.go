package event

import (
	"be101_golang/models/user"
)

type Signup struct {
	EventFactory
}

func (s Signup) GetEventName() string {
	return "register"
}

func (s *Signup) Trigger(user user.User) {
	s.EventFactory.Trigger(user, s.GetEventName())
}
