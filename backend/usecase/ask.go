package usecase

import (
	"context"
	"gomini/domain"
	"log"
)

type AskService struct {
	AiService AiService
}

func NewAskService(aiService AiService) *AskService {
	return &AskService{
		AiService: aiService,
	}
}

type GameAnswer struct {
	Answer string
	Finish bool
}

func (a *AskService) Ask(ctx context.Context, topic string, prompt string) (*GameAnswer, error) {
	t := domain.Topic{
		Content:     topic,
		Description: "",
	}
	answer, err := a.AiService.Ask(ctx, t, prompt)
	if err != nil {
		return nil, err
	}
	log.Printf("reason: %s", answer.Reason)
	return &GameAnswer{
		Answer: answer.Answer,
		Finish: answer.Finish,
	}, nil
}
