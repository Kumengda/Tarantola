package main

import (
	. "github.com/Kumengda/Tarantola/runtime"
	"github.com/Kumengda/Tarantola/tarantola"
	"github.com/Kumengda/Tarantola/testCrawler"
)

func main() {
	myCrawler := testCrawler.NewMyCrawler(tarantola.BaseOptions{})

	t := tarantola.NewTarantola()
	t.AddCrawler(myCrawler)
	MainInsp.Print(LEVEL_INFO, Text("asda"))
	t.MonoCrawl(1)
}
