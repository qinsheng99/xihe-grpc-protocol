package competition

type CompetitionIndex struct {
	Id string
}

type SubmissionInfo struct {
	Id       string
	Status   string
	Score    float32
	Phase    string
	PlayerId string
}

type CompetitionService interface {
	SetSubmissionInfo(*CompetitionIndex, *SubmissionInfo) error
}
