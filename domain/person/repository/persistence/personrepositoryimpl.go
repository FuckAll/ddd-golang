package persistence

import (
	"github.com/FuckAll/ddd-golang/domain/person/repository/po"
)

type PersonRepositoryImpl struct {
	// 初始化数据库
}

func NewPersonRepositoryImpl() *PersonRepositoryImpl {
	return &PersonRepositoryImpl{}
}

func (p *PersonRepositoryImpl) Insert(personPO *po.PersonPO) error {
	return nil
}

func (p *PersonRepositoryImpl) Update(personPO *po.PersonPO) error {
	return nil
}

func (p *PersonRepositoryImpl) FindById(id string) (*po.PersonPO, error) {
	return nil, nil
}

func (p *PersonRepositoryImpl) FindLeaderByPersonId(personId string) (*po.PersonPO, error) {
	return nil, nil
}
