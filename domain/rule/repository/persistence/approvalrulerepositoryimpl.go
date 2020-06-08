package persistence

import "github.com/FuckAll/ddd-golang/domain/rule/entity"

type ApprovalRuleRepositoryImpl struct {
	//

}

func (a *ApprovalRuleRepositoryImpl) GetLeaderMaxLevel(rule *entity.ApprovalRule) (int, error) {
	return 0, nil
}

func NewApprovalRuleRepositoryImpl() *ApprovalRuleRepositoryImpl {
	return &ApprovalRuleRepositoryImpl{}

}
