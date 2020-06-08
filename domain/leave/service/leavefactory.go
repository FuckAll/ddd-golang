package service

import (
	"github.com/FuckAll/ddd-golang/domain/leave/entity"
	"github.com/FuckAll/ddd-golang/domain/leave/event"
	"github.com/FuckAll/ddd-golang/domain/leave/repository/po"
)

type LeaveFactory struct {
}

func NewLeaveFactory() *LeaveFactory {
	return &LeaveFactory{}
}

func (l *LeaveFactory) CreateLeavePO(leave *entity.Leave) *po.LeavePO {
	return &po.LeavePO{}
}

func (l *LeaveFactory) GetLeave(leavePO *po.LeavePO) (*entity.Leave, error) {
	return &entity.Leave{}, nil
}

func (l *LeaveFactory) CreateLeaveEventPO(leaveEvent *event.LeaveEvent) *po.LeaveEventPO {
	return &po.LeaveEventPO{}
}

func (l *LeaveFactory) ApprovalInfoPOListFromDO(leave *entity.Leave) []*po.ApprovalInfoPO {
	return []*po.ApprovalInfoPO{}
}
