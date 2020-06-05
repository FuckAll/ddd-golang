package service

import (
	"github.com/FuckAll/ddd-golang/leave/entity"
	"github.com/FuckAll/ddd-golang/leave/event"
	"github.com/FuckAll/ddd-golang/leave/repository/po"
)

type LeaveFactory struct {
}

func (l *LeaveFactory) CreateLeavePO(leave entity.Leave) po.LeavePO {
	return po.LeavePO{}
}

func (l *LeaveFactory) GetLeave(leavePO po.LeavePO) entity.Leave {
	return entity.Leave{}
}

func (l *LeaveFactory) CreateLeaveEventPO(leaveEvent event.LeaveEvent) po.LeaveEventPO {
	return po.LeaveEventPO{}
}

func (l *LeaveFactory) ApprovalInfoPOListFromDO(leave entity.Leave) []po.ApprovalInfoPO {
	return []po.ApprovalInfoPO{}
}
