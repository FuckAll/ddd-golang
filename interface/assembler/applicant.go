package assembler

import (
	"github.com/FuckAll/ddd-golang/domain/leave/entity/applicant"
	"github.com/FuckAll/ddd-golang/interface/dto"
)

//  ToApplicantDTO DTO(Data Transfer Object) --> DO (Domain Object)
func ToApplicantDTO(applicant applicant.Applicant) dto.ApplicantDTO {
	return dto.ApplicantDTO{
		PersonId:   applicant.PersonId,
		PersonName: applicant.PersonName,
	}
}

func ToApplicantDO(applicantDTO dto.ApplicantDTO) applicant.Applicant {
	return applicant.Applicant{
		PersonId:   applicantDTO.PersonId,
		PersonName: applicantDTO.PersonName,
	}
}
