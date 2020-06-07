package dto

type LeaveDTO struct {
	LeaveId                    string
	applicantDTO               ApplicantDTO
	approverDTO                ApproverDTO
	leaveType                  string
	currentApprovalInfoDTO     ApprovalInfoDTO
	historyApprovalInfoDTOList []ApprovalInfoDTO
	startTime                  string
	endTime                    string
	duration                   int64
	status                     string
}
