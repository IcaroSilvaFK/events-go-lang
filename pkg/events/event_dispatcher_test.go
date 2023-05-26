package events_test

import (
	"testing"
	"time"

	"github.com/IcaroSilvaFK/events-go-lang/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

type TestEventHandler struct {
}

type EventDispatcherTestSuite struct {
	suite.Suite

	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *events.EventDispatcher
}

func (e *TestEvent) GetName() string {
	return e.Name
}
func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (h *TestEventHandler) Handle(event events.EventInterface) error {
	return nil
}

func TestShouldCreateEvent(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (s *EventDispatcherTestSuite) SetupTest() {

	s.event = TestEvent{
		Name:    "event1",
		Payload: "test",
	}
	s.event2 = TestEvent{
		Name:    "event2",
		Payload: "test",
	}

	s.dispatcher = events.NewEventDispatcher()
	s.handler = TestEventHandler{}
	s.handler2 = TestEventHandler{}
	s.handler3 = TestEventHandler{}

}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {

	assert.True(s.T(), true)

}
