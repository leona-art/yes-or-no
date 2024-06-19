package usecase

import (
	"context"
	"gomini/domain"
)

type AskService struct {
	AiService AiService
}

func NewAskService(aiService AiService) *AskService {
	return &AskService{
		AiService: aiService,
	}
}

func (a *AskService) Ask(ctx context.Context, topic string, prompt string) (string, error) {
	t := domain.Topic{
		Content:     topic,
		Description: "",
	}
	answer, err := a.AiService.Ask(ctx, t, prompt)
	if err != nil {
		return "", err
	}
	return answer.Answer, nil
}
