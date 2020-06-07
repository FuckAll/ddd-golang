package entity

import "github.com/FuckAll/ddd-golang/domain/shared"

type ApprovalRule struct {
	PersonType     shared.PersonType
	LeaveType      shared.LeaveType
	Duration       int64
	MaxLeaderLevel int
}

func GetByLeave(personType shared.PersonType, leaveType shared.LeaveType, duration int64) *ApprovalRule {
	return &ApprovalRule{
		PersonType: personType,
		LeaveType:  leaveType,
		Duration:   duration,
	}
}
