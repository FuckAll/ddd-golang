package facade

import (
	"github.com/FuckAll/ddd-golang/application/service"
	"github.com/gin-gonic/gin"
)

type LeaveAPI struct {
	engine                  *gin.Engine
	leaveApplicationService *service.LeaveApplicationService
}

func NewLeaveAPI(engine *gin.Engine, leaveApplicationService *service.LeaveApplicationService) *LeaveAPI {
	return &LeaveAPI{leaveApplicationService: leaveApplicationService}
}
