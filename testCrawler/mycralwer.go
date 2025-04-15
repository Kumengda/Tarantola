package testCrawler

import (
	"github.com/Kumengda/Tarantola/tarantola"
)

type MyCrawler struct {
	tarantola.BaseCrawler
}

func (m *MyCrawler) Crawl() error {
	return nil
}

func NewMyCrawler(options tarantola.BaseOptions) *MyCrawler {
	return &MyCrawler{
		BaseCrawler: tarantola.BaseCrawler{
			BaseOptions: options,
		},
	}
}
