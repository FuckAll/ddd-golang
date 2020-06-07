package facade

import "github.com/FuckAll/ddd-golang/domain/rule/entity"

type ApprovalRuleRepository interface {
	GetLeaderMaxLevel(rule *entity.ApprovalRule) (int, error)
}
