package approver

type Approver struct {
	PersonId   string
	PersonName string
	Level      int
}

func NewApprover(personId string, personName string, level int) *Approver {
	return &Approver{PersonId: personId, PersonName: personName, Level: level}
}
