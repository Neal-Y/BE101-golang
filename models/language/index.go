package language

import "fmt"

type Language interface {
	GetMessage(event_name string) (string, error)
}

var EventMessage = map[string]map[string]string{
	"registerSuccess": {
		"zh-tw": "註冊成功。歡迎加入我們的社群！",
		"en-us": "Registration successful. Welcome to our community!",
	},
	"bookingSuccess": {
		"zh-tw": "課程預定成功。期待見到你！",
		"en-us": "Course successfully booked. We look forward to seeing you!",
	},
	"cancellationSuccess": {
		"zh-tw": "課程取消成功。希望在其他課程見到你。",
		"en-us": "Course successfully cancelled. We hope to see you in other courses.",
	},
}

type ZnTW struct{}
type EnUS struct{}

func (z ZnTW) GetMessage(event_name string) (string, error) {
	lang := "zh-tw"

	if _, ok := EventMessage[event_name][lang]; !ok {
		return "沒有這個事件訊息", fmt.Errorf("event name %s not found in %s", event_name, lang)
	}

	return EventMessage[event_name][lang], nil
}

func (e EnUS) GetMessage(event_name string) (string, error) {
	lang := "en-us"

	if _, ok := EventMessage[event_name][lang]; !ok {
		return "No such Event or Language", fmt.Errorf("event name %s not found in %s", event_name, lang)
	}

	return EventMessage[event_name][lang], nil
}

/*
{
    "registerSuccess": {
        "en-us": "Registration successful. Welcome to our community!",
        "zh-tw": "註冊成功。歡迎加入我們的社群！"
    },
    "bookingSuccess": {
        "en-us": "Course successfully booked. We look forward to seeing you!",
        "zh-tw": "課程預定成功。期待見到你！"
     },
    "cancellationSuccess": {
        "en-us": "Course successfully cancelled. We hope to see you in other courses.",
        "zh-tw": "課程取消成功。希望在其他課程見到你。"
     },
}
*/
