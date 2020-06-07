package service

import (
	"github.com/FuckAll/ddd-golang/domain/person/entity"
	personService "github.com/FuckAll/ddd-golang/domain/person/service"
)

type PersonApplicationService struct {
	personDomainService personService.PersonDomainService
}

func (p *PersonApplicationService) Create(person *entity.Person) error {
	return p.personDomainService.Create(person)
}

func (p *PersonApplicationService) Update(person *entity.Person) error {
	return p.personDomainService.Update(person)
}

func (p *PersonApplicationService) DeleteById(personId string) error {
	return p.personDomainService.DeleteById(personId)
}

func (p *PersonApplicationService) FindById(personId string) error {
	return p.FindById(personId)
}

func (p *PersonApplicationService) FindFirstApprover(applicantId string, leaderMaxLevel int) (*entity.Person, error) {
	return p.personDomainService.FindFirstApprover(applicantId, leaderMaxLevel)
}
