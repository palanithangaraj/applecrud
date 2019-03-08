package model

type JokeType struct {
	Type  string   `json:"type"`
	Value JokeItem `json:"value"`
}

type JokeItem struct {
	Id         string   `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"catagories"`
	//Could add date times
}
