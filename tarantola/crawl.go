package tarantola

type CrawlKit interface {
	Crawl() error
	CrawlErrorHandler(err error)
	DataProcessErrorHandler(err error)
	DataProcessorHandler(crawlRes interface{}) error
	Init()
	GetResChan() chan interface{}
	GetCrawlerName() string
}
