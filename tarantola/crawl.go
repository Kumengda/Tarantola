package tarantola

import (
	"github.com/B9O2/Multitasking"
	"github.com/Kumengda/Tarantola/request"
)

type CrawlKit interface {
	Crawl() error
	crawlErrorHandler(err error)
	dataProcessErrorHandler(err error)
	dataProcessHandler(crawlRes interface{}, request *request.HttpRequest) error
	init()
	getResChan() chan interface{}
	getMiddleware() Multitasking.Middleware
	getHttpRequest() *request.HttpRequest
	getCrawlerName() string
}
