package event

import "github.com/FuckAll/ddd-golang/leave/event"

type EventPublisher struct {
}

func (e *EventPublisher) Publish(event event.LeaveEvent) {
	// send

}
