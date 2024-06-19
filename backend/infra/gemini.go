package infra

import (
	"context"
	"fmt"
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

func (g *Gemini) GenerateText(ctx context.Context, prompt string) (string, error) {
	model := g.client.GenerativeModel(GENERATIVE_MODEL)
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"answer": {
				Type: genai.TypeString,
				Enum: []string{"yes", "no", "neither", "unrelated"},
			},
		},
	}
	p := []genai.Part{
		genai.Text(`
			次のトピックに関して質問されるので,それにyes/no/neither/unrelatedのいずれかで答えてください。
		`),
		genai.Text(`
			topic: お寿司
			`),
		genai.Text(fmt.Sprintf("question: %s", prompt)),
	}
	res, err := model.GenerateContent(ctx, p...)
	if err != nil {
		return "", err
	}

	for _, part := range res.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			log.Println(txt)
		}
	}
	return "", nil
}

func (g *Gemini) Close() {
	g.client.Close()
}
