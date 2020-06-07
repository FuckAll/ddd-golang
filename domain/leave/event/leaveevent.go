package event

import (
	"encoding/json"
	"github.com/FuckAll/ddd-golang/domain/leave/entity"
	"time"
)

type LeaveEvent struct {
	leaveEventType LeaveEventType
	id             string
	timestamp      time.Time
	source         string
	data           string
}

func NewLeaveEvent(leaveEventType LeaveEventType, leave *entity.Leave) *LeaveEvent {
	bytes, _ := json.Marshal(leave)
	// id 是生成的
	return &LeaveEvent{leaveEventType: leaveEventType, id: "1", timestamp: time.Now(), data: string(bytes)}
}
