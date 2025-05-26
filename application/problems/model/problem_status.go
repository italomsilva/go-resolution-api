package model

type ProblemStatus int 

const (
	Open ProblemStatus = 0
	InProgress ProblemStatus = 1
	Resolved ProblemStatus = 2
	Canceled ProblemStatus = 3
)