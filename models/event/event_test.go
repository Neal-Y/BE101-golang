package event_test

import (
	"be101_golang/models/constant"
	"be101_golang/models/event"
	"be101_golang/models/language"
	"be101_golang/models/notifier"
	"be101_golang/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventTrigger(t *testing.T) {
	var flagtests = map[string]struct {
		notifier []notifier.Notifier
		user     user.User
	}{
		constant.Booking: {[]notifier.Notifier{notifier.SMSNotifier{}, notifier.TelegramNotifier{}}, user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}},
		constant.Cancel:  {[]notifier.Notifier{notifier.TelegramNotifier{}}, user.Student{Name: "jonny", PreferredLanguage: language.ZnTW{}}},
	}

	for name, tt := range flagtests {
		t.Run(name, func(t *testing.T) {
			tmp := event.NewEventFactory()
			for _, notifier := range tt.notifier {
				tmp.AddNotifier(notifier)
			}
			tmp.Trigger(tt.user, event.TriggerHelper(tt.user, name))
		})
	}
}

func TestEvent(t *testing.T) {
	eventContainer := event.NewEventFactory()
	eventContainer.AddNotifier(notifier.SMSNotifier{})

	testNotifier := notifier.TelegramNotifier{}

	var flagtests = map[string]struct {
		event event.Event
		user  user.User
	}{
		constant.Register: {&event.Signup{Methods: eventContainer}, user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}},
		constant.Booking:  {&event.BookClass{Methods: eventContainer}, user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}},
		constant.Cancel:   {&event.CancelClass{Methods: eventContainer}, user.Student{Name: "jonny", PreferredLanguage: language.ZnTW{}}},
		constant.Newyear:  {&event.NewYear{Methods: eventContainer}, user.Student{Name: "jonny", PreferredLanguage: language.ZnTW{}}},
	}

	for _, tt := range flagtests {
		t.Run("", func(t *testing.T) {
			tt.event.AddNotifier(testNotifier)
			tt.event.Trigger(tt.user)
		})
	}
}

func TestEventFactoryEarlyReturn(t *testing.T) {
	user := user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}
	errorEventName := ""

	event.TriggerHelper(user, "")

	assert.Empty(t, event.TriggerHelper(user, errorEventName))
}
