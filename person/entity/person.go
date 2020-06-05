package entity

import (
	valueobject2 "github.com/FuckAll/ddd-golang/person/entity/valueobject"
	"time"
)

type Person struct {
	personId       string
	personName     string
	personType     valueobject2.PersonType
	relationships  []Relationship
	roleLevel      int
	createTime     time.Time
	lastModifyTime time.Time
	status         valueobject2.PersonStatus
}
