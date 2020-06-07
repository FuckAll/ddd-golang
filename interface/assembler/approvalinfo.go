package assembler

import (
	"github.com/FuckAll/ddd-golang/domain/leave/entity"
	"github.com/FuckAll/ddd-golang/interface/dto"
)

func ToApprovalInfoDTO(info entity.ApprovalInfo) dto.ApprovalInfoDTO {
	return dto.ApprovalInfoDTO{
		ApprovalInfoId: info.ApprovalInfoId,
		Msg:            info.Msg,
		Time:           info.Time,
	}
}

func ToApprovalInfoDO(infoDTO dto.ApprovalInfoDTO) entity.ApprovalInfo {
	return entity.ApprovalInfo{
		ApprovalInfoId: infoDTO.ApprovalInfoId,
		Msg:            infoDTO.Msg,
		Time:           infoDTO.Time,
	}
}
