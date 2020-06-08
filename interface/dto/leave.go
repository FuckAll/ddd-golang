package dto

type LeaveDTO struct {
	LeaveId                    string
	ApplicantDTO               ApplicantDTO
	ApproverDTO                ApproverDTO
	LeaveType                  int
	CurrentApprovalInfoDTO     ApprovalInfoDTO
	HistoryApprovalInfoDTOList []ApprovalInfoDTO
	StartTime                  string
	EndTime                    string
	Duration                   int64
	Status                     int
}
