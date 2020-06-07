package facade

import "github.com/FuckAll/ddd-golang/domain/person/repository/po"

type PersonRepository interface {
	Insert(personPO *po.PersonPO) error

	Update(personPO *po.PersonPO) error

	FindById(id string) (*po.PersonPO, error)

	FindLeaderByPersonId(personId string) (*po.PersonPO, error)
}
