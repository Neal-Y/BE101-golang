package event

import (
	"be101_golang/models/user"
)

type Signup struct {
	methods *EventFactory
}

func (s *Signup) Trigger(user user.User) {
	s.methods.Trigger(user, Register)
}
