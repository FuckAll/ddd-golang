package entity

import "github.com/FuckAll/ddd-golang/leave/entity/approver"

type ApprovalInfo struct {
	approvalInfoId string
	approver       approver.Approver
	approvalType   approver.ApprovalType
	msg            string
	time           float64
}
