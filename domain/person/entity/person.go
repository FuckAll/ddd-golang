package entity

import (
	"github.com/FuckAll/ddd-golang/domain/shared"
	"time"
)

type Person struct {
	PersonId       string
	PersonName     string
	PersonType     shared.PersonType
	Relationships  []Relationship
	RoleLevel      int
	CreateTime     time.Time
	LastModifyTime time.Time
	Status         PersonStatus
}

func (p *Person) Create() {
	p.CreateTime = time.Now()
	p.Status = Enable
}

func (p *Person) Enable() {
	p.LastModifyTime = time.Now()
	p.Status = Enable
}

func (p *Person) Disable() {
	p.LastModifyTime = time.Now()
	p.Status = Disable
}
