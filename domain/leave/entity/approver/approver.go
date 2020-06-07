package approver

type Approver struct {
	personId   string
	personName string
	level      int
}

func NewApprover(personId string, personName string, level int) *Approver {
	return &Approver{personId: personId, personName: personName, level: level}
}
