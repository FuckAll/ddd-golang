package persistence

import (
	"github.com/FuckAll/ddd-golang/domain/leave/repository/po"
)

type LeaveRepositoryImpl struct {
	// 初始化一些资源
	//@Autowired
	//LeaveDao leaveDao;
	//@Autowired
	//ApprovalInfoDao approvalInfoDao;
	//@Autowired
	//LeaveEventDao leaveEventDao;
}

func (l *LeaveRepositoryImpl) Save(LeavePO po.LeavePO) {
	panic("implement me")
}

func (l *LeaveRepositoryImpl) SaveEvent(leaveEventPO po.LeaveEventPO) {
	panic("implement me")
}

func (l *LeaveRepositoryImpl) findById(id string) po.LeavePO {
	panic("implement me")
}

func (l *LeaveRepositoryImpl) queryByApplicantId(applicantId string) []po.LeavePO {
	panic("implement me")
}

func (l *LeaveRepositoryImpl) queryByApproverId(approverId string) []po.LeavePO {
	panic("implement me")
}
