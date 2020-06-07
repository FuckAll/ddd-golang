package dto

type ApprovalInfoDTO struct {
	ApprovalInfoId string
	ApprovalDTO    ApproverDTO
	Msg            string
	Time           int64
}
