package competition

type SubmissionInfo struct {
	Id       string
	Status   string
	Score    float32
	Phase    string
	PlayerId string
}

type CompetitionService interface {
	SetSubmissionInfo(string, *SubmissionInfo) error
}
