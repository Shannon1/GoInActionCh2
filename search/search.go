package search

import (
	"log"
	"sync"
)


// 这个变量没有定义在任何函数作用域内，所以会被当成包级变量
// 注册用于搜索的匹配器的映射
// 实际上只有defaultMatcher和rssMather两个
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 获取需要搜索的数据源列表([]*Feed)
	// 数据源(Feed)的类型定义见feed.go，该类型作用是将data中的数据进行反序列化存储
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 waitGroup，以便处理所有的数据源
	// WaitGroup 是一个计数信号量，我们可以利用它来统计所有的 goroutine 是不是都完成了工作。
	var waitGroup sync.WaitGroup

	// 设置需要等待处理的每个数据源 goroutine 的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			// Match 函数会搜索数 据源的数据，并将匹配结果输出到 results 通道
			Match(matcher, feed, searchTerm, results)

			// 递减 WaitGroup 的计数。一旦每个 goroutine 都执行调用 Match 函数和 Done 方法，程序就知道每个数据源都处理完成。
			waitGroup.Done()
		} (matcher, feed)
	}

	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式，通知 Display 函数可以退出程序了
		close(results)
	} ()

	// Display函数内阻塞等待通道results中的数据并打印到屏幕上，results通道关闭后，Display函数将退出
	Display(results)
}
