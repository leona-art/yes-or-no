package domain

type Topic struct {
	Content     string `json:"content"`
	Description string `json:"description"`
}

func NewTopic(content, description string) *Topic {
	return &Topic{
		Content:     content,
		Description: description,
	}
}
