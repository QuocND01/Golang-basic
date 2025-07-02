package model

type Article struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	URL       string `json:"url"`
	Published string `json:"published"`
	Source    string `json:"source"`
}
