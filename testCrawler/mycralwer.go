package testCrawler

import (
	"fmt"
	"github.com/Kumengda/Tarantola/tarantola"
)

type MyCrawler struct {
	tarantola.BaseCrawler
}

func NewMyCrawler(options tarantola.BaseOptions) *MyCrawler {
	return &MyCrawler{
		BaseCrawler: tarantola.BaseCrawler{
			BaseOptions: options,
		},
	}
}

func (m *MyCrawler) Crawl() error {
	return nil
}

func (m *MyCrawler) CrawlErrorHandler(err error) {

}
func (m *MyCrawler) DataProcessErrorHandler(err error) {

}

func (m *MyCrawler) DataProcessorHandler(crawlRes interface{}) error {
	fmt.Println("模拟数据接收")
	fmt.Println(crawlRes)
	return nil
}
