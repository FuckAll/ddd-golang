package event

type LeaveEventType int

const (
	CreateEvent   LeaveEventType = 1
	AgreeEvent    LeaveEventType = 2
	RejectEvent   LeaveEventType = 3
	ApprovedEvent LeaveEventType = 4
)
