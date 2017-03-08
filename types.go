package main

// StoryItem is a story/post on HN
type StoryItem struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	StoryURL string `json:"story_url"`
	Score    int    `json:"score"`
}
