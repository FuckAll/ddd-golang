package po

import (
	"github.com/FuckAll/ddd-golang/domain/person/entity"
	"github.com/FuckAll/ddd-golang/domain/shared"
	"time"
)

type PersonPO struct {
	PersonId       string
	PersonName     string
	DepartmentId   string
	PersonType     shared.PersonType
	LeaderId       string
	RoleLevel      int
	CreateTime     time.Time
	LastModifyTime time.Time
	Status         entity.PersonStatus
	RelationshipPO RelationshipPO
}
