package tarantola

import (
	"github.com/Kumengda/Tarantola/request"
	"net/http"
)

type BaseOptions struct {
	Headers              map[string]interface{}
	ProxyUrl             string
	Timeout              int
	RandomWaitTimeoutMin int
	RandomWaitTimeoutMax int
}

type BaseCrawler struct {
	BaseOptions
	BaseUrl     string
	resChain    chan interface{}
	HttpRequest *request.HttpRequest
	CrawlerName string
}

func (b *BaseCrawler) Init() {
	if b.resChain == nil {
		b.resChain = make(chan interface{})
	}
	if b.HttpRequest == nil {
		b.HttpRequest = request.NewHttpRequest(b.Headers, b.ProxyUrl, b.Timeout, b.RandomWaitTimeoutMin, b.RandomWaitTimeoutMin)
	}
}

func (b *BaseCrawler) PushResult(res interface{}) {
	b.resChain <- res
}
func (b *BaseCrawler) GetResChan() chan interface{} {
	return b.resChain
}
func (b *BaseCrawler) GetCrawlerName() string {
	return b.CrawlerName
}

func (b *BaseCrawler) SetRetryFunc(retryFunc func(respData []byte, respHeader http.Header, err error) bool, maxRetry int) {
	b.HttpRequest.SetRetryFunc(retryFunc, maxRetry)
}
