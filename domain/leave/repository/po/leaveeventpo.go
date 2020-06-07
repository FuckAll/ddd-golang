package po

import (
	"github.com/FuckAll/ddd-golang/domain/leave/event"
	"time"
)

type LeaveEventPO struct {
	id             int
	leaveEventType event.LeaveEventType
	timestamp      time.Time
	source         string
	data           string
}
