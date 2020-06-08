package service

import (
	"fmt"
	"github.com/FuckAll/ddd-golang/domain/person/entity"
	"github.com/FuckAll/ddd-golang/domain/person/repository/facade"
	"github.com/FuckAll/ddd-golang/domain/person/repository/po"
)

type PersonFactory struct {
	// 要通过依赖注入
	personRepository facade.PersonRepository
}

func NewPersonFactory(personRepository facade.PersonRepository) *PersonFactory {
	return &PersonFactory{personRepository: personRepository}
}

func (p *PersonFactory) getPerson(personPO *po.PersonPO) (*entity.Person, error) {
	person, err := p.personRepository.FindById(personPO.PersonId)
	if err != nil {
		return nil, err
	}
	if person == nil {
		return nil, fmt.Errorf("person_id %d does not exist", personPO.PersonId)
	}
	return CreatePerson(person), nil
}

func CreatePersonPO(person *entity.Person) *po.PersonPO {
	return &po.PersonPO{
		PersonId:       person.PersonId,
		PersonName:     person.PersonName,
		PersonType:     person.PersonType,
		RoleLevel:      person.RoleLevel,
		CreateTime:     person.CreateTime,
		LastModifyTime: person.LastModifyTime,
	}
}

func CreatePerson(personPO *po.PersonPO) *entity.Person {
	return &entity.Person{
		PersonId:       personPO.PersonId,
		PersonName:     personPO.PersonName,
		PersonType:     personPO.PersonType,
		Relationships:  nil,
		RoleLevel:      personPO.RoleLevel,
		CreateTime:     personPO.CreateTime,
		LastModifyTime: personPO.LastModifyTime,
		Status:         personPO.Status,
	}
}
