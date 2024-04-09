package testCrawler

import (
	"github.com/Kumengda/Tarantola/tarantola"
	"time"
)

type MyCrawler struct {
	tarantola.BaseCrawler
}

func (m *MyCrawler) Crawl() error {
	for {
		m.PushResult("111")
		time.Sleep(1 * time.Second)
	}

	return nil
}

func NewMyCrawler(options tarantola.BaseOptions) *MyCrawler {
	return &MyCrawler{
		BaseCrawler: tarantola.BaseCrawler{
			BaseOptions: options,
		},
	}
}
