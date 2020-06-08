package facade

import (
	"github.com/FuckAll/ddd-golang/application/service"
	"github.com/FuckAll/ddd-golang/domain/person/entity"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthApi struct {
	engine                  *gin.Engine
	loginApplicationService *service.LoginApplicationService
}

func NewAuthApi(engine *gin.Engine, loginApplicationService *service.LoginApplicationService) *AuthApi {
	return &AuthApi{engine: engine, loginApplicationService: loginApplicationService}
}

func (a *AuthApi) InitApi() {
	a.engine.GET("/login", func(c *gin.Context) {
		//
		if c.Query("password") != "1024" {
			c.JSON(403, gin.H{
				"message": "login fail",
			})
			return
		}
		person := entity.Person{
			PersonId:       "007",
			PersonName:     "wu",
			PersonType:     0,
			Relationships:  nil,
			RoleLevel:      0,
			CreateTime:     time.Time{},
			LastModifyTime: time.Time{},
			Status:         0,
		}
		login, err := a.loginApplicationService.Login(person)
		if err != nil {
			c.String(200, "err :%+v", err)
			return
		}

		c.JSON(200, gin.H{
			"message": login,
		})
	})
}
