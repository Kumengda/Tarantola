package main

import (
	"github.com/Kumengda/Tarantola/tarantola"
	"github.com/Kumengda/Tarantola/testCrawler"
)

func main() {
	myCrawler := testCrawler.NewMyCrawler(tarantola.BaseOptions{})
	t := tarantola.NewTarantola()

	t.AddCrawler(myCrawler)

	t.MonoCrawl()

}
