package search

import (
	"log"
	"fmt"
)

type Result struct {
	Field 	string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Print(err)
		return
	}

	for _, result := range searchResult {
		results <- result
	}
}

func Display(results chan *Result) {
	// 通道会一直被阻塞，知道有结果写入
	// 一旦通道被关闭，for循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Mather already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}