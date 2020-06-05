package facade

import (
	"github.com/FuckAll/ddd-golang/leave/repository/po"
)

type ILeaveRepository interface {
	Save(LeavePO po.LeavePO)
	//
	SaveEvent(leaveEventPO po.LeaveEventPO)
	//
	findById(id string) po.LeavePO
	//
	queryByApplicantId(applicantId string) []po.LeavePO
	//
	queryByApproverId(approverId string) []po.LeavePO
}
