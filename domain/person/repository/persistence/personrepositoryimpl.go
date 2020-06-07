package persistence

import (
	"github.com/FuckAll/ddd-golang/domain/person/repository/po"
)

type PersonRepositoryImpl struct {
	// 初始化数据库
}

func (p PersonRepositoryImpl) Insert(personPO *po.PersonPO) error {
	panic("implement me")
}

func (p PersonRepositoryImpl) Update(personPO *po.PersonPO) error {
	panic("implement me")
}

func (p PersonRepositoryImpl) FindById(id string) (*po.PersonPO, error) {
	panic("implement me")
}

func (p PersonRepositoryImpl) FindLeaderByPersonId(personId string) (*po.PersonPO, error) {
	panic("implement me")
}
