package search

import (
	"log"
	"fmt"
)

type Result struct {
	Field 	string
	Content string
}

// 这个接口声明了结构类型(struct)或者具名类型(type)需要实现的行为。
// 一个接口的行为最终由在这个接口类型中声明的方法决定。
// 一个类型如果实现了该接口类型所声明的所有方法，那该类型就是该接口（或理解为继承该接口）
// 如果接口类型只包含一个方法，那么这 个类型的名字以 er 结尾
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// matcher参数只有实现了 Matcher 接口的值或者指针能被接受。
	// 因为 defaultMatcher 类型使用值作为接收者，实现了这个接口，
	// 所以 defaultMatcher 类型的 值或者指针可以传入这个函数
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
	// 通道会一直被阻塞，直到有结果写入
	// 一旦通道被关闭，for循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}


// 将一个 Matcher 值加入到保存注册匹配器的映射中
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Mather already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}