package po

import (
	"github.com/FuckAll/ddd-golang/domain/leave/entity"
	"github.com/FuckAll/ddd-golang/domain/shared"

	"time"
)

type LeavePO struct {
	Id                        string
	ApplicantId               string
	ApplicantName             string
	applicantType             shared.PersonType
	ApplicantType             int
	ApproverId                string
	ApproverName              string
	LeaveType                 shared.LeaveType
	Status                    entity.Status
	StartTime                 time.Time
	EndTime                   time.Time
	Duration                  int64
	HistoryApprovalInfoPOList []ApprovalInfoPO
}
