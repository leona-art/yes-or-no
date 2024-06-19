package usecase

import (
	"context"
	"gomini/domain"
)

type AiService interface {
	Ask(ctx context.Context, topic domain.Topic, prompt string) (*domain.Answer, error)
}
