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

func (l *LeaveRepositoryImpl) Save(LeavePO *po.LeavePO) error {
	return nil

}

func (l *LeaveRepositoryImpl) SaveEvent(leaveEventPO *po.LeaveEventPO) error {
	return nil
}

func (l *LeaveRepositoryImpl) FindById(id string) (*po.LeavePO, error) {
	return nil, nil
}

func (l *LeaveRepositoryImpl) QueryByApplicantId(applicantId string) ([]*po.LeavePO, error) {
	return nil, nil
}

func (l *LeaveRepositoryImpl) QueryByApproverId(approverId string) ([]*po.LeavePO, error) {
	return nil, nil
}

func NewLeaveRepositoryImpl() *LeaveRepositoryImpl {
	return &LeaveRepositoryImpl{}
}
