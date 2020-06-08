package service

import (
	"github.com/FuckAll/ddd-golang/domain/rule/entity"
	"github.com/FuckAll/ddd-golang/domain/rule/repository/facade"
	"github.com/FuckAll/ddd-golang/domain/shared"
)

type ApprovalRuleDomainService struct {
	approvalRuleRepository facade.ApprovalRuleRepository
}

func NewApprovalRuleDomainService(approvalRuleRepository facade.ApprovalRuleRepository) *ApprovalRuleDomainService {
	return &ApprovalRuleDomainService{approvalRuleRepository: approvalRuleRepository}
}

func (a *ApprovalRuleDomainService) GetLeaderMaxLevel(personType shared.PersonType, levelType shared.LeaveType, duration int64) (int, error) {
	rule := &entity.ApprovalRule{
		PersonType: personType,
		LeaveType:  levelType,
		Duration:   duration,
	}
	return a.approvalRuleRepository.GetLeaderMaxLevel(rule)
}
