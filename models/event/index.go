package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type Event interface {
	AddNotifier(notifier notifier.Notifier)
	Trigger(user user.User)
}

type Signup struct {
	methods   []notifier.Notifier
	studentID string
}

type BookClass struct {
	methods []notifier.Notifier
	classID string
}

type CancelClass struct {
	methods []notifier.Notifier
	classID string
}

func (s Signup) AddNotifier(notifier notifier.Notifier) {
	s.methods = append(s.methods, notifier)
}
func (s Signup) Trigger(user user.User) {
	for _, notifier := range s.methods{
		notifier.Notify(user, "Welcome to the class!")
	}
}
//Notify(user user.User, message string) 
func (b BookClass) AddNotifier(notifier notifier.Notifier) {
	b.methods = append(b.methods, notifier)
}
func (b BookClass) Trigger(user user.User) {
	for _, notifier := range b.methods{
		notifier.Notify(user, "Class booked!")
}

func (c CancelClass) AddNotifier(notifier notifier.Notifier) {
	c.methods = append(c.methods, notifier)
}
func (c CancelClass) Trigger(user user.User) {
	for _, notifier := range c.methods{
		notifier.Notify(user, "Class cancelled!")
	}
}
