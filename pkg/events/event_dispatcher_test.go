package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

type TestEventHandler struct {
	ID int
}

type EventDispatcherTestSuite struct {
	suite.Suite

	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *EventDispatcher
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

func (h *TestEventHandler) Handle(event EventInterface) error {
	return nil
}

func TestSuite(t *testing.T) {
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

	s.dispatcher = NewEventDispatcher()
	s.handler = TestEventHandler{
		ID: 1,
	}
	s.handler2 = TestEventHandler{
		ID: 2,
	}
	s.handler3 = TestEventHandler{
		ID: 3,
	}

}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {

	err := s.dispatcher.Register(s.event.GetName(), &s.handler)

	assert.Nil(s.T(), err)

	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)

	assert.Nil(s.T(), err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))

	assert.Equal(s.T(), &s.handler, s.dispatcher.handlers[s.event.GetName()][0])
	assert.Equal(s.T(), &s.handler2, s.dispatcher.handlers[s.event.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {

	s.dispatcher.Register(s.event.GetName(), &s.handler)

	err := s.dispatcher.Register(s.event.GetName(), &s.handler)

	s.Error(err)
	s.Equal(ErrorHandlerAlreadyRegistered, err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {

	err := s.dispatcher.Register(s.event.GetName(), &s.handler)

	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Register(s.event2.GetName(), &s.handler2)

	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Register(s.event2.GetName(), &s.handler3)
	s.Nil(err)

	s.dispatcher.Clear()

	s.Equal(0, len(s.dispatcher.handlers))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {

	err := s.dispatcher.Register(s.event.GetName(), &s.handler)

	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)

	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))

	s.True(s.dispatcher.Has(s.event.GetName(), &s.handler))
	s.True(s.dispatcher.Has(s.event.GetName(), &s.handler2))
}
