package tarantola

type CrawlKit interface {
	Crawl() error
	crawlErrorHandler(err error)
	dataProcessErrorHandler(err error)
	dataProcessHandler(crawlRes interface{}) error
	init()
	getResChan() chan interface{}
	getCrawlerName() string
}
