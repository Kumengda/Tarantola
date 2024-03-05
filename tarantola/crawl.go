package tarantola

import "github.com/Kumengda/Tarantola/request"

type CrawlKit interface {
	Crawl() error
	crawlErrorHandler(err error)
	dataProcessErrorHandler(err error)
	dataProcessHandler(crawlRes interface{}, request *request.HttpRequest) error
	init()
	getResChan() chan interface{}
	getHttpRequest() *request.HttpRequest
	getCrawlerName() string
}
