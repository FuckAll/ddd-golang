package assembler

import (
	"github.com/FuckAll/ddd-golang/domain/leave/entity/approver"
	"github.com/FuckAll/ddd-golang/interface/dto"
)

func ToApproverDTO(approver approver.Approver) dto.ApproverDTO {
	return dto.ApproverDTO{
		PersonId:   approver.PersonId,
		PersonName: approver.PersonName,
	}
}

func ToApproverDO(approverDTO dto.ApproverDTO) approver.Approver {
	return approver.Approver{
		PersonId:   approverDTO.PersonId,
		PersonName: approverDTO.PersonName,
	}
}
