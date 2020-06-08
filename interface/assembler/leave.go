package assembler

import (
	leaveEntity "github.com/FuckAll/ddd-golang/domain/leave/entity"
	"github.com/FuckAll/ddd-golang/interface/dto"
)

func ToLeaveDTO(leave leaveEntity.Leave) dto.LeaveDTO {
	leaveDTO := dto.LeaveDTO{}
	leaveDTO.LeaveId = leave.Id
	leaveDTO.LeaveType = int(leave.LeaveType)
	leaveDTO.Status = int(leave.Status)
	leaveDTO.StartTime = leave.StartTime.String()
	leaveDTO.EndTime = leave.EndTime.String()
	leaveDTO.CurrentApprovalInfoDTO = ToApprovalInfoDTO(leave.CurrentApprovalInfo)
	var historyApprovalInfoDTOList []dto.ApprovalInfoDTO
	for _, x := range leave.HistoryApprovalInfos {
		historyApprovalInfoDTOList = append(historyApprovalInfoDTOList, ToApprovalInfoDTO(x))
	}
	leaveDTO.HistoryApprovalInfoDTOList = historyApprovalInfoDTOList
	leaveDTO.Duration = leave.Duration
	return leaveDTO
}

func ToLeaveDO(leaveDTO dto.LeaveDTO) leaveEntity.Leave {
	leave := leaveEntity.Leave{}
	leave.Id = leaveDTO.LeaveId
	leave.Applicant = ToApplicantDO(leaveDTO.ApplicantDTO)
	leave.Approver = ToApproverDO(leaveDTO.ApproverDTO)
	leave.CurrentApprovalInfo = ToApprovalInfoDO(leaveDTO.CurrentApprovalInfoDTO)
	var historyApprovalInfos []leaveEntity.ApprovalInfo
	for _, x := range leaveDTO.HistoryApprovalInfoDTOList {
		historyApprovalInfos = append(historyApprovalInfos, ToApprovalInfoDO(x))
	}
	leave.HistoryApprovalInfos = historyApprovalInfos
	return leave
}
