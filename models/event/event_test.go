package event_test

import (
	"be101_golang/models/constant"
	"be101_golang/models/event"
	"be101_golang/models/language"
	"be101_golang/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockNotifier is a mock implementation of the Notifier interface
// 意思就是mock一个Notifier接口
type MockNotifier struct {
	mock.Mock
}

// 試試看加上Table-Driven Tests 的精神，以及委託策略
type NotifierTest struct {
	Notifier  *MockNotifier
	Name      string
	CallTimes int
}

// 那mock一個接口，就是實現這個接口的所有方法，這裡就是Notify和GetName
func (m *MockNotifier) Notify(user user.User, message string) {
	m.Called(user, message)
}

func (m *MockNotifier) GetName() string {
	args := m.Called() // args是一個mock.Call，這個Call是一個struct，裡面有一個Arguments的interface{}，這個interface{}是一個"slice"(這就是為啥有[0])，裡面裝的是我們傳進去的參數，也就是Notifier.Name
	return args.String(0)
}

// ? 確保在Trigger被呼叫時，每個Notifier的Notify方法都被按預期調用了正確的次數
func TestEvent(t *testing.T) {
	eventFactory := event.NewEventFactory()

	Notifiers := []NotifierTest{
		{Notifier: new(MockNotifier), Name: constant.Line, CallTimes: 1},
		{Notifier: new(MockNotifier), Name: constant.Telegram, CallTimes: 1},
		{Notifier: new(MockNotifier), Name: constant.SMS, CallTimes: 1},
		{Notifier: new(MockNotifier), Name: constant.Email, CallTimes: 1},
	}

	// 這邊就是把所有的Notifier都加進去EventFactory的Methods中
	for _, notifier := range Notifiers {
		notifier.Notifier.On("GetName").Return(notifier.Name)
		notifier.Notifier.On("Notify", mock.Anything, mock.Anything).Return().Times(notifier.CallTimes)
		eventFactory.AddNotifier(notifier.Notifier)
	}

	// 觸發事件
	testObject := user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}
	eventFactory.Trigger(testObject, "expected message")

	// 驗證每個MockNotifier的Notify方法是否都各被调用了一次
	for _, notifier := range Notifiers {
		notifier.Notifier.AssertNumberOfCalls(t, "Notify", notifier.CallTimes)
	}
}

func TestEventFactoryEarlyReturn(t *testing.T) {
	user := user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}
	errorEventName := ""

	event.TriggerHelper(user, "")

	assert.Empty(t, event.TriggerHelper(user, errorEventName))
}
