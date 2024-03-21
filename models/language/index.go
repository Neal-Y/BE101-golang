package language

type Language interface {
	GetMessage(event_name string) string
}

var ZnTWEventMessage = map[string]string{
	"registerSuccess":     "註冊成功。歡迎加入我們的社群！",
	"bookingSuccess":      "課程預定成功。期待見到你！",
	"cancellationSuccess": "課程取消成功。希望在其他課程見到你。",
}

var EnUSEventMessage = map[string]string{
	"registerSuccess":     "Registration successful. Welcome to our community!",
	"bookingSuccess":      "Course successfully booked. We look forward to seeing you!",
	"cancellationSuccess": "Course successfully cancelled. We hope to see you in other courses.",
}

type ZnTW struct{}
type EnUS struct{}

func (z ZnTW) GetMessage(event_name string) string {
	return ZnTWEventMessage[event_name]
}

func (e EnUS) GetMessage(event_name string) string {
	return EnUSEventMessage[event_name]
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
