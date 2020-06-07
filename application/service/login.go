package service

import "github.com/FuckAll/ddd-golang/domain/person/entity"

type LoginApplicationService struct {
	// 登录服务
}

func (l *LoginApplicationService) Login(person entity.Person) (string, error) {
	return "ok", nil
}
