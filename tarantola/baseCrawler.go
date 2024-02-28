package tarantola

import (
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
	dataProcessFunc        func(crawlRes interface{}) error
	dataProcessErrorFunc   func(err error)
	crawlErrorFunc         func(err error)
	BaseUrl                string
	resChain               chan interface{}
	HttpRequest            *request.HttpRequest
	JsExec                 *otto.Otto
	CrawlerName            string
	chromeJsExecuteTimeout int
}

func (b *BaseCrawler) init() {
	if b.resChain == nil {
		b.resChain = make(chan interface{})
	}
	if b.HttpRequest == nil {
		b.HttpRequest = request.NewHttpRequest(b.Headers, b.ProxyUrl, b.Timeout, b.RandomWaitTimeoutMin, b.RandomWaitTimeoutMax)
	}
	if b.JsExec == nil {
		b.JsExec = otto.New()
	}
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
func (b *BaseCrawler) getResChan() chan interface{} {
	return b.resChain
}
func (b *BaseCrawler) getCrawlerName() string {
	return b.CrawlerName
}

func (b *BaseCrawler) dataProcessErrorHandler(err error) {
	if b.dataProcessErrorFunc != nil {
		b.dataProcessErrorFunc(err)
	}
}
func (b *BaseCrawler) dataProcessHandler(crawlRes interface{}) error {
	if b.dataProcessFunc != nil {
		return b.dataProcessFunc(crawlRes)
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
		b.HttpRequest = request.NewHttpRequest(b.Headers, b.ProxyUrl, b.Timeout, b.RandomWaitTimeoutMin, b.RandomWaitTimeoutMin)
	}
	b.HttpRequest.SetRetryFunc(retryFunc, maxRetry)
}
func (b *BaseCrawler) SetCrawlErrorHandler(crawlErrorFunc func(err error)) {
	b.crawlErrorFunc = crawlErrorFunc
}
func (b *BaseCrawler) SetDataProcessErrorHandler(dataProcessErrorFunc func(err error)) {
	b.dataProcessErrorFunc = dataProcessErrorFunc
}
func (b *BaseCrawler) SetDataProcessFunc(dataProcessFunc func(crawlRes interface{}) error) {
	b.dataProcessFunc = dataProcessFunc
}
