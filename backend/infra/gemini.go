package infra

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gomini/domain"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const (
	GENERATIVE_MODEL = "gemini-1.5-pro-latest"
)

type Gemini struct {
	client *genai.Client
}

func NewGemini(ctx context.Context) (*Gemini, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &Gemini{
		client: client,
	}, nil
}

func (g *Gemini) Ask(ctx context.Context, topic domain.Topic, prompt string) (*domain.Answer, error) {
	model := g.client.GenerativeModel(GENERATIVE_MODEL)
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"answer": {
				Type: genai.TypeString,
				Enum: []string{"yes", "no", "neither"},
			},
			"reason": {
				Type: genai.TypeString,
			},
		},
	}
	p := []genai.Part{
		genai.Text(`
			"topic"と"question"があります。
			"question"は"topic"に関する質問です。
			"topic"に関する質問に対して、"yes"、"no"、"neither"のいずれかで回答してください。
			"reason"には回答の理由を記入してください。
		`),
		genai.Text(fmt.Sprintf("topic: %s", topic)),
		genai.Text(fmt.Sprintf("question: %s", prompt)),
	}
	res, err := model.GenerateContent(ctx, p...)
	if err != nil {
		return nil, err
	}

	for _, part := range res.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			var answer domain.Answer
			err := json.Unmarshal([]byte(txt), &answer)
			if err != nil {
				log.Printf("failed to unmarshal: %v", txt)
				return nil, err
			}
			return &answer, nil
		}
	}
	return nil, errors.New("no answer")
}

func (g *Gemini) Close() {
	g.client.Close()
}
