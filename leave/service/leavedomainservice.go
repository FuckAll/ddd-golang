package service

import (
	"github.com/FuckAll/ddd-golang/infrastructure/common/event"
	"github.com/FuckAll/ddd-golang/leave/entity"
	"github.com/FuckAll/ddd-golang/leave/entity/approver"
	event2 "github.com/FuckAll/ddd-golang/leave/event"
	"github.com/FuckAll/ddd-golang/leave/repository/facade"
)

type LeaveDomainService struct {
	eventPublisher           event.EventPublisher
	leaveRepositoryInterface facade.ILeaveRepository
	leaveFactory             LeaveFactory
}

func (l *LeaveDomainService) CreateLeave(leave entity.Leave, leaderMaxLevel int, approver approver.Approver) {
	leave.LeaderMaxLevel = leaderMaxLevel
	leave.Approver = approver
	leave = *leave.Create()
	l.leaveRepositoryInterface.Save(l.leaveFactory.CreateLeavePO(leave))
	leaveEvent := event2.NewLeaveEvent(event2.CreateEvent, leave)
	l.leaveRepositoryInterface.SaveEvent(l.leaveFactory.CreateLeaveEventPO(leaveEvent))
	l.eventPublisher.Publish(leaveEvent)
}