package service

import (
	leaveEntity "github.com/FuckAll/ddd-golang/domain/leave/entity"
	approver "github.com/FuckAll/ddd-golang/domain/leave/entity/approver"
	leaveService "github.com/FuckAll/ddd-golang/domain/leave/service"
	personService "github.com/FuckAll/ddd-golang/domain/person/service"
	ruleService "github.com/FuckAll/ddd-golang/domain/rule/service"
)

type LeaveApplicationService struct {
	leaveDomainService        leaveService.LeaveDomainService
	personDomainService       personService.PersonDomainService
	approvalRuleDomainService ruleService.ApprovalRuleDomainService
}

// CreateLeaveInfo 创建请假申请，并为审批人生成人物
func (l *LeaveApplicationService) CreateLeaveInfo(leave *leaveEntity.Leave) error {
	// get approval leader max level by rule
	leaderMaxLevel, err := l.approvalRuleDomainService.GetLeaderMaxLevel(leave.Applicant.PersonType, leave.LeaveType,
		leave.Duration)
	if err != nil {
		return err
	}

	approverPerson, err := l.personDomainService.FindFirstApprover(leave.Applicant.PersonId, leaderMaxLevel)
	if err != nil {
		return err
	}

	l.leaveDomainService.CreateLeave(leave, leaderMaxLevel, *approver.NewApprover(approverPerson.PersonId,
		approverPerson.PersonName, approverPerson.RoleLevel))
	return nil
}

// UpdateLeaveInfo 更新请假单基本信息
func (l *LeaveApplicationService) UpdateLeaveInfo(leave *leaveEntity.Leave) error {
	return l.leaveDomainService.UpdateLeaveInfo(*leave)
}

// SubmitApproval 提交审批，更新请假单信息
func (l *LeaveApplicationService) SubmitApproval(leave *leaveEntity.Leave) error {
	nextApprover, err := l.personDomainService.FindNextApprover(leave.Approver.PersonId, leave.LeaderMaxLevel)
	if err != nil {
		return err
	}
	return l.leaveDomainService.SubmitApproval(leave, approver.NewApprover(nextApprover.PersonId, nextApprover.PersonName, nextApprover.RoleLevel))
}

// GetLeaveInfo 获取请假单
func (l *LeaveApplicationService) GetLeaveInfo(leaveId string) (*leaveEntity.Leave, error) {
	return l.leaveDomainService.GetLeaveInfo(leaveId)
}

// QueryLeaveInfosByApplicant 通过申请人获取请假单
func (l *LeaveApplicationService) QueryLeaveInfosByApplicant(applicantId string) ([]*leaveEntity.Leave, error) {
	return l.leaveDomainService.QueryLeaveInfosByApplicant(applicantId)
}

// QueryLeaveInfosByApprover 通过申请人获取请假单
func (l *LeaveApplicationService) QueryLeaveInfosByApprover(approverId string) ([]*leaveEntity.Leave, error) {
	return l.leaveDomainService.QueryLeaveInfosByApprover(approverId)
}
