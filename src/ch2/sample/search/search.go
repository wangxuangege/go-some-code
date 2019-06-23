package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		// 调用后会终止应用程序
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接受匹配后的结果
	results := make(chan *Result)

	// 构造一个waitGroup，处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理，需要等待的数量为数据源的个数
	waitGroup.Add(len(feeds))

	// 每个数据源启动一个goroutine查找结果
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine监控是否所有工作完成
	go func() {
		waitGroup.Wait()

		// 关闭通道的方式，通知Display函数
		close(results)
	}()

	// 启动函数，显示所有的结果，把那个且在最后一个结果显示完后返回
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
