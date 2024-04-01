package notifier_test

import (
	"be101_golang/models/constant"
	"be101_golang/models/notifier"
	"be101_golang/models/user"
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifierGetName(t *testing.T) {
	var flagtests = []struct {
		notifier      notifier.Notifier
		expected_name string
	}{
		{notifier.EmailNotifier{}, constant.Email},
		{notifier.SMSNotifier{}, constant.SMS},
		{notifier.TelegramNotifier{}, constant.Telegram},
		{notifier.LineNotifier{}, constant.Line},
	}

	for _, tt := range flagtests {
		t.Run(tt.expected_name, func(t *testing.T) {
			assert.Equal(t, tt.expected_name, tt.notifier.GetName())
		})
	}
}

func TestNotifierNotify(t *testing.T) {
	var flagtests = []struct {
		name             user.User
		notifier         notifier.Notifier
		message          string
		expected_message string
	}{
		{user.Guest{}, notifier.EmailNotifier{}, "Hello", "Email sent to Guest with message: Hello\n"},
		{user.Student{Name: "neal"}, notifier.SMSNotifier{}, "Hello", "SMS sent to neal with message: Hello\n"},
		{user.Student{Name: "jonny"}, notifier.TelegramNotifier{}, "Hello", "Telegram sent to jonny with message: Hello\n"},
		{user.Student{Name: "yang"}, notifier.LineNotifier{}, "Hello", "Line sent to yang with message: Hello\n"},
	} // 老實說這邊expected_message 是用硬編碼的方式，本來想說可以使用fmt.Sprintf() 來做，但是發現fmt.Sprintf() 也是硬編碼，所以就沒有使用fmt.Sprintf() 來做

	for _, tt := range flagtests {
		t.Run(tt.expected_message, func(t *testing.T) {
			// 捕获输出
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			tt.notifier.Notify(tt.name, tt.message)

			// 关闭写端，完成输出捕获
			w.Close()
			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)
			os.Stdout = oldStdout

			actualMessage := buf.String()
			assert.Contains(t, actualMessage, tt.expected_message)
		})
	}
}
