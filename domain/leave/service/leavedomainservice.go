package service

import (
	"github.com/FuckAll/ddd-golang/domain/leave/entity"
	"github.com/FuckAll/ddd-golang/domain/leave/entity/approver"
	leaveEvent "github.com/FuckAll/ddd-golang/domain/leave/event"
	"github.com/FuckAll/ddd-golang/domain/leave/repository/facade"
	"github.com/FuckAll/ddd-golang/infrastructure/common/event"
)

type LeaveDomainService struct {
	eventPublisher           event.EventPublisher
	leaveRepositoryInterface facade.ILeaveRepository
	leaveFactory             LeaveFactory
}

func (l *LeaveDomainService) CreateLeave(leave *entity.Leave, leaderMaxLevel int, approver approver.Approver) {
	leave.LeaderMaxLevel = leaderMaxLevel
	leave.Approver = approver
	leave.Create()
	l.leaveRepositoryInterface.Save(l.leaveFactory.CreateLeavePO(leave))
	event := leaveEvent.NewLeaveEvent(leaveEvent.CreateEvent, leave)
	l.leaveRepositoryInterface.SaveEvent(l.leaveFactory.CreateLeaveEventPO(event))
	l.eventPublisher.Publish(event)
}

func (l *LeaveDomainService) UpdateLeaveInfo(leave entity.Leave) error {
	leavePO, err := l.leaveRepositoryInterface.FindById(leave.Id)
	if err != nil {
		return err
	}
	return l.leaveRepositoryInterface.Save(leavePO)
}

func (l *LeaveDomainService) SubmitApproval(leave *entity.Leave, apr *approver.Approver) error {
	var event *leaveEvent.LeaveEvent
	if leave.CurrentApprovalInfo.ApprovalType == approver.Reject {
		leave.Reject(*apr)
		event = leaveEvent.NewLeaveEvent(leaveEvent.RejectEvent, leave)
	} else {
		if apr == nil {
			leave.Agree(*apr)
			event = leaveEvent.NewLeaveEvent(leaveEvent.ApprovedEvent, leave)
		}
	}
	leave.AddHistoryApprovalInfo(leave.CurrentApprovalInfo)
	l.leaveRepositoryInterface.Save(l.leaveFactory.CreateLeavePO(leave))
	l.leaveRepositoryInterface.SaveEvent(l.leaveFactory.CreateLeaveEventPO(event))
	return l.eventPublisher.Publish(event)
}

func (l *LeaveDomainService) GetLeaveInfo(leaveId string) (*entity.Leave, error) {
	leavePO, err := l.leaveRepositoryInterface.FindById(leaveId)
	if err != nil {
		return nil, err
	}
	return l.leaveFactory.GetLeave(leavePO)
}

func (l *LeaveDomainService) QueryLeaveInfosByApplicant(applicantId string) ([]*entity.Leave, error) {
	leavePOList, err := l.leaveRepositoryInterface.QueryByApplicantId(applicantId)
	if err != nil {
		return nil, err
	}
	var ret []*entity.Leave
	for _, leavePO := range leavePOList {
		leave, err := l.leaveFactory.GetLeave(leavePO)
		if err != nil {
			return nil, err
		}
		ret = append(ret, leave)
	}
	return ret, nil
}

func (l *LeaveDomainService) QueryLeaveInfosByApprover(approverId string) ([]*entity.Leave, error) {
	leavePOList, err := l.leaveRepositoryInterface.QueryByApproverId(approverId)
	if err != nil {
		return nil, err
	}
	var ret []*entity.Leave
	for _, leavePO := range leavePOList {
		leave, err := l.leaveFactory.GetLeave(leavePO)
		if err != nil {
			return nil, err
		}
		ret = append(ret, leave)
	}
	return ret, nil
}
