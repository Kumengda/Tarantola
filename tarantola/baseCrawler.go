package tarantola

import (
	"errors"
	"fmt"
	"github.com/B9O2/Multitasking"
	"github.com/Kumengda/Tarantola/request"
	"github.com/Kumengda/easyChromedp/chrome"
	"github.com/chromedp/chromedp"
	"github.com/robertkrimen/otto"
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
	dataProcessFunc        func(crawlRes interface{}, httpRequest *request.HttpRequest) error
	dataProcessErrorFunc   func(err error)
	crawlErrorFunc         func(err error)
	BaseUrl                string
	resChain               chan interface{}
	HttpRequest            *request.HttpRequest
	JsExec                 *otto.Otto
	CrawlerName            string
	chromeJsExecuteTimeout int
	middleware             Multitasking.Middleware
}

func (b *BaseCrawler) init() {
	b.resChain = make(chan interface{})
	b.HttpRequest = request.NewHttpRequest(b.Headers, b.ProxyUrl, b.Timeout, b.RandomWaitTimeoutMin, b.RandomWaitTimeoutMax)
	b.JsExec = otto.New()
	b.chromeJsExecuteTimeout = 10
}
func (b *BaseCrawler) SetChromeJsExecuteTimeout(timeout int) {
	b.chromeJsExecuteTimeout = timeout
}
func (b *BaseCrawler) ExecJsWithChrome(jsCode string) (interface{}, error) {
	myChrome, err := chrome.NewChromeWithTimout(
		b.chromeJsExecuteTimeout,
	)
	defer func() {
		if myChrome != nil {
			myChrome.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = myChrome.RunWithOutListen(
		chromedp.Evaluate(jsCode, &result),
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BaseCrawler) PushResult(res interface{}) {
	b.resChain <- res
}
func (b *BaseCrawler) SetMiddleware(middleware Multitasking.Middleware) {
	b.middleware = middleware
}
func (b *BaseCrawler) getMiddleware() Multitasking.Middleware {
	return b.middleware
}
func (b *BaseCrawler) getResChan() chan interface{} {
	return b.resChain
}
func (b *BaseCrawler) getHttpRequest() *request.HttpRequest {
	return b.HttpRequest
}
func (b *BaseCrawler) getCrawlerName() string {
	return b.CrawlerName
}

func (b *BaseCrawler) dataProcessErrorHandler(err error) {
	if b.dataProcessErrorFunc != nil {
		b.dataProcessErrorFunc(err)
	}
}
func (b *BaseCrawler) dataProcessHandler(crawlRes interface{}, httpRequest *request.HttpRequest) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				err = r.(error)
			} else {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()
	if b.dataProcessFunc != nil {
		return b.dataProcessFunc(crawlRes, httpRequest)
	}
	return nil
}
func (b *BaseCrawler) crawlErrorHandler(err error) {
	if b.crawlErrorFunc != nil {
		b.crawlErrorFunc(err)
	}
}

func (b *BaseCrawler) SetRetryFunc(retryFunc func(respData []byte, respHeader http.Header, err error) bool, maxRetry int) {
	if b.HttpRequest == nil {
		b.HttpRequest = request.NewHttpRequest(b.Headers, b.ProxyUrl, b.Timeout, b.RandomWaitTimeoutMin, b.RandomWaitTimeoutMax)
	}
	b.HttpRequest.SetRetryFunc(retryFunc, maxRetry)
}
func (b *BaseCrawler) SetCrawlErrorHandler(crawlErrorFunc func(err error)) {
	b.crawlErrorFunc = crawlErrorFunc
}
func (b *BaseCrawler) SetDataProcessErrorHandler(dataProcessErrorFunc func(err error)) {
	b.dataProcessErrorFunc = dataProcessErrorFunc
}
func (b *BaseCrawler) SetDataProcessFunc(dataProcessFunc func(crawlRes interface{}, httpRequest *request.HttpRequest) error) {
	b.dataProcessFunc = dataProcessFunc
}
