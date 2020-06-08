package assembler

import (
	"github.com/FuckAll/ddd-golang/domain/person/entity"
	"github.com/FuckAll/ddd-golang/domain/shared"
	"github.com/FuckAll/ddd-golang/interface/dto"
	"time"
)

func ToPersonDTO(person entity.Person) dto.PersonDTO {
	personDTO := dto.PersonDTO{}
	personDTO.PersonId = person.PersonId
	personDTO.PersonName = person.PersonName
	personDTO.PersonType = int(person.PersonType)
	personDTO.Status = int(person.Status)
	personDTO.CreateTime = person.CreateTime.String()
	personDTO.LastModifyTime = person.LastModifyTime.String()
	return personDTO
}

func ToPersonDO(personDTO dto.PersonDTO) entity.Person {
	createTime, _ := time.Parse(time.RFC3339, personDTO.CreateTime)
	lastModifyTime, _ := time.Parse(time.RFC3339, personDTO.LastModifyTime)
	person := entity.Person{
		PersonId:       personDTO.PersonId,
		PersonName:     personDTO.PersonName,
		PersonType:     shared.PersonType(personDTO.PersonType),
		Relationships:  nil,
		RoleLevel:      0,
		CreateTime:     createTime,
		LastModifyTime: lastModifyTime,
		Status:         entity.PersonStatus(personDTO.Status),
	}
	return person
}
