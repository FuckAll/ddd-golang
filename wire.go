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

var leaveAPISet = wire.NewSet(leaveDomainService.NewLeaveDomainService, personDomainService.NewPersonDomainService, ruleDomainService.NewApprovalRuleDomainService)
var leaveDomainServiceSet = wire.NewSet(event.NewEventPublisher, wire.Bind(new(leaveRepositoryFacade.ILeaveRepository), new(*leavePersistence.LeaveRepositoryImpl)), leavePersistence.NewLeaveRepositoryImpl, leaveDomainService.NewLeaveFactory)
var personDomainServiceSet = wire.NewSet(wire.Bind(new(personRepositoryFacade.PersonRepository), new(*personPersistence.PersonRepositoryImpl)), personPersistence.NewPersonRepositoryImpl, personDomainService.NewPersonFactory)
var approvalRuleDomainServiceSet = wire.NewSet(wire.Bind(new(ruleRepositoryFacade.ApprovalRuleRepository), new(*rulePersistence.ApprovalRuleRepositoryImpl)), rulePersistence.NewApprovalRuleRepositoryImpl)

func InitLeaveAPI() *facade.LeaveAPI {
	wire.Build(facade.NewLeaveAPI, wire.Value(DefaultEngine), service.NewLeaveApplicationService, leaveAPISet, leaveDomainServiceSet, personDomainServiceSet, approvalRuleDomainServiceSet)
	return &facade.LeaveAPI{}
}
