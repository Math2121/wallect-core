package events_test

import (
	"sync"
	"testing"
	"time"

	"github.com/Math2121/walletcore/pkg/eventos/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}
func (e *TestEvent) SetPayload(payload interface{}){
	e.Payload = payload
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {

}

type EventDispatcherSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	eventDispatcher *events.EventDispatcher
}

func (suite *EventDispatcherSuite) SetupTest() {
	suite.event = TestEvent{Name: "TestEvent", Payload: "TestPayload"}
	suite.event2 = TestEvent{Name: "TestEvent2", Payload: "TestPayload2"}
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.eventDispatcher = events.NewEventDispatcher()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherSuite))
}

func (suite *EventDispatcherSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)

	suite.Equal(2, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))

}

func (suite *EventDispatcherSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(events.ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherSuite) TestEventDispatcher_Clear() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)

	suite.Equal(2, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.Handlers))

}

func (suite *EventDispatcherSuite) TestEventDispatcher_Has() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)

	suite.Equal(2, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler))
}

func (suite *EventDispatcherSuite) TestEventDispatcher_Remove(){
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.event.GetName()]))



}
type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event events.EventInterface,  wg *sync.WaitGroup) {
	m.Called(event)
}

func (suite *EventDispatcherSuite) TestEventDispatcher_Dispatch() {
	handler := new(MockHandler)
	handler.On("Handle", &suite.event).Return()

	err := suite.eventDispatcher.Register(suite.event.GetName(), handler)
	suite.Nil(err)

	err = suite.eventDispatcher.Dispatch(&suite.event)
	suite.Nil(err)

	handler.AssertCalled(suite.T(), "Handle", &suite.event)
	handler.AssertNumberOfCalls(suite.T(), "Handle", 1)
}
