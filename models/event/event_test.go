package event_test

import (
	"be101_golang/models/constant"
	"be101_golang/models/event"
	"be101_golang/models/language"
	"be101_golang/models/notifier"
	"be101_golang/models/user"
	"testing"
)

func TestEventTrigger(t *testing.T) {
	var flagtests = map[string]struct {
		notifier []notifier.Notifier
		user     user.User
	}{
		constant.Booking: {[]notifier.Notifier{notifier.SMSNotifier{}, notifier.TelegramNotifier{}}, user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}},
		constant.Cancel:  {[]notifier.Notifier{notifier.TelegramNotifier{}}, user.Student{Name: "jonny", PreferredLanguage: language.ZnTW{}}}}

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
