package event

import "github.com/FuckAll/ddd-golang/domain/leave/event"

type EventPublisher struct {
}

func (e *EventPublisher) Publish(event *event.LeaveEvent) error {
	// send

	return nil
}
