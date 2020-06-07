package entity

import "github.com/FuckAll/ddd-golang/domain/leave/entity/approver"

type ApprovalInfo struct {
	ApprovalInfoId string
	Approver       approver.Approver
	ApprovalType   approver.ApprovalType
	Msg            string
	Time           int64
}
