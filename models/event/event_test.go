package event_test

import (
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

// 那mock一個接口，就是實現這個接口的所有方法，這裡就是Notify和GetName
func (m *MockNotifier) Notify(user user.User, message string) {
	m.Called(user, message)
}

func (m *MockNotifier) GetName() string {
	args := m.Called()
	return args.String(0)
}

func TestEvent(t *testing.T) {
	eventFactory := event.NewEventFactory()

	// 創建三個MockNotifier
	mockLineNotifier := new(MockNotifier)
	mockTelegramNotifier := new(MockNotifier)
	mockSMSNotifier := new(MockNotifier)

	// 為每個MockNotifier設置GetName方法的預期行為
	mockLineNotifier.On("GetName").Return("Line")
	mockTelegramNotifier.On("GetName").Return("Telegram")
	mockSMSNotifier.On("GetName").Return("SMS")

	// 設置每個MockNotifier的Notify方法的預期行為
	testUser := user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}
	expectedMessage := "expected message"
	mockLineNotifier.On("Notify", testUser, expectedMessage).Return().Once()
	mockTelegramNotifier.On("Notify", testUser, expectedMessage).Return().Once()
	mockSMSNotifier.On("Notify", testUser, expectedMessage).Return().Once()

	// 將MockNotifier添加到EventFactory的Methods中
	eventFactory.AddNotifier(mockLineNotifier)
	eventFactory.AddNotifier(mockTelegramNotifier)
	eventFactory.AddNotifier(mockSMSNotifier)

	// 觸發事件
	eventFactory.Trigger(testUser, expectedMessage)

	// 驗證每個MockNotifier的Notify方法是否都各被调用了一次
	mockLineNotifier.AssertNumberOfCalls(t, "Notify", 1)
	mockTelegramNotifier.AssertNumberOfCalls(t, "Notify", 1)
	mockSMSNotifier.AssertNumberOfCalls(t, "Notify", 1)
}

func TestEventFactoryEarlyReturn(t *testing.T) {
	user := user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}
	errorEventName := ""

	event.TriggerHelper(user, "")

	assert.Empty(t, event.TriggerHelper(user, errorEventName))
}
