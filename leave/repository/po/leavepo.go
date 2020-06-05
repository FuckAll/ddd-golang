package po

import (
	"github.com/FuckAll/ddd-golang/leave/entity"
	"time"
)

type LeavePO struct {
	Id            string
	ApplicantId   string
	ApplicantName string
	//applicantType PersonType;
	ApplicantType             int
	ApproverId                string
	ApproverName              string
	LeaveType                 entity.LeaveType
	Status                    entity.Status
	StartTime                 time.Time
	EndTime                   time.Time
	Duration                  int64
	HistoryApprovalInfoPOList []ApprovalInfoPO
}
