package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type Signup struct {
	methods []notifier.Notifier
}

func (s Signup) GetEventName() string {
	return "register"
}

func (s *Signup) AddNotifier(n notifier.Notifier) {
	s.methods = append(s.methods, n)
}

func (s *Signup) Trigger(user user.User) {
	for _, method := range s.methods {
		method.Notify(user, s.GetEventName())
	}
}
