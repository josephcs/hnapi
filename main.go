package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const hnURL string = "https://news.ycombinator.com"

func prefixRelativeURL(relativePath string) string {
	return hnURL + "/" + relativePath
}

func getStories() []StoryItem {
	stories := make([]StoryItem, 0)
	doc, err := goquery.NewDocument(hnURL)
	if err != nil {
		log.Println(err)
	}

	doc.Find(".athing").Each(func(i int, s *goquery.Selection) {
		domItem := s.Find(".storylink")
		title := domItem.Text()
		link, _ := domItem.Attr("href")
		subtextRow := s.Next()
		storyLink, _ := subtextRow.Find(".age a").Attr("href")
		scoreAsString := strings.Split(subtextRow.Find(".score").Text(), " ")[0]
		// TODO: ^ Use RegExp
		score, _ := strconv.Atoi(scoreAsString)
		story := StoryItem{
			Title:    title,
			URL:      link,
			StoryURL: prefixRelativeURL(storyLink),
			Score:    score}
		stories = append(stories, story)
	})
	return stories
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		stories := getStories()
		if len(stories) == 0 {
			// Either HN is down (mostly not)
			// or you're making too many zero-interval requests
			c.Status(204)
			return
		}
		c.JSON(200, gin.H{"posts": stories})
	})
	r.Run(":8080")
}
