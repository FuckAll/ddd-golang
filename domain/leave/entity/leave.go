package entity

import (
	"github.com/FuckAll/ddd-golang/domain/leave/entity/applicant"
	"github.com/FuckAll/ddd-golang/domain/leave/entity/approver"
	"github.com/FuckAll/ddd-golang/domain/shared"
	"time"
)

/*
	聚合根
*/
type Leave struct {
	Id        string
	Applicant applicant.Applicant
	Approver  approver.Approver
	LeaveType shared.LeaveType
	Status    Status
	StartTime time.Time
	EndTime   time.Time
	Duration  int64
	//审批领导的最大级别
	LeaderMaxLevel       int
	CurrentApprovalInfo  ApprovalInfo
	HistoryApprovalInfos []ApprovalInfo
}

func (l *Leave) GetDuration() int64 {
	return int64(l.EndTime.Sub(l.StartTime))
}

func (l *Leave) AddHistoryApprovalInfo(approvalInfo ApprovalInfo) {
	l.HistoryApprovalInfos = append(l.HistoryApprovalInfos, approvalInfo)
}

func (l *Leave) Create() {
	l.Status = APPROVING
	l.StartTime = time.Now()
}

func (l *Leave) Agree(nextApprover approver.Approver) {
	l.Status = APPROVING
	l.Approver = nextApprover
}

func (l *Leave) Reject(approver approver.Approver) {
	l.Approver = approver
	l.Status = REJECTED
}

func (l *Leave) Finish() *Leave {
	l.Approver = approver.Approver{}
	l.Status = APPROVED
	l.EndTime = time.Now()
	l.Duration = l.GetDuration()
	return l
}
