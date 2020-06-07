package facade

import (
	"github.com/FuckAll/ddd-golang/domain/leave/repository/po"
)

type ILeaveRepository interface {
	Save(LeavePO *po.LeavePO) error
	//
	SaveEvent(leaveEventPO *po.LeaveEventPO) error
	//
	FindById(id string) (*po.LeavePO, error)
	//
	QueryByApplicantId(applicantId string) ([]*po.LeavePO, error)
	//
	QueryByApproverId(approverId string) ([]*po.LeavePO, error)
}
