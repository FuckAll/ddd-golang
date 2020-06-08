package service

import "github.com/FuckAll/ddd-golang/domain/person/entity"

type LoginApplicationService struct {
	// 登录服务
}

func NewLoginApplicationService() *LoginApplicationService {
	return &LoginApplicationService{}
}
func (l *LoginApplicationService) Login(person entity.Person) (string, error) {
	return "ok", nil
}
