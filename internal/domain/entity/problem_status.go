package entity

type ProblemStatus int

const (
	ProblemStatusOpen       ProblemStatus = 0
	ProblemStatusInProgress ProblemStatus = 1
	ProblemStatusResolved   ProblemStatus = 2
	ProblemStatusCanceled   ProblemStatus = 3
)
