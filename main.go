package main

import (
	. "github.com/Kumengda/Tarantola/runtime"
	"github.com/Kumengda/Tarantola/tarantola"
	"github.com/Kumengda/Tarantola/testCrawler"
)

func main() {
	myCrawler := testCrawler.NewMyCrawler(tarantola.BaseOptions{
		Timeout:  10,
		ProxyUrl: "http://127.0.0.1:7890",
	})
	t := tarantola.NewTarantola()
	t.AddCrawler(myCrawler)
	MainInsp.Print(LEVEL_INFO, Text("asda"))
	t.MonoCrawl(1)
}
