package domain

import "fmt"

type NewsEvent struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

func (n *NewsEvent) GetID() string {
	return fmt.Sprintf("newsapi-%s-%s", n.Source.Name, n.URL)
}
