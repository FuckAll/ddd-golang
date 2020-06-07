package service

import (
	"fmt"
	"github.com/FuckAll/ddd-golang/domain/person/entity"
	"github.com/FuckAll/ddd-golang/domain/person/repository/facade"
	"time"
)

type PersonDomainService struct {
	// 通过依赖注入
	personRepository facade.PersonRepository

	// 依赖注入
	personFactory PersonFactory
}

func (p *PersonDomainService) Create(person *entity.Person) error {
	personPO, err := p.personRepository.FindById(person.PersonId)
	if err != nil {
		return err
	}
	if personPO == nil {
		return fmt.Errorf("person_id %d does not exist", person.PersonId)
	}
	person.Create()
	return p.personRepository.Insert(personPO)
}

func (p *PersonDomainService) Update(person *entity.Person) error {
	person.LastModifyTime = time.Now()
	return p.personRepository.Update(CreatePersonPO(person))
}

func (p *PersonDomainService) DeleteById(personId string) error {
	personPO, err := p.personRepository.FindById(personId)
	if err != nil {
		return err
	}
	person, err := p.personFactory.getPerson(personPO)
	if err != nil {
		return err
	}
	person.Disable()
	return p.personRepository.Update(CreatePersonPO(person))
}

func (p *PersonDomainService) FindById(userId string) (*entity.Person, error) {
	personPO, err := p.personRepository.FindById(userId)
	if err != nil {
		return nil, err
	}
	return p.personFactory.getPerson(personPO)
}

func (p *PersonDomainService) FindFirstApprover(applicantId string, leaderMaxLevel int) (*entity.Person, error) {
	leaderPO, err := p.personRepository.FindLeaderByPersonId(applicantId)
	if err != nil {
		return nil, err
	}
	if leaderPO.RoleLevel > leaderMaxLevel || leaderPO == nil {
		return nil, nil
	} else {
		return CreatePerson(leaderPO), nil
	}
}

func (p *PersonDomainService) FindNextApprover(currentApproverId string, leaderMaxLevel int) (*entity.Person, error) {
	leaderPO, err := p.personRepository.FindLeaderByPersonId(currentApproverId)
	if err != nil {
		return nil, err
	}
	if leaderPO == nil {
		return nil, nil
	}
	return CreatePerson(leaderPO), nil
}
