package tarantola

import (
	. "github.com/Kumengda/Tarantola/runtime"
	"sync"
)

type Tarantola struct {
	Crawlers []CrawlKit
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

func (t *Tarantola) MonoCrawl() (finalRes []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				err = r.(error)
			}
		}
	}()
	var wg sync.WaitGroup
	for _, c := range t.Crawlers {
		c.init()
		resChan := c.getResChan()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for res := range resChan {
				finalRes = append(finalRes, res)
				err := c.dataProcessHandler(res, c.getHttpRequest())
				if err != nil {
					c.dataProcessErrorHandler(err)
				}
			}
			return
		}()
		err := c.Crawl()
		if err != nil {
			c.crawlErrorHandler(err)
		}
		close(resChan)
	}
	wg.Wait()
	return finalRes, err
}

func (t *Tarantola) MultiCrawl() error {
	return nil
}
