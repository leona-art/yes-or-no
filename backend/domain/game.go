package domain

type Game struct {
	Topic  *Topic   `json:"topic"`
	Record []Record `json:"record"`
}

type Record struct {
	Question *Question `json:"question"`
	Answer   *Answer   `json:"answer"`
}

func NewGame(topic *Topic) *Game {
	return &Game{
		Topic:  topic,
		Record: make([]Record, 0),
	}
}
