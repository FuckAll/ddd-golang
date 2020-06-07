package po

import "github.com/FuckAll/ddd-golang/domain/shared"

type ApprovalRulePO struct {
	Id              string
	LeaveType       shared.LeaveType
	PersonType      shared.PersonType
	Duration        int
	ApplicantRoleId string
}
