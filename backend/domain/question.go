package domain

type Question struct {
	Content string `json:"content"`
}

func NewQuestion(content string) *Question {
	return &Question{
		Content: content,
	}
}
