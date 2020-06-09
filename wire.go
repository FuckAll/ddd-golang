//+build wireinject

package main

import (
	"github.com/FuckAll/ddd-golang/application/service"
	leaveRepositoryFacade "github.com/FuckAll/ddd-golang/domain/leave/repository/facade"
	leavePersistence "github.com/FuckAll/ddd-golang/domain/leave/repository/persistence"
	leaveDomainService "github.com/FuckAll/ddd-golang/domain/leave/service"
	personRepositoryFacade "github.com/FuckAll/ddd-golang/domain/person/repository/facade"
	personPersistence "github.com/FuckAll/ddd-golang/domain/person/repository/persistence"
	personDomainService "github.com/FuckAll/ddd-golang/domain/person/service"
	ruleRepositoryFacade "github.com/FuckAll/ddd-golang/domain/rule/repository/facade"
	rulePersistence "github.com/FuckAll/ddd-golang/domain/rule/repository/persistence"
	ruleDomainService "github.com/FuckAll/ddd-golang/domain/rule/service"
	"github.com/FuckAll/ddd-golang/infrastructure/common/event"
	"github.com/FuckAll/ddd-golang/interface/facade"
	"github.com/google/wire"
)

func InitAuthApi() *facade.AuthApi {
	wire.Build(facade.NewAuthApi, wire.Value(DefaultEngine), service.NewLoginApplicationService)
	return &facade.AuthApi{}
}

var _ leaveRepositoryFacade.ILeaveRepository = &leavePersistence.LeaveRepositoryImpl{}

// NewLeaveApplicationService 所需依赖
var leaveApplicationServiceSet = wire.NewSet(
	leaveDomainService.NewLeaveDomainService,
	personDomainService.NewPersonDomainService,
	ruleDomainService.NewApprovalRuleDomainService)

// LeaveDomainService 所需依赖
var leaveDomainServiceSet = wire.NewSet(
	event.NewEventPublisher,
	wire.Bind(new(leaveRepositoryFacade.ILeaveRepository), new(*leavePersistence.LeaveRepositoryImpl)), // 注入interface方法要用Bind
	leavePersistence.NewLeaveRepositoryImpl, // 需要注入interface的实现类
	leaveDomainService.NewLeaveFactory)

// PersonDomainService 所需依赖
var personDomainServiceSet = wire.NewSet(
	wire.Bind(new(personRepositoryFacade.PersonRepository), new(*personPersistence.PersonRepositoryImpl)),
	personPersistence.NewPersonRepositoryImpl,
	personDomainService.NewPersonFactory)

var approvalRuleDomainServiceSet = wire.NewSet(
	wire.Bind(new(ruleRepositoryFacade.ApprovalRuleRepository), new(*rulePersistence.ApprovalRuleRepositoryImpl)),
	rulePersistence.NewApprovalRuleRepositoryImpl)

func InitLeaveAPI() *facade.LeaveAPI {
	wire.Build(
		facade.NewLeaveAPI,
		wire.Value(DefaultEngine),
		service.NewLeaveApplicationService,
		leaveApplicationServiceSet,
		leaveDomainServiceSet,
		personDomainServiceSet,
		approvalRuleDomainServiceSet)
	return &facade.LeaveAPI{}
}
