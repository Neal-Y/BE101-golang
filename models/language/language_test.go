package language_test

import (
	"be101_golang/models/constant"
	"be101_golang/models/language"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguageCode(t *testing.T) {
	var flagtests = []struct {
		in  language.Language
		out string
	}{
		{language.ZnTW{}, constant.ZhTWCode},
		{language.EnUS{}, constant.EnUSCode},
	}

	for _, tt := range flagtests {
		t.Run(tt.out, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.GetLanguage())
		})
	}
}

func TestGetMessage(t *testing.T) {
	var flagtests = []struct {
		name string
		in   language.Language
		out  string
	}{
		{name: constant.Newyear, in: language.ZnTW{}, out: "新年快樂"},
		{name: constant.Newyear, in: language.EnUS{}, out: "Happy New Year"},
		{name: constant.Register, in: language.ZnTW{}, out: "註冊成功。歡迎加入我們的社群！"},
		{name: constant.Register, in: language.EnUS{}, out: "Registration successful. Welcome to our community!"},
		{name: constant.Booking, in: language.ZnTW{}, out: "課程預定成功。期待見到你！"},
		{name: constant.Booking, in: language.EnUS{}, out: "Course successfully booked. We look forward to seeing you!"},
		{name: constant.Cancel, in: language.ZnTW{}, out: "課程取消成功。希望在其他課程見到你。"},
		{name: constant.Cancel, in: language.EnUS{}, out: "Course successfully cancelled. We hope to see you in other courses."},
	}

	for _, tt := range flagtests {
		t.Run(tt.out, func(t *testing.T) {
			msg, _ := tt.in.GetMessage(tt.name)
			assert.Equal(t, tt.out, msg)
		})
	}
}

// 因為GetMessage()裡面就是使用LanguageCheck()，所以這個測試就不需要了

func TestLanguageCheck(t *testing.T) {
	var flagtests = []struct {
		name      string
		eventName string
		lang      string
		wantMsg   string
		wantErr   bool
	}{
		{"EventNotFound", "nonexistentEvent", constant.ZhTWCode, "", true},
		{"LanguageNotFound", constant.Newyear, "en", "", true},
		{"SuccessfulCheck", constant.Register, constant.ZhTWCode, "註冊成功。歡迎加入我們的社群！", false},
	}

	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := language.LanguageCheck(tt.lang, tt.eventName)
			if tt.wantErr {
				assert.Error(t, err, "")
			} else {
				assert.Equal(t, tt.wantMsg, msg)
			}
		})
	}
}
