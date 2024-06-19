package domain

import "errors"

const (
	Yes       = "yes"
	No        = "no"
	Neither   = "neither"
	Unrelated = "unrelated"
)

var (
	errInvalidateAnswer = errors.New("invalid answer")
)

type Answer struct {
	Answer string `json:"answer"`
}

func NewAnswer(answer string) (*Answer, error) {
	if answer != Yes && answer != No && answer != Neither && answer != Unrelated {
		return nil, errInvalidateAnswer
	}
	return &Answer{
		Answer: answer,
	}, nil
}
