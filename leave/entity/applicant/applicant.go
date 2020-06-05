package applicant

// Applicant 申请人值对象
type Applicant struct {
	PersonId   string
	PersonName string
	PersonType string
}

func NewApplicant(personId string, personName string, personType string) *Applicant {
	return &Applicant{PersonId: personId, PersonName: personName, PersonType: personType}
}
