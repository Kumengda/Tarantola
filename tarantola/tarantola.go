package tarantola

import (
	"context"
	"github.com/B9O2/Multitasking"
	. "github.com/Kumengda/Tarantola/runtime"
)

type Tarantola struct {
	Crawlers    []CrawlKit
	errCallBack func(controller Multitasking.Controller, err error)
}

func NewTarantola() *Tarantola {
	InitDecoration()
	return &Tarantola{}
}

func (t *Tarantola) AddCrawler(Crawlers ...CrawlKit) {
	for _, c := range Crawlers {
		t.Crawlers = append(t.Crawlers, c)
	}
}
func (t *Tarantola) ClearCrawler() {
	t.Crawlers = nil
}
func (t *Tarantola) SetMtErrorCallBack(errCallBack func(controller Multitasking.Controller, err error)) {
	t.errCallBack = errCallBack
}

func (t *Tarantola) MonoCrawl(threads uint) []interface{} {
	crawlerMT := Multitasking.NewMultitasking("crawler", nil)
	crawlerMT.SetErrorCallback(t.errCallBack)
	crawlerMT.Register(func(dc Multitasking.DistributeController) {
		for _, c := range t.Crawlers {
			dc.AddTask(c)
		}
	}, func(ec Multitasking.ExecuteController, a any) any {
		crawler := a.(CrawlKit)
		crawler.init()
		resChan := crawler.getResChan()
		go func() {
			err := crawler.Crawl()
			if err != nil {
				crawler.crawlErrorHandler(err)
			}
			close(resChan)
		}()
		dataProcessMT := Multitasking.NewMultitasking("dataProcessMT", nil)
		middleware := crawler.getMiddleware()
		if middleware != nil {
			dataProcessMT.SetResultMiddlewares(middleware)
		}
		dataProcessMT.Register(func(dc Multitasking.DistributeController) {
			for res := range resChan {
				dc.AddTask(res)
			}
		}, func(ec Multitasking.ExecuteController, a any) any {
			err := crawler.dataProcessHandler(a, crawler.getHttpRequest())
			if err != nil {
				crawler.dataProcessErrorHandler(err)
			}
			return a
		})
		finalRes, _ := dataProcessMT.Run(context.Background(), 1)
		return finalRes
	})
	finalRes, _ := crawlerMT.Run(context.Background(), threads)

	//var wg sync.WaitGroup
	//for _, c := range t.Crawlers {
	//	c.init()
	//	resChan := c.getResChan()
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for res := range resChan {
	//			finalRes = append(finalRes, res)
	//			err := c.dataProcessHandler(res, c.getHttpRequest())
	//			if err != nil {
	//				c.dataProcessErrorHandler(err)
	//			}
	//		}
	//		return
	//	}()
	//	err := c.Crawl()
	//	if err != nil {
	//		c.crawlErrorHandler(err)
	//	}
	//	close(resChan)
	//}
	//wg.Wait()
	return finalRes
}
